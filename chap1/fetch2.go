package main

import (
	"fmt"
	"time"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //send to channel
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start)
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func main(){
	start := time.Now()
	ch := make(chan string)
	d, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	for _, url := range strings.Split(string(d), "\n") {
		go fetch("https://www."+url, ch)
	}
	for range strings.Split(string(d), "\n") {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elased\n", time.Since(start).Seconds())
}
