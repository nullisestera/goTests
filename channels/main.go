package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://facebook.com",
		"http://youtube.com",
	}

	// Create a channel in order to communicate go routines
	c := make(chan string)

	for _, link := range links {
		// go routine
		go checkLink(link, c)
	}

	// Receive message from channel
	// Iterate over channel
	for l := range c {
		// link. It can be whatever variable name
		go func(link string) {
			time.Sleep(5 * time.Second) // Each 5 seconds
			checkLink(link, c)
		}(l) // Function literal, equivalent to Python Lambdas
	}

	// Iterate over length of links
	/*for i := 0; i < len(links); i++ {
		// Print channel messages
		fmt.Println(<-c)
	}*/
}

// expected: link and channel, return string
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		// Send message to channel
		c <- link
		return
	}

	fmt.Println(link, "is up")
	// Send message to channel
	c <- link
}
