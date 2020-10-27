package models

//PaymentBillingAddress a
type PaymentBillingAddress struct {
	ID              int `gorm:"primaryKey"`
	PaymentMethodID string
	Country         string `gorm:"size:100"`
	PostalCode      string `gorm:"size:100"`
	Latitude        float64
	Longitude       float64
	AddresseeName   string `gorm:"size:100"`
	Street1         string `gorm:"size:100"`
	Street2         string `gorm:"size:100"`
	Neighbourhood   string `gorm:"size:100"`
	Zone            string `gorm:"size:100"`
	City            string `gorm:"size:100"`
	Region          string `gorm:"size:100"`
	PoBoxNumber     string `gorm:"size:100"`
	CreatedAt       int64  `gorm:"autoCreateTime"`       // Use unix seconds as creating time
	UpdatedAt       int64  `gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
}
