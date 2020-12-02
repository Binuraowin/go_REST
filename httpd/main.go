package main

import (
	"go_REST/platform/newsfeed"
"fmt"

)

func main() {
	// r := gin.Default()
	// r.GET("/ping", handler.PingGet())
	
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	feed := newsfeed.New()

	fmt.Println(feed)
	feed.Add(newsfeed.Item{"Hello","how are you"})
	fmt.Println(feed)
}
