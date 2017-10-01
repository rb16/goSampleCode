package main

import (
	"net/http"
	"io/ioutil"
	"os"
	"fmt"
	"io"
	"strings"
)

func main() {
	d, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v",  err)
		//os.Exit(1)
	}
	for _, url := range strings.Split(string(d), "\n") {
		resp, err := http.Get("https://"+url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			//os.Exit(1)
		}
		_, err = io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			//os.Exit(1)
		}
		fmt.Println(url)
	}
}
