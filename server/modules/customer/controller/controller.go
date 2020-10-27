package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rogaha/go-postgres-jwt-react-starter/server/errors"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/modules/customer/repository"
	customer_viewmodel "github.com/rogaha/go-postgres-jwt-react-starter/server/modules/customer/viewmodel"
	payment_repository "github.com/rogaha/go-postgres-jwt-react-starter/server/modules/payment/repository"
)

//CreateOrUpdate new customer
func CreateOrUpdate(c *gin.Context) {
	var customer customer_viewmodel.Customer
	c.Bind(&customer)

	valErr := customer.ValidateRegister(errors.ValidationErrors)

	if len(valErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
		return
	}

	ok := repository.CreateOrUpdate(customer)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
	return
}

//GetAll a
func GetAll(c *gin.Context) {
	customers := repository.GetAll()

	c.JSON(http.StatusOK, gin.H{"success": true, "customers": customers})
	return
}

//GetByID a
func GetByID(c *gin.Context) {
	id := c.Params.ByName("id")

	customer := repository.GetByID(id)

	c.JSON(http.StatusOK, gin.H{"success": true, "customer": customer})
	return
}

//Delete a
func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	repository.Delete(id)

	c.JSON(http.StatusOK, gin.H{"success": true})
	return
}

//GetAllByCustomer a
func GetAllByCustomer(c *gin.Context) {
	id := c.Params.ByName("id")

	payments := payment_repository.GetAllByCustomer(id)

	c.JSON(http.StatusOK, gin.H{"success": true, "payments": payments})
	return
}

//GetCountAllByCustomer a
func GetCountAllByCustomer(c *gin.Context) {
	id := c.Params.ByName("id")

	payments := payment_repository.GetCountAllByCustomer(id)

	c.JSON(http.StatusOK, gin.H{"success": true, "payments": payments})
	return
}
