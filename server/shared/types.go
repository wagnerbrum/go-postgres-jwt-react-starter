package shared

import (
	"github.com/gin-gonic/gin"
)

//Route struct
type Route struct {
	Method      string
	Pattern     string
	HandlerFunc func(*gin.Context)
	UseAuth     bool
}
