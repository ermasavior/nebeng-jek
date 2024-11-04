package utils

import (
	"encoding/json"
	"io"
)

func ParseResponseBody(resBody io.ReadCloser, res interface{}) error {
	b, err := io.ReadAll(resBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &res)
}
