package middlewares

import (
	"net/http"
	"strings"
	"time"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := "4etjkdfk?jse07ghf8ffper9fd!@78&L!"

		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			c.JSON(http.StatusUnauthorized, myerror.ErrorApp{
				Code:      http.StatusUnauthorized,
				Message:   "precisa fazer login para acessar",
				Timestamp: time.Now().Format(time.RFC3339),
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authorization, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, myerror.ErrorApp{
				Code:      http.StatusUnauthorized,
				Message:   "token inválido",
				Timestamp: time.Now().Format(time.RFC3339),
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return
		}

		userID := claims["id"].(float64)

		userIDUint := uint(userID)

		c.Set("userID", userIDUint)
		c.Next()

	}
}
