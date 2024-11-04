package payment

import (
	"context"
	"errors"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/internal/rides/repository"
	"nebeng-jek/pkg/configs"
	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/http_client"
	"nebeng-jek/pkg/logger"
	"nebeng-jek/pkg/utils"
	"net/http"

	"github.com/google/uuid"
)

const (
	deductCreditPath = "/v1/deduct/credit"
	addCreditPath    = "/v1/add/credit"
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
		Url:    s.BaseUrl + addCreditPath,
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
		var res httpUtils.Response
		_ = utils.ParseResponseBody(httpRes.Body, &res)

		logger.Error(ctx, "error http response", map[string]interface{}{
			"status_code": httpRes.StatusCode,
			"response":    res,
		})
		return ErrorInternalPaymentService
	}

	return nil
}

func (s *paymentRepo) DeductCredit(ctx context.Context, req model.DeductCreditRequest) error {
	idempotencyKey := uuid.New().String()
	transport := http_client.RestTransport{
		Url:    s.BaseUrl + deductCreditPath,
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
		var res httpUtils.Response
		_ = utils.ParseResponseBody(httpRes.Body, &res)

		logger.Error(ctx, "error http response", map[string]interface{}{
			"status_code": httpRes.StatusCode,
			"response":    res,
		})
		return ErrorInternalPaymentService
	}

	return nil
}
