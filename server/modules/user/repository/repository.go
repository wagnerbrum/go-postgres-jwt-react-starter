package repository

import (
	"github.com/rogaha/go-postgres-jwt-react-starter/server/db"
	models "github.com/rogaha/go-postgres-jwt-react-starter/server/models"
	user_viewmodel "github.com/rogaha/go-postgres-jwt-react-starter/server/modules/user/viewmodel"
)

//GetByEmail a
func GetByEmail(email string) (err, userExists bool, user *models.User) {
	var userDB models.User

	result := db.DB.First(&userDB, "email = ?", email)

	if result.Error != nil {
		return true, false, nil
	}

	if result.RowsAffected < 1 {
		return false, false, nil
	}

	return false, true, &userDB
}

//Create a
func Create(user user_viewmodel.Register) bool {
	userDB := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	userDB.HashPassword()

	result := db.DB.Create(userDB)

	if result.Error != nil {
		return false
	}

	return true
}
