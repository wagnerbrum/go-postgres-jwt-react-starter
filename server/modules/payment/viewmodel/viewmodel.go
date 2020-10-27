package viewmodel

import (
	customer_viewmodel "github.com/rogaha/go-postgres-jwt-react-starter/server/modules/customer/viewmodel"
)

//PaymentMethod a
type PaymentMethod struct {
	ID                     string                      `json:"paymentMethodID"`
	CustomerID             string                      `json:"customerID"`
	InstrumentID           string                      `json:"instrumentID"`
	MethodType             string                      `json:"methodType"`
	CardBin                string                      `json:"cardBin"`
	CardLastFour           string                      `json:"cardLastFour"`
	ExpiryMonth            int16                       `json:"expiryMonth"`
	ExpiryYear             int16                       `json:"expiryYear"`
	EWallet                string                      `json:"eWallet"`
	NameOnCard             string                      `json:"nameOnCard"`
	SuccessfulRegistration bool                        `json:"successfulRegistration"`
	RegistrationTime       int64                       `json:"registrationTime"`
	LastVerified           int64                       `json:"lastVerified"`
	PaymentBillingAddress  PaymentBillingAddress       `json:"billingAddress"`
	Customer               customer_viewmodel.Customer `json:"customer"`
}

//PaymentBillingAddress a
type PaymentBillingAddress struct {
	Country       string  `json:"country"`
	PostalCode    string  `json:"postalCode"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	AddresseeName string  `json:"addresseeName"`
	Street1       string  `json:"street1"`
	Street2       string  `json:"street2"`
	Neighbourhood string  `json:"neighbourhood"`
	Zone          string  `json:"zone"`
	City          string  `json:"city"`
	Region        string  `json:"region"`
	PoBoxNumber   string  `json:"poBoxNumber"`
}
