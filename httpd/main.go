package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joshjwelsh/newsfeeder/httpd/handler"
	"github.com/joshjwelsh/newsfeeder/platform/newsfeed"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// feed := newsfeed.New()
	// fmt.Println(feed)
	// feed.Add(newsfeed.Item{"Breaking news", "Confirmed! The chicken came before the egg..."})
	// fmt.Println(feed)

}
