package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rogaha/go-postgres-jwt-react-starter/server/errors"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/modules/payment/repository"
	payment_viewmodel "github.com/rogaha/go-postgres-jwt-react-starter/server/modules/payment/viewmodel"
)

//CreateOrUpdate new payment
func CreateOrUpdate(c *gin.Context) {
	var paymentMethod payment_viewmodel.PaymentMethod
	c.Bind(&paymentMethod)

	valErr := paymentMethod.ValidateRegister(errors.ValidationErrors)

	if len(valErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
		return
	}

	ok := repository.CreateOrUpdate(paymentMethod)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
	return
}

//GetAll a
func GetAll(c *gin.Context) {
	paymentMethods := repository.GetAll()

	c.JSON(http.StatusOK, gin.H{"success": true, "payments": paymentMethods})
	return
}

//GetByID a
func GetByID(c *gin.Context) {
	id := c.Params.ByName("id")

	paymentMethod := repository.GetByID(id)

	c.JSON(http.StatusOK, gin.H{"success": true, "payment": paymentMethod})
	return
}

//Delete a
func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	repository.Delete(id)

	c.JSON(http.StatusOK, gin.H{"success": true})
	return
}
