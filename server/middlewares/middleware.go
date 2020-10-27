// Package middlewares contains gin middlewares
// Usage: router.Use(middlewares.Connect)
package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/rogaha/go-postgres-jwt-react-starter/server/config"
	user_viewmodel "github.com/rogaha/go-postgres-jwt-react-starter/server/modules/user/viewmodel"
)

// ErrorHandler is a middleware to handle errors encountered during requests
func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors,
		})
	}
}

// CORSMiddleware //
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.ClientUrl)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

//AuthMiddleware a
func AuthMiddleware(useAuth bool) gin.HandlerFunc {
	if !useAuth {
		return func(c *gin.Context) {
			c.Next()
			return
		}
	}

	return func(c *gin.Context) {
		var jwtKey = []byte(config.JwtKey)

		//obtain session token from the requests cookies
		ck, err := c.Request.Cookie("token")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "unauthorized"})
			return
		}

		// Get the JWT string from the cookie
		tokenString := ck.Value

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return jwtKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "msg": err})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var userAuth = user_viewmodel.User{
				Name:      claims["name"].(string),
				Email:     claims["email"].(string),
				CreatedAt: claims["created_at"].(string),
				UpdatedAt: claims["updated_at"].(string),
			}

			c.Set("user_auth", userAuth)
			c.Next()

			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "unauthorized"})
	}
}
