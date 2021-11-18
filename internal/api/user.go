package api

import (
	"context"
	"fmt"
	"net/http"
	auth "voting-app/internal/auth"
	repo "voting-app/internal/platform/repository"

	"github.com/gin-gonic/gin"
)

func AuthenticateUser(repos *repo.Repositiory) gin.HandlerFunc {
	userAuthentication := func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")
		fmt.Println(email)
		fmt.Println(password)
		ctx := context.Background()
		id, name := (*repos).AuthenticateUser(ctx, email, password)
		if id == -1 {
			c.JSON(http.StatusUnauthorized, gin.H{"Success": false})
			return
		}

		jwtToken, err := auth.JWTAuthService().GenerateToken(email, true)
		if err != nil {
			c.JSON(500, gin.H{"status": 500, "message": "Auth token Not Generated"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"id":     id,
			"name":   name,
			"email":  email,
			"token":  jwtToken,
		})
	}
	return gin.HandlerFunc(userAuthentication)
}

func RegisterUser(repo *repo.Repositiory) gin.HandlerFunc {
	userRegistration := func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		password := c.PostForm("password")
		fmt.Println(name)
		fmt.Println(email)
		fmt.Println(password)
		ctx := context.Background()
		id := (*repo).CreateUser(ctx, name, email, password)
		if id == -1 {
			c.JSON(http.StatusUnauthorized, gin.H{"Success": false})
			return
		}
		jwtToken, err := auth.JWTAuthService().GenerateToken(email, true)
		if err != nil {
			c.JSON(500, gin.H{"status": 500, "message": "Auth token Not Generated"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"id":     id,
			"name":   name,
			"email":  email,
			"token":  jwtToken,
		})
	}
	return gin.HandlerFunc(userRegistration)
}
