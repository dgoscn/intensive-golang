package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - channel onlye reading
func title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	t1 := title("https://www.twitter.com", "https://www.google.com")
	t2 := title("https://www.amazon.com", "https://www.youtube.com")
	fmt.Println("The first:", <-t1, "|", <-t2)
	fmt.Println("The Second:", <-t1, "|", <-t2)
}
