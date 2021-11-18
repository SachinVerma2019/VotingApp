package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	authz "voting-app/internal/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthorizeJWT ... API security with JWT authorisation
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("Entered authorizeJWT")
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimSpace(authHeader[len(BEARER_SCHEMA):])
		//log.Print("tokenString :", tokenString)
		token, err := authz.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
