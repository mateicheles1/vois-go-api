package middlewares

import (
	"gogin-api/logs"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {

	c.Next()
	
	for _, err := range c.Errors {
		logs.ErrorLogger.Error().
			Str("Method", c.Request.Method).
			Str("Path", c.Request.URL.Path).
			Int("Status code", c.Writer.Status()).
			Msgf("Could not bind the request body to desired struct due to: %s", err.Error())
	}

	if len(c.Errors) != 0 {
		c.JSON(-1, "error processing request; invalid syntax")
	}
}
