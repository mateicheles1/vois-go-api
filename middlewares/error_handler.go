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

			switch ctx.Writer.Status() {

			case http.StatusBadRequest:
				logs.ErrorLogger.Error().
					Str("Method", ctx.Request.Method).
					Str("Path", ctx.Request.URL.Path).
					Int("Status code", http.StatusBadRequest).
					Msgf("Bad request: %s", err)

				ctx.JSON(http.StatusBadRequest, err.Error())

			case http.StatusInternalServerError:
				logs.ErrorLogger.Error().
					Str("Method", ctx.Request.Method).
					Str("Path", ctx.Request.URL.Path).
					Int("Status code", http.StatusInternalServerError).
					Msgf("Internal server error: %s", err)

				ctx.JSON(http.StatusInternalServerError, "something went wrong")

			case http.StatusUnauthorized:
				logs.ErrorLogger.Error().
					Str("Method", ctx.Request.Method).
					Str("Path", ctx.Request.URL.Path).
					Int("Status code", http.StatusUnauthorized).
					Msgf("Unauthorized: %s", err)

				ctx.JSON(http.StatusUnauthorized, "invalid credentials")

			case http.StatusForbidden:
				logs.ErrorLogger.Error().
					Str("Method", ctx.Request.Method).
					Str("Path", ctx.Request.URL.Path).
					Int("Status code", http.StatusForbidden).
					Msgf("Forbidden: %s", err)

				ctx.JSON(http.StatusForbidden, "action not allowed")
			}
		}
	}
}
