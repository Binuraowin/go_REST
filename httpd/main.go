package main

import (
	"go_REST/httpd/handler"
"go_REST/platform/newsfeed"

"github.com/gin-gonic/gin"

)

func main() {
	feed := newsfeed.New()
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// feed := newsfeed.New()

	// fmt.Println(feed)
	// feed.Add(newsfeed.Item{"Hello","how are you"})
	// fmt.Println(feed)
}
