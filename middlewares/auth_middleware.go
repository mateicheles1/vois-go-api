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
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("no authorization header provided"))
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid authorization header"))
			return
		}

		tokenVar := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenVar, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		ctx.Set("username", claims["username"])
		ctx.Set("role", claims["role"])
		ctx.Next()
	}
}
