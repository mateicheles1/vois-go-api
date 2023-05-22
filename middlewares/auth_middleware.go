package middlewares

import (
	"errors"
	"fmt"
	"gogin-api/data"
	"gogin-api/logs"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(db *data.ToDoListDB) gin.HandlerFunc {
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

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok {
				return nil, errors.New("invalid token claims")
			}

			username, ok := claims["username"].(string)

			if !ok {
				return nil, errors.New("invalid username claim")
			}

			user, err := db.FindUserByUsername(username)

			if err != nil {
				return nil, fmt.Errorf("could not find username: %s", err)
			}

			return []byte(user.SecretKey), nil
		})

		if err != nil {
			logs.ErrorLogger.Error().Msgf("Invalid token: %s", err)
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
