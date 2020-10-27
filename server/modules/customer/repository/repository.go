package repository

import (
	"github.com/rogaha/go-postgres-jwt-react-starter/server/db"
	models "github.com/rogaha/go-postgres-jwt-react-starter/server/models"
	customer_viewmodel "github.com/rogaha/go-postgres-jwt-react-starter/server/modules/customer/viewmodel"
)

//CreateOrUpdate a
func CreateOrUpdate(customer customer_viewmodel.Customer) bool {
	customerDB := &models.Customer{
		ID:                    customer.CustomerID,
		Email:                 customer.Email,
		EmailVerifiedTime:     customer.EmailVerifiedTime,
		Name:                  customer.Name,
		FamilyName:            customer.FamilyName,
		GivenName:             customer.GivenName,
		Telephone:             customer.Telephone,
		TelephoneVerifiedTime: customer.TelephoneVerifiedTime,
		TelephoneCountry:      customer.TelephoneCountry,
	}

	customerLocationDB := &models.CustomerLocation{
		CustomerID:    customer.CustomerID,
		Country:       customer.CustomerLocation.Country,
		PostalCode:    customer.CustomerLocation.PostalCode,
		Latitude:      customer.CustomerLocation.Latitude,
		Longitude:     customer.CustomerLocation.Longitude,
		AddresseeName: customer.CustomerLocation.AddresseeName,
		Street1:       customer.CustomerLocation.Street1,
		Street2:       customer.CustomerLocation.Street2,
		Neighbourhood: customer.CustomerLocation.Neighbourhood,
		Zone:          customer.CustomerLocation.Zone,
		City:          customer.CustomerLocation.City,
		Region:        customer.CustomerLocation.Region,
		PoBoxNumber:   customer.CustomerLocation.PoBoxNumber,
	}

	// Insert Customer
	result := db.DB.Model(&customerDB).Where("id = ?", customer.CustomerID).Updates(&customerDB)

	if result.RowsAffected == 0 {
		result = db.DB.Create(&customerDB)
	}

	if result.Error != nil {
		return false
	}

	// Insert CustomerLocation
	result = db.DB.Model(&customerLocationDB).Where("customer_id = ?", customerLocationDB.CustomerID).Updates(&customerLocationDB)

	if result.RowsAffected == 0 {
		result = db.DB.Create(&customerLocationDB)
	}

	if result.Error != nil {
		return false
	}

	return true
}

//GetAll a
func GetAll() *[]models.Customer {
	var customerDB []models.Customer

	result := db.DB.Joins("CustomerLocation").Find(&customerDB)

	if result.RowsAffected == 0 || result.Error != nil {
		return nil
	}

	return &customerDB
}

//GetByID a
func GetByID(customerID string) *models.Customer {
	var customerDB models.Customer

	result := db.DB.Joins("CustomerLocation").First(&customerDB, `"customers"."id" = ?`, customerID)

	if result.RowsAffected == 0 || result.Error != nil {
		return nil
	}

	return &customerDB
}

//Delete a
func Delete(customerID string) bool {
	result := db.DB.Table("customers").Delete(&models.Customer{}, `id = ?`, customerID)

	if result.RowsAffected == 0 || result.Error != nil {
		return false
	}

	return true
}
