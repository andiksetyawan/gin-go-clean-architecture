package middleware

import (
	"gin-go-clean-architecture/model"
	"gin-go-clean-architecture/service"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		splitToken := strings.Split(bearerToken, "Bearer ")

		if len(splitToken) != 2 {
			c.JSON(http.StatusUnauthorized, model.ApiResponse{
				Success: false,
				Message: http.StatusText(http.StatusUnauthorized),
				Data:    gin.H{},
			})
			c.Abort()
			return
		}

		tokenString := splitToken[1]
		token, _ := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("guid", claims["guid"])
			c.Set("email", claims["email"])
			c.Set("admin", claims["admin"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, model.ApiResponse{
				Success: false,
				Message: http.StatusText(http.StatusUnauthorized),
				Data:    gin.H{},
			})
			c.Abort()
			return
		}
	}
}
