// Fetch prints the content found at a URL. The difference here with the orig.
// version is that it pastes everything to stdout.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: reading %s %v \n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
