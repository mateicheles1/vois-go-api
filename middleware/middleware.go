package middleware

import (
	"gogin-api/logs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		logs.Logger.Error().
			Str("Method", c.Request.Method).
			Str("Path", c.Request.URL.Path).
			Int("Status code", http.StatusBadRequest).
			Msgf("Could not unmarshal the request body into the requestBody struct due to: %s", err.Error())
	}

	if len(c.Errors) != 0 {
		c.JSON(-1, "error processing request; invalid syntax")
	}
}
