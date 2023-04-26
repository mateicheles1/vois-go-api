package middlewares

import (
	"gogin-api/logs"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		for _, err := range ctx.Errors {
			logs.ErrorLogger.Error().
				Str("Method", ctx.Request.Method).
				Str("Path", ctx.Request.URL.Path).
				Int("Status code", ctx.Writer.Status()).
				Msgf("JSON syntax error in request body: %s", err.Error())
		}

		if len(ctx.Errors) > 0 {
			ctx.JSON(ctx.Writer.Status(), "Error processing request: invalid JSON syntax in request body")
		}

	}
}
