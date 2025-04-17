package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c) // Will hang up the program

	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		//time.Sleep(5 * time.Second)
		//go checkLink(l, c)
		//go func() {
		//	time.Sleep(5 * time.Second)
		//	checkLink(l, c)
		//}()
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		//c <- "Might be down, I think"
		c <- link
	}

	fmt.Println(link, "is up!")
	//c <- "Yep, it's up"
	c <- link
}
