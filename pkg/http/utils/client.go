package http_utils

import "encoding/json"

type ClientResponse struct {
	Meta MetaResponse    `json:"meta"`
	Data json.RawMessage `json:"data,omitempty"`
}
