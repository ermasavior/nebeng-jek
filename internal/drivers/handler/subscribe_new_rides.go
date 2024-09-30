package handler

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/pkg/amqp"
	"nebeng-jek/pkg/logger"
	pkg_ws "nebeng-jek/pkg/websocket"

	"github.com/gorilla/websocket"
)

func (h *driversHandler) SubscribeNewRides(ctx context.Context, ridesChannel amqp.AMQPChannel) {
	err := ridesChannel.ExchangeDeclare(
		constants.RideRequestsExchange,
		"fanout", // exchange type: fanout
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		logger.Fatal(context.Background(), "failed to declare an amqp exchange", map[string]interface{}{
			"error": err,
		})
	}

	q, err := ridesChannel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		logger.Fatal(context.Background(), "failed to declare amqp queue", map[string]interface{}{
			"error": err,
		})
	}

	err = ridesChannel.QueueBind(
		q.Name,                         // queue name
		"",                             // routing key
		constants.RideRequestsExchange, // exchange
		false,
		nil,
	)
	if err != nil {
		logger.Fatal(context.Background(), "failed to bind amqp queue", map[string]interface{}{
			"error": err,
		})
	}

	msgs, err := ridesChannel.Consume(
		q.Name,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		logger.Fatal(context.Background(), "failed to consume amqp queue", map[string]interface{}{
			"error": err,
		})
	}

	for msg := range msgs {
		var data model.RideRequestMessage
		err := json.Unmarshal(msg.Body, &data)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{"error": err})
		}

		h.broadcastToActiveDrivers(ctx, data)
	}
}

func (h *driversHandler) broadcastToActiveDrivers(ctx context.Context, msg model.RideRequestMessage) {
	for driver := range msg.AvailableDrivers {
		conn, ok := h.connStorage.Load(driver)
		if !ok {
			return
		}

		wsConn, ok := conn.(pkg_ws.WebsocketInterface)
		if !ok {
			logger.Error(ctx, "error loading driver connection websocket", nil)
			return
		}

		broadcastMsg := model.DriverAllocationMessage{
			Event: model.EventNewRideRequest,
			Data: model.NewRideRequestBroadcast{
				RideID:         msg.RideID,
				RiderID:        msg.RiderID,
				PickupLocation: msg.PickupLocation,
				Destination:    msg.Destination,
			},
		}
		msgBytes, err := json.Marshal(broadcastMsg)
		if err != nil {
			logger.Error(ctx, "error unmarshalling message broadcast", map[string]interface{}{
				"error": err,
			})
			return
		}

		if err := wsConn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
			logger.Error(ctx, "error broadcasting to drivers via websocket", map[string]interface{}{
				"error": err,
			})
			return
		}
	}
}
