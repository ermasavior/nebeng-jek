package payment

type AddCreditRequest struct {
	MSISDN string
	Value  float64
}

type DeductCreditRequest struct {
	MSISDN string
	Value  float64
}

type AddRevenueRequest struct {
	Value float64
}
