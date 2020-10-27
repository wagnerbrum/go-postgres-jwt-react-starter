package controller

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/rogaha/go-postgres-jwt-react-starter/server/config"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/errors"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/modules/user/repository"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/modules/user/viewmodel"
	user_viewmodel "github.com/rogaha/go-postgres-jwt-react-starter/server/modules/user/viewmodel"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/shared"
)

var jwtKey = []byte(config.JwtKey)

//Claims jwt claims struct
type Claims struct {
	user_viewmodel.User
	jwt.StandardClaims
}

//Create new user
func Create(c *gin.Context) {
	var user user_viewmodel.Register
	c.Bind(&user)

	valErr := user.ValidateRegister(errors.ValidationErrors)

	_, userExists, _ := repository.GetByEmail(user.Email)

	if userExists {
		valErr = append(valErr, "email already exists")
	}

	if len(valErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
		return
	}

	okCreate := repository.Create(user)

	if !okCreate {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true})
	return
}

// Session returns JSON of user info
func Session(c *gin.Context) {
	user, err := shared.GetUserAuth(c)

	if err {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "user": user})
}

// Login controller
func Login(c *gin.Context) {
	var user user_viewmodel.Login
	c.Bind(&user)

	err, userExists, userDB := repository.GetByEmail(user.Email)

	if err || !userExists || !userDB.CheckPasswordHash(user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "incorrect credentials"})
		return
	}

	expirationTime := time.Now().Add(30 * time.Minute)

	// Create the JWT claims, which includes the User struct and expiry time
	claims := &Claims{
		User: viewmodel.User{
			Name:  userDB.Name,
			Email: userDB.Email,
		},
		StandardClaims: jwt.StandardClaims{
			//expiry time, expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	jwtKey := []byte(config.JwtKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT token string
	tokenString, tokenErr := token.SignedString(jwtKey)
	errors.HandleErr(c, tokenErr)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "logged in succesfully", "user": claims.User, "token": tokenString})
	return
}
