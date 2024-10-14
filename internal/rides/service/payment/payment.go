package payment

import "context"

func NewPaymentService() PaymentService {
	return &paymentSvc{}
}

func (s *paymentSvc) AddCredit(ctx context.Context, req AddCreditRequest) error {
	return nil
}

func (s *paymentSvc) DeductCredit(ctx context.Context, req DeductCreditRequest) error {
	return nil
}

func (s *paymentSvc) AddRevenue(ctx context.Context, req AddRevenueRequest) error {
	return nil
}
