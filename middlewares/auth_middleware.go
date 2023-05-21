package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("authorization header required"))
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.Set("username", claims["username"])
			ctx.Set("role", claims["role"])
			ctx.Next()
		}

		ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))

	}
}
