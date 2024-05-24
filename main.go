package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://go.dev",
		"http://amazon.com",
		"wowclassic.blizzard.com", // just add http:// for correct adress
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)

	}

	for l := range c {
		go func(link string) { // you can use just "l"
			time.Sleep(5 * time.Second)
			checkLink(link, c) // you can use just "l"
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		//c <- "Might be down i think"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	//c <- "is up!"
	c <- link
}
