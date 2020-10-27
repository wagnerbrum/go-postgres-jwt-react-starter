package models

//Customer a
type Customer struct {
	ID                    string
	CustomerLocation      CustomerLocation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Email                 string
	EmailVerifiedTime     int64
	Name                  string
	FamilyName            string
	GivenName             string
	Telephone             string
	TelephoneVerifiedTime int64
	TelephoneCountry      string
	CreatedAt             int64 `gorm:"autoCreateTime"`       // Use unix seconds as creating time
	UpdatedAt             int64 `gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
}
