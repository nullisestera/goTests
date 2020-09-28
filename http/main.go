package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// GET ENDPOINT HTTP
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// Calling io.Copy
	// function to copy res.Body Reader interface to os.Stdout (output) using res.Body Writer interface
	// Copy (dst Writer, src Reader)
	// dst: destiny
	// src: source
	// io.Copy(os.Stdout, res.Body)

	lw := logWriter{}
	io.Copy(lw, res.Body)
}

// Custom Writer
func (logWriter) Write(bs []byte) (int, error) {
	// print the byte slice returned by the Reader interface
	fmt.Println(string(bs))

	// return length of byteslice
	fmt.Println("Length of Byte Slice", len(bs))
	return len(bs), nil
}
