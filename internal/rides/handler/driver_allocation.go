package handler

// const (
// 	EventNotifyDriverNewCustomer = "notify_driver_new_customer"
// )

// type DriverAllocationMessage struct {
// 	Event string      `json:"event"`
// 	Data  interface{} `json:"data"`
// }

// type Coordinate struct {
// 	Latitude  float64 `json:"latitude"`
// 	Longitude float64 `json:"longitude"`
// }

// type NewCustomerMessage struct {
// 	TripID         string     `json:"trip_id"`
// 	Distance       float32    `json:"distance"`
// 	PickupLocation Coordinate `json:"pickup_location"`
// 	Destination    Coordinate `json:"destination"`
// }

// func (h *ridesHandler) DriverAllocation(c *gin.Context) {
// 	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		logger.Error(c.Request.Context(), "error upgrade to websocket", map[string]interface{}{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	defer conn.Close()

// 	go func() {
// 		for {
// 			messageType, message, err := conn.ReadMessage()
// 			if err != nil {
// 				log.Println("Error reading message:", err)
// 				break
// 			}

// 			fmt.Printf("Received: %s\n", message)

// 			if err := conn.WriteMessage(messageType, message); err != nil {
// 				log.Println("Error writing message:", err)
// 				break
// 			}
// 		}
// 	}()
// }
