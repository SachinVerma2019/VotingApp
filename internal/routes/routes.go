package routes

import (
	"voting-app/internal/api"
	"voting-app/internal/middleware"
	repo "voting-app/internal/platform/repository"

	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter(router *gin.Engine) {

	dbConnection := repo.NewEntClient()
	router.POST("/user/login", api.AuthenticateUser(&dbConnection))
	router.POST("/user/register", api.RegisterUser(&dbConnection))

	// router.POST("/poll/", api.Poll(&dbConnection))
	// router.POST("/poll/vote", api.Vote(&dbConnection))
	// router.GET("/poll/:voterId", api.Polls_V2(&dbConnection))
	router.GET("/poll/record/:pollId/:option/", api.GetAllUsersVotedOption(&dbConnection))

	//Routes with JWT authorisation
	router.POST("/poll/", middleware.AuthorizeJWT(), api.Poll(&dbConnection))
	router.POST("/poll/vote", middleware.AuthorizeJWT(), api.Vote(&dbConnection))
	router.GET("/poll/:voterId", middleware.AuthorizeJWT(), api.Polls_V2(&dbConnection))

}
