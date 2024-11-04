package http_client

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	http_utils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/logger"
	"net/http"
)

var ErrorStatusCode = errors.New("response status is not success")

func RequestHTTPAndParseResponse(ctx context.Context, client *http.Client, transport RestTransport) (http_utils.ClientResponse, error) {
	httpRes, err := Request(ctx, client, transport)
	if err != nil {
		logger.Error(ctx, "failed http request", map[string]interface{}{
			logger.ErrorKey: err,
			"url":           transport.Url,
			"method":        transport.Method,
			"payload":       transport.Payload,
		})
		return http_utils.ClientResponse{}, err
	}
	defer httpRes.Body.Close()

	var res http_utils.ClientResponse
	_ = parseResponseBody(httpRes.Body, &res)

	if httpRes.StatusCode/100 != 2 {
		logger.Error(ctx, "error http response", map[string]interface{}{
			"status_code": httpRes.StatusCode,
			"response":    res,
			"url":         transport.Url,
			"method":      transport.Method,
			"payload":     transport.Payload,
		})
		return http_utils.ClientResponse{}, ErrorStatusCode
	}

	return res, nil
}

func parseResponseBody(resBody io.ReadCloser, res interface{}) error {
	b, err := io.ReadAll(resBody)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &res)
}
