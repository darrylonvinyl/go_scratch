// Fetch prints the content found at a URL. The difference here with the orig.
// version is that it pastes everything to stdout.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// have_http, err := regexp.Match(`(?m)^http(s?)`, []byte(url))
		have_http := strings.HasPrefix(url, "http")
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: address error - %s %v \n", url, err)
			os.Exit(1)
		}
		if !have_http {
			fmt.Printf("Adding required https header to %s \n", url)
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "HTTP status code: %s", resp.Status)
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: reading %s %v \n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
