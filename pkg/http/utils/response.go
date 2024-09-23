package http_utils

import "net/http"

const MessageInternalServerError = "Internal server error"

type Response struct {
	Meta MetaResponse `json:"meta"`
	Data interface{}  `json:"data,omitempty"`
}

type MetaResponse struct {
	Code               int      `json:"code"`
	Message            string   `json:"message"`
	ErrorDetailMessage []string `json:"errors,omitempty"`
}

func NewFailedResponse(errorCode int, message string) Response {
	return Response{
		Meta: MetaResponse{
			Code:    errorCode,
			Message: message,
		},
	}
}

func NewSuccessResponse(data interface{}) Response {
	return Response{
		Meta: MetaResponse{
			Code:    http.StatusOK,
			Message: "success",
		},
		Data: data,
	}
}
