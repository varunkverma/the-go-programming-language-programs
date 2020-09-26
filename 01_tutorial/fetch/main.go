// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	readResponseUsingIoCopy()
}

func readResponseUsingRealAll() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

func readResponseUsingIoCopy() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		println("RESPONSE CODE: ", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			errURL := url
			if strings.HasPrefix(url, "http://") {
				errURL = "http://" + errURL
			}
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", errURL, err)
			os.Exit(1)
		}
	}
}
