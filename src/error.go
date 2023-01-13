package src

import (
	"github.com/gin-gonic/gin"
)

// Error function is used for the error handling throughout the code
func Error(err error) gin.H {
	return gin.H{
		"statuscode": 400,
		"error":      err.Error(),
	}
}
