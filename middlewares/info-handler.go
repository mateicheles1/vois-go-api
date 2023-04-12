package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InfoHandler() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] - %s %s %d \n",
		param.TimeStamp.Format(time.RFC822),
		param.Method,
		param.Path,
		param.StatusCode,
	)
	})
}
