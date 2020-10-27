package viewmodel

import "github.com/rogaha/go-postgres-jwt-react-starter/server/utils"

// ValidateRegister returns a slice of string of validation errors
func (customer Customer) ValidateRegister(erros []string) []string {
	if err, message := utils.ValidateStrMinLength(customer.CustomerID, "CustomerID", 1); err {
		erros = append(erros, message)
	}

	return erros
}
