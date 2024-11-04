package payment

import (
	"context"
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
	deductCreditPath = "/v1/deduct/credit"
	addCreditPath    = "/v1/add/credit"
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

	_, err := http_client.RequestHTTPAndParseResponse(ctx, s.HttpClient, transport)
	if err != nil {
		logger.Error(ctx, "error request http", map[string]interface{}{
			logger.ErrorKey: err,
		})
		return err
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

	_, err := http_client.RequestHTTPAndParseResponse(ctx, s.HttpClient, transport)
	if err != nil {
		logger.Error(ctx, "error request http", map[string]interface{}{
			logger.ErrorKey: err,
		})
		return err
	}

	return nil
}
