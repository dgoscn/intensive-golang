package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

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

func fowrward(source <-chan string, destiny chan string) {
	for {
		destiny <- <-source
	}
}

// multiplex - mixing (mensages) in a channel
func join(entry1, entry2 <-chan string) <-chan string {
	c := make(chan string)
	go fowrward(entry1, c)
	go fowrward(entry1, c)
	return c
}

func main() {
	c := join(
		html.title("https://www.twitter.com", "https://www.google.com"),
		html.title("https://www.amazon.com", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
