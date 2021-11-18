package main

import (
	"fmt"
	"net/http"
	"voting-app/internal/routes"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Started voting app backend!")
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("../client/build", true)))
	router.Use(cors.Default())
	// router.Use(CORSMiddleware())
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	routes.SetupRouter(router)
	router.Run(":8080")
}

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }
