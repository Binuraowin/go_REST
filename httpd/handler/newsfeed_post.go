package handler


import (
	"go_REST/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsFeedPost(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context){
		c.JSON(http.StatusOK, map[string]string{
			"hello": "Binura",
		})
	}

}