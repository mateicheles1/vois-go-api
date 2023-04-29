package middlewares

import (
	"gogin-api/logs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		for _, err := range ctx.Errors {

			if ctx.Writer.Status() == http.StatusBadRequest {
				logs.ErrorLogger.Error().
					Str("Method", ctx.Request.Method).
					Str("Path", ctx.Request.URL.Path).
					Int("Status code", http.StatusBadRequest).
					Msgf("Bad request: %s", err.Error())

				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}

			if ctx.Writer.Status() == http.StatusInternalServerError {
				logs.ErrorLogger.Error().
					Str("Method", ctx.Request.Method).
					Str("Path", ctx.Request.URL.Path).
					Int("Status code", http.StatusInternalServerError).
					Msgf("Internal server error: %s", err.Error())

				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}

		}
	}
}
