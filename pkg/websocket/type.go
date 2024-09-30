package websocket

type WebsocketInterface interface {
	WriteMessage(messageType int, data []byte) error
}
