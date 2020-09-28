package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// GET ENDPOINT HTTP
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// Calling io.Copy 
	// function to copy res.Body Reader interface to os.Stdout (output) using res.Body Writer interface
	io.Copy(os.Stdout, res.Body)
}
