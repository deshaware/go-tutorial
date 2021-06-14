package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/10")
	if err != nil {
		fmt.Print("Error:", err)
		os.Exit(1)
	}
	// fmt.Println(res)
	lw := logWriter{}

	io.Copy(lw, res.Body)

}

// Remember, just defining the function and associating with logWriter
// The logWriter is implementing Writer interface

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote these many bytes:", len(bs))
	return len(bs), nil
}
