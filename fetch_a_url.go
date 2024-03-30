// Fetch the content found at a URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "https://www.bing.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("response status code: ", resp.StatusCode)
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	// fmt.Println(b)
	fmt.Printf("%s", b)
}
