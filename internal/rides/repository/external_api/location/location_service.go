package location

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/internal/rides/repository"
	"nebeng-jek/pkg/configs"
	http_utils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/http_client"
	"nebeng-jek/pkg/logger"
	"nebeng-jek/pkg/utils"
	"net/http"
)

const (
	addAvailableDriver         = "/v1/drivers/available"
	removeAvailableDriver      = "/v1/drivers/available/%d"
	getNearestAvailableDrivers = "/v1/drivers/available/nearby"
	getRidePath                = "/v1/drivers/ride-path"
)

var (
	ErrorInternalLocationService = errors.New("failed http request to location service")
	ErrorParsingResponse         = errors.New("error parsing response")
)

type locationRepo struct {
	HttpClient *http.Client
	BaseUrl    string
	APIKey     string
}

func NewLocationRepository(cfg *configs.Config, httpClient *http.Client) repository.RidesLocationRepository {
	return &locationRepo{
		HttpClient: httpClient,
		BaseUrl:    cfg.LocationServiceURL,
		APIKey:     cfg.LocationServiceAPIKey,
	}
}

func (s *locationRepo) AddAvailableDriver(ctx context.Context, driverID int64, location model.Coordinate) error {
	req := model.AddAvailableDriverRequest{
		DriverID: driverID,
		Location: location,
	}
	transport := http_client.RestTransport{
		Url:    s.BaseUrl + addAvailableDriver,
		Method: http.MethodPost,
		Header: http.Header{
			http_utils.HeaderApiKey: []string{s.APIKey},
		},
		Payload: req,
	}

	httpRes, err := http_client.Request(ctx, s.HttpClient, transport)
	if err != nil {
		logger.Error(ctx, "failed http request", map[string]interface{}{
			logger.ErrorKey: err,
			"url":           transport.Url,
			"method":        transport.Method,
			"payload":       req,
		})
		return err
	}
	defer httpRes.Body.Close()

	if httpRes.StatusCode/100 != 2 {
		var res http_utils.ClientResponse
		_ = utils.ParseResponseBody(httpRes.Body, &res)

		logger.Error(ctx, "error http response", map[string]interface{}{
			"status_code": httpRes.StatusCode,
			"response":    res,
			"url":         transport.Url,
			"method":      transport.Method,
			"payload":     req,
		})
		return ErrorInternalLocationService
	}

	return nil
}

func (s *locationRepo) RemoveAvailableDriver(ctx context.Context, driverID int64) error {
	transport := http_client.RestTransport{
		Url:    s.BaseUrl + fmt.Sprintf(removeAvailableDriver, driverID),
		Method: http.MethodDelete,
		Header: http.Header{
			http_utils.HeaderApiKey: []string{s.APIKey},
		},
	}

	httpRes, err := http_client.Request(ctx, s.HttpClient, transport)
	if err != nil {
		logger.Error(ctx, "failed http request", map[string]interface{}{
			logger.ErrorKey: err,
			"url":           transport.Url,
			"method":        transport.Method,
		})
		return err
	}
	defer httpRes.Body.Close()

	if httpRes.StatusCode/100 != 2 {
		var res http_utils.ClientResponse
		_ = utils.ParseResponseBody(httpRes.Body, &res)

		logger.Error(ctx, "error http response", map[string]interface{}{
			"status_code": httpRes.StatusCode,
			"response":    res,
			"url":         transport.Url,
			"method":      transport.Method,
		})
		return ErrorInternalLocationService
	}

	return nil
}

func (s *locationRepo) GetNearestAvailableDrivers(ctx context.Context, location model.Coordinate) ([]int64, error) {
	req := model.GetNearestAvailableDriversRequest{
		Location: location,
	}
	transport := http_client.RestTransport{
		Url:    s.BaseUrl + getNearestAvailableDrivers,
		Method: http.MethodGet,
		Header: http.Header{
			http_utils.HeaderApiKey: []string{s.APIKey},
		},
		Payload: req,
	}

	httpRes, err := http_client.Request(ctx, s.HttpClient, transport)
	if err != nil {
		logger.Error(ctx, "failed http request", map[string]interface{}{
			logger.ErrorKey: err,
			"url":           transport.Url,
			"method":        transport.Method,
			"payload":       req,
		})
		return nil, err
	}
	defer httpRes.Body.Close()

	var res http_utils.ClientResponse
	_ = utils.ParseResponseBody(httpRes.Body, &res)

	if httpRes.StatusCode/100 != 2 {
		logger.Error(ctx, "error http response", map[string]interface{}{
			"status_code": httpRes.StatusCode,
			"response":    res,
			"url":         transport.Url,
			"method":      transport.Method,
			"payload":     req,
		})
		return nil, ErrorInternalLocationService
	}

	var data model.GetNearestAvailableDriversResponse
	err = json.Unmarshal(res.Data, &data)
	if err != nil {
		logger.Error(ctx, "error parsing response data", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return data.DriverIDs, nil
}

func (s *locationRepo) GetRidePath(ctx context.Context, rideID int64, driverID int64) ([]model.Coordinate, error) {
	req := model.GetRidePathRequest{
		RideID:   rideID,
		DriverID: driverID,
	}
	transport := http_client.RestTransport{
		Url:    s.BaseUrl + getRidePath,
		Method: http.MethodGet,
		Header: http.Header{
			http_utils.HeaderApiKey: []string{s.APIKey},
		},
		Payload: req,
	}

	httpRes, err := http_client.Request(ctx, s.HttpClient, transport)
	if err != nil {
		logger.Error(ctx, "failed http request", map[string]interface{}{
			logger.ErrorKey: err,
			"url":           transport.Url,
			"method":        transport.Method,
			"payload":       req,
		})
		return nil, err
	}
	defer httpRes.Body.Close()

	var res http_utils.ClientResponse
	_ = utils.ParseResponseBody(httpRes.Body, &res)

	if httpRes.StatusCode/100 != 2 {
		logger.Error(ctx, "error http response", map[string]interface{}{
			"status_code": httpRes.StatusCode,
			"response":    res,
			"url":         transport.Url,
			"method":      transport.Method,
			"payload":     req,
		})
		return nil, ErrorInternalLocationService
	}

	var data model.GetRidePathResponse
	err = json.Unmarshal(res.Data, &data)
	if err != nil {
		logger.Error(ctx, "error parsing response data", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}

	return data.Path, nil
}
