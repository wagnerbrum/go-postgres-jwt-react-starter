package viewmodel

//Customer a
type Customer struct {
	CustomerID            string           `json:"customerId"`
	Email                 string           `json:"email"`
	EmailVerifiedTime     int64            `json:"emailVerifiedTime"`
	Name                  string           `json:"name"`
	FamilyName            string           `json:"familyName"`
	GivenName             string           `json:"givenName"`
	Telephone             string           `json:"telephone"`
	TelephoneVerifiedTime int64            `json:"telephoneVerifiedTime"`
	TelephoneCountry      string           `json:"telephoneCountry"`
	CreatedAt             int64            `json:"registrationTime"` // Use unix seconds as creating time
	CustomerLocation      CustomerLocation `json:"location"`
}

//CustomerLocation a
type CustomerLocation struct {
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
