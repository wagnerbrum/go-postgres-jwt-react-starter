package repository

import (
	"fmt"

	"github.com/rogaha/go-postgres-jwt-react-starter/server/db"
	models "github.com/rogaha/go-postgres-jwt-react-starter/server/models"
	payment_viewmodel "github.com/rogaha/go-postgres-jwt-react-starter/server/modules/payment/viewmodel"
)

//CreateOrUpdate a
func CreateOrUpdate(paymentMethod payment_viewmodel.PaymentMethod) bool {
	fmt.Println("repository - Teste 1")

	paymentMethodDB := &models.PaymentMethod{
		ID:                     paymentMethod.ID,
		CustomerID:             paymentMethod.CustomerID,
		InstrumentID:           paymentMethod.InstrumentID,
		MethodType:             paymentMethod.MethodType,
		CardBin:                paymentMethod.CardBin,
		CardLastFour:           paymentMethod.CardLastFour,
		ExpiryMonth:            paymentMethod.ExpiryMonth,
		ExpiryYear:             paymentMethod.ExpiryYear,
		EWallet:                paymentMethod.EWallet,
		NameOnCard:             paymentMethod.NameOnCard,
		SuccessfulRegistration: paymentMethod.SuccessfulRegistration,
		RegistrationTime:       paymentMethod.RegistrationTime,
		LastVerified:           paymentMethod.LastVerified,
	}
	fmt.Println("repository - Teste 2")

	paymentBillingAddressDB := &models.PaymentBillingAddress{
		PaymentMethodID: paymentMethod.ID,
		Country:         paymentMethod.PaymentBillingAddress.Country,
		PostalCode:      paymentMethod.PaymentBillingAddress.PostalCode,
		Latitude:        paymentMethod.PaymentBillingAddress.Latitude,
		Longitude:       paymentMethod.PaymentBillingAddress.Longitude,
		AddresseeName:   paymentMethod.PaymentBillingAddress.AddresseeName,
		Street1:         paymentMethod.PaymentBillingAddress.Street1,
		Street2:         paymentMethod.PaymentBillingAddress.Street2,
		Neighbourhood:   paymentMethod.PaymentBillingAddress.Neighbourhood,
		Zone:            paymentMethod.PaymentBillingAddress.Zone,
		City:            paymentMethod.PaymentBillingAddress.City,
		Region:          paymentMethod.PaymentBillingAddress.Region,
		PoBoxNumber:     paymentMethod.PaymentBillingAddress.PoBoxNumber,
	}
	// Insert PaymentMethod
	result := db.DB.Model(&paymentMethodDB).Where("id = ?", paymentMethod.ID).Updates(&paymentMethodDB)

	if result.RowsAffected == 0 {
		result = db.DB.Create(&paymentMethodDB)
	}

	if result.Error != nil {
		return false
	}

	// Insert PaymentMethodLocation
	result = db.DB.Model(&paymentBillingAddressDB).Where("payment_method_id = ?", paymentMethod.ID).Updates(&paymentBillingAddressDB)

	if result.RowsAffected == 0 {
		result = db.DB.Create(&paymentBillingAddressDB)
	}

	if result.Error != nil {
		return false
	}

	return true
}

//GetAll a
func GetAll() *[]models.PaymentMethod {
	var paymentMethodDB []models.PaymentMethod

	result := db.DB.Find(&paymentMethodDB)

	if result.RowsAffected == 0 || result.Error != nil {
		return nil
	}

	return &paymentMethodDB
}

//GetByID a
func GetByID(paymentMethodID string) *models.PaymentMethod {
	var paymentMethodDB models.PaymentMethod

	result := db.DB.First(&paymentMethodDB, `"payment_methods"."id" = ?`, paymentMethodID)

	if result.RowsAffected == 0 || result.Error != nil {
		return nil
	}

	return &paymentMethodDB
}

//Delete a
func Delete(paymentMethodID string) bool {
	result := db.DB.Table("payment_methods").Delete(&models.PaymentMethod{}, `id = ?`, paymentMethodID)

	if result.RowsAffected == 0 || result.Error != nil {
		return false
	}

	return true
}

//GetAllByCustomer a
func GetAllByCustomer(customerID string) *[]models.PaymentMethod {
	var paymentMethodDB []models.PaymentMethod

	result := db.DB.Find(&paymentMethodDB, `"payment_methods"."customer_id" = ?`, customerID)

	if result.RowsAffected == 0 || result.Error != nil {
		return nil
	}

	return &paymentMethodDB
}

//GetCountAllByCustomer a
func GetCountAllByCustomer(customerID string) int {
	var paymentMethodDB []models.PaymentMethod

	result := db.DB.Find(&paymentMethodDB, `"payment_methods"."customer_id" = ?`, customerID)

	if result.RowsAffected == 0 || result.Error != nil {
		return 0
	}

	return len(paymentMethodDB)
}

// get a count of all payment methods for a customer
