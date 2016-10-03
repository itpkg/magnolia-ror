package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Redirect redirect handler
func Redirect(fn func(*gin.Context) (string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if val, err := fn(c); err == nil {
			c.Redirect(http.StatusTemporaryRedirect, val)
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	}
}

//JSON json handler
func JSON(fn func(*gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if val, err := fn(c); err == nil {
			c.JSON(http.StatusOK, val)
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	}
}
