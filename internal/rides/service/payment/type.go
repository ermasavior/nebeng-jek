package payment

import (
	"context"
	"net/http"
)

type PaymentService interface {
	AddCredit(context.Context, AddCreditRequest) error
	DeductCredit(context.Context, DeductCreditRequest) error
	AddRevenue(context.Context, AddRevenueRequest) error
}

type paymentSvc struct {
	HttpClient *http.Client
	BaseUrl    string
	APIKey     string
}
