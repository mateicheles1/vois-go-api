package routes

import (
	"gogin-api/logs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func check(err error, c *gin.Context) {
	logs.Logger.Panic().
		Str("Method", c.Request.Method).
		Str("Path", c.Request.URL.Path).
		Int("Status code", http.StatusBadRequest).
		Msgf("Could not unmarshal the request body into the requestBody struct due to: %s", err)
}
