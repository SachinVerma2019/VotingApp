package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	repo "voting-app/internal/platform/repository"

	"github.com/gin-gonic/gin"
)

func Poll(repo *repo.Repositiory) gin.HandlerFunc {
	createPoll := func(c *gin.Context) {
		ownerId, _ := strconv.Atoi(c.PostForm("ownerId"))
		topic := c.PostForm("topic")
		options := c.PostForm("options")
		var opt []string
		_ = json.Unmarshal([]byte(options), &opt)
		log.Printf("Unmarshaled: %v", opt)

		fmt.Println(ownerId)
		fmt.Println(topic)
		fmt.Println(options)

		ctx := context.Background()

		res, err := (*repo).CreatePoll(ctx, ownerId, topic, opt)
		if err || res == -1 {
			c.JSON(http.StatusInternalServerError, gin.H{"Success": false})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": 201,
			"pollId": res,
		})
	}
	return gin.HandlerFunc(createPoll)
}

//Vote - function to cast a vote
func Vote(repo *repo.Repositiory) gin.HandlerFunc {
	castVote := func(c *gin.Context) {
		voterId, _ := strconv.Atoi(c.PostForm("voterId"))
		pollId, _ := strconv.Atoi(c.PostForm("pollId"))
		option := c.PostForm("option")

		fmt.Println(voterId)
		fmt.Println(pollId)
		fmt.Println(option)

		ctx := context.Background()
		if !(*repo).ValidateOption(ctx, pollId, option) {
			c.JSON(http.StatusBadRequest, gin.H{"Success": false})
			return
		}
		res, err := (*repo).CastVote(ctx, voterId, pollId, option)
		if err || res == -1 {
			c.JSON(http.StatusInternalServerError, gin.H{"Success": false})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":   201,
			"votingId": res,
		})
	}
	return gin.HandlerFunc(castVote)
}

//Vote - function to cast a vote
func Polls(repo *repo.Repositiory) gin.HandlerFunc {
	allPoles := func(c *gin.Context) {
		ctx := context.Background()
		res := (*repo).GetAllPolls(ctx)
		if res == nil {
			c.JSON(http.StatusNoContent, gin.H{"Success": false})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 201,
			"polls":  res,
		})
	}
	return gin.HandlerFunc(allPoles)
}

//Polls_V2 - function to cast a vote
func Polls_V2(repo *repo.Repositiory) gin.HandlerFunc {
	allPoles := func(c *gin.Context) {
		// userId, _ := strconv.Atoi(c.PostForm("voterId"))
		userId, _ := strconv.Atoi(c.Param("voterId"))
		ctx := context.Background()
		votedSummary := (*repo).GetVotedPolls(ctx, userId)
		if votedSummary == nil {
			fmt.Println("No vote has been casted by the user yet")
		}
		notVotedPolls := (*repo).GetNotVotedPolls(ctx, userId)
		if notVotedPolls == nil {
			fmt.Println("All polls have been voted by the user")
		}
		c.JSON(http.StatusOK, gin.H{
			"status":           201,
			"votedPollSummary": votedSummary,
			"polls":            notVotedPolls,
		})
	}
	return gin.HandlerFunc(allPoles)
}

//Polls_V2 - function to cast a vote
func GetAllUsersVotedOption(repo *repo.Repositiory) gin.HandlerFunc {
	userForVotedOptions := func(c *gin.Context) {
		pollId, _ := strconv.Atoi(c.Param("pollId"))
		option := c.Param("option")
		ctx := context.Background()
		emails := (*repo).GetUsersForVoteOption(ctx, pollId, option)
		c.JSON(http.StatusOK, gin.H{
			"status": 201,
			"emails": emails,
		})
	}
	return gin.HandlerFunc(userForVotedOptions)
}
