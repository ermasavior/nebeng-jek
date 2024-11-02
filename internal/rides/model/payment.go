package model

import httpUtils "nebeng-jek/pkg/http/utils"

type AddCreditRequest struct {
	MSISDN string  `json:"msisdn"`
	Value  float64 `json:"value"`
}

type DeductCreditRequest struct {
	MSISDN string  `json:"msisdn"`
	Value  float64 `json:"value"`
}

type PaymentResponse struct {
	Meta httpUtils.MetaResponse `json:"meta"`
}
