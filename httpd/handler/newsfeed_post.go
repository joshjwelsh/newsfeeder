package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshjwelsh/newsfeeder/platform/newsfeed"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(feed *newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsfeedPostRequest{}
		c.Bind(&requestBody)
		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)
		c.Status(http.StatusNoContent) // 204 Response - no body but all is well
	}

}
