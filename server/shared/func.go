package shared

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"

	user_viewmodel "github.com/rogaha/go-postgres-jwt-react-starter/server/modules/user/viewmodel"
)

// GetUserAuth return Object of user info
func GetUserAuth(c *gin.Context) (user *user_viewmodel.User, err bool) {
	if userAuthMap, ok := c.Keys["user_auth"]; ok {
		var userAuth *user_viewmodel.User

		mapstructure.Decode(userAuthMap, &userAuth)

		return userAuth, false
	}

	return nil, true
}
