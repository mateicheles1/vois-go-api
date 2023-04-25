package middlewares

import (
	"gogin-api/logs"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errorOccured := false
		for _, err := range c.Errors {
			logs.ErrorLogger.Error().
				Str("Method", c.Request.Method).
				Str("Path", c.Request.URL.Path).
				Int("Status code", c.Writer.Status()).
				Msgf("JSON syntax error in request body: %s", err.Error())
			errorOccured = true
		}

		if errorOccured {
			c.JSON(c.Writer.Status(), "Error processing request: invalid JSON syntax in request body")
		}
	}
}
