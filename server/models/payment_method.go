package models

//PaymentMethod a
type PaymentMethod struct {
	ID                     string
	CustomerID             string
	InstrumentID           string
	MethodType             string
	CardBin                string
	CardLastFour           string
	ExpiryMonth            int16
	ExpiryYear             int16
	EWallet                string
	NameOnCard             string
	SuccessfulRegistration bool
	RegistrationTime       int64
	LastVerified           int64
	Customer               Customer
	PaymentBillingAddress  PaymentBillingAddress `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt              int64                 `gorm:"autoCreateTime"`       // Use unix seconds as creating time
	UpdatedAt              int64                 `gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
}
