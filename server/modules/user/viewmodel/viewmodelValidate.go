package viewmodel

import "github.com/rogaha/go-postgres-jwt-react-starter/server/utils"

// ValidateRegister returns a slice of string of validation errors
func (user Register) ValidateRegister(erros []string) []string {
	if err, message := utils.ValidateEmail(user.Email); err {
		erros = append(erros, message)
	}

	if err, message := utils.ValidateStrMinLength(user.Name, "Name", 1); err {
		erros = append(erros, message)
	}

	if err, message := utils.ValidateStrMinLength(user.Password, "Password", 4); err {
		erros = append(erros, message)
	}

	return erros
}
