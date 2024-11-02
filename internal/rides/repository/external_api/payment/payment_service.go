package payment

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/internal/rides/repository"
	"nebeng-jek/pkg/configs"
	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/http_client"
	"nebeng-jek/pkg/logger"
	"net/http"

	"github.com/google/uuid"
)

const (
	DeductCreditPath = "/v1/deduct/credit"
	AddCreditPath    = "/v1/add/credit"
)

var (
	ErrorInternalPaymentService = errors.New("failed http request to payment service")
)

type paymentRepo struct {
	HttpClient *http.Client
	BaseUrl    string
	APIKey     string
}

func NewPaymentRepository(cfg *configs.Config, httpClient *http.Client) repository.PaymentRepository {
	return &paymentRepo{
		HttpClient: httpClient,
		BaseUrl:    cfg.PaymentServiceURL,
		APIKey:     cfg.PaymentServiceAPIKey,
	}
}

func (s *paymentRepo) AddCredit(ctx context.Context, req model.AddCreditRequest) error {
	idempotencyKey := uuid.New().String()
	transport := http_client.RestTransport{
		Url:    s.BaseUrl + AddCreditPath,
		Method: http.MethodPost,
		Header: http.Header{
			httpUtils.HeaderApiKey:   []string{s.APIKey},
			httpUtils.IdempotencyKey: []string{idempotencyKey},
		},
		Payload: req,
	}

	httpRes, err := http_client.Request(ctx, s.HttpClient, transport)
	if err != nil {
		logger.Error(ctx, "failed http request", map[string]interface{}{
			logger.ErrorKey: err,
		})
		return err
	}
	defer httpRes.Body.Close()

	if httpRes.StatusCode/100 != 2 {
		logger.Error(ctx, "error http response", map[string]interface{}{
			"status_code": httpRes.StatusCode,
			"response":    parseResponseBody(httpRes.Body),
		})
		return ErrorInternalPaymentService
	}

	return nil
}

func (s *paymentRepo) DeductCredit(ctx context.Context, req model.DeductCreditRequest) error {
	idempotencyKey := uuid.New().String()
	transport := http_client.RestTransport{
		Url:    s.BaseUrl + DeductCreditPath,
		Method: http.MethodPost,
		Header: http.Header{
			httpUtils.HeaderApiKey:   []string{s.APIKey},
			httpUtils.IdempotencyKey: []string{idempotencyKey},
		},
		Payload: req,
	}

	httpRes, err := http_client.Request(ctx, s.HttpClient, transport)
	if err != nil {
		logger.Error(ctx, "failed http request", map[string]interface{}{
			logger.ErrorKey: err,
		})
		return err
	}
	defer httpRes.Body.Close()

	if httpRes.StatusCode/100 != 2 {
		logger.Error(ctx, "error http response", map[string]interface{}{
			"status_code": httpRes.StatusCode,
			"response":    parseResponseBody(httpRes.Body),
		})
		return ErrorInternalPaymentService
	}

	return nil
}

func parseResponseBody(resBody io.ReadCloser) model.PaymentResponse {
	var res model.PaymentResponse
	b, _ := io.ReadAll(resBody)
	_ = json.Unmarshal(b, &res)
	return res
}
