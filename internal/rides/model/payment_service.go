package model

type AddCreditRequest struct {
	MSISDN string  `json:"msisdn"`
	Value  float64 `json:"value"`
}

type DeductCreditRequest struct {
	MSISDN string  `json:"msisdn"`
	Value  float64 `json:"value"`
}
