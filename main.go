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
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// to receive data from channel
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// an easier way is to do a for loop
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// Repeating Routines using infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Alternative way of writing Repeating Routines using infinite loop
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// inserting a pause
	// for l := range c {
	// 	time.Sleep(5 * time.Second)
	// 	go checkLink(l, c)
	// }

	// adding a function literal
	// for l := range c {
	// 	go func() {
	// 		time.Sleep(5 * time.Second)
	// 		checkLink(l, c)
	// 	}()
	// }

	// fixing the function literal
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// func checkLink(link string, c chan string) {
// 	// we don't need the resp (_) to verify whether the website is up
// 	_, err := http.Get(link)
// 	if err != nil {
// 		fmt.Println(link, "might be down!")
// 		c <- "Might be down I think"
// 		return
// 	}
// 	fmt.Println(link, "is up!")
// 	c <- "Yep it's up"
// }

// Repeating Routines
func checkLink(link string, c chan string) {
	// we don't need the resp (_) to verify whether the website is up
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
