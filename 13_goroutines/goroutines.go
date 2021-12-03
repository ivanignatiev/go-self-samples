package main

import (
	"fmt"
	"time"
)

func abc(syncChannelABC chan string) {
	syncChannelABC <- "a"
	syncChannelABC <- "b"
	time.Sleep(4 * time.Second)
	syncChannelABC <- "c"
}

func def(syncChannelDEF chan string) {
	syncChannelDEF <- "d"
	syncChannelDEF <- "e"
	time.Sleep(2 * time.Second)
	syncChannelDEF <- "f"
}

func greeting(greetingChannel chan string) {
	greetingChannel <- "hi!"
}

func main() {

	greetingChannel := make(chan string)
	go greeting(greetingChannel)
	fmt.Println(<-greetingChannel)

	syncChannelABC := make(chan string)
	syncChannelDEF := make(chan string)
	go abc(syncChannelABC)
	go def(syncChannelDEF)

	fmt.Println(<-syncChannelABC)
	fmt.Println(<-syncChannelDEF)
	fmt.Println(<-syncChannelABC)
	fmt.Println(<-syncChannelDEF)
	fmt.Println(<-syncChannelABC)
	fmt.Println(<-syncChannelDEF)

	time.Sleep(10 * time.Second)
	fmt.Println("end main()")
}
