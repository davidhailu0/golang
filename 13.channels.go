package main

import (
	"fmt"
	"time"
)

func sendEmail() {
	// the go keyword makes the anonymous function work asynchroneously
	go func() {
		time.Sleep(1000)
		fmt.Println("Email Sent")
	}()
}

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

func main() {
	//to create a channel
	// channel := make(chan int)
	//to send data to channel
	// channel <- 12
	//to receive data from channel
	// val := <-channel
	// we can also use buffered channel
	channel := addEmailsToQueue([]string{"abcd", "abcd", "abcd"})
	for i := 0; i < 3; i++ {
		fmt.Println(<-channel)
	}
	//it breaks out of the loop if the channel is closed and blocks the process until an item is received from other routine
	for item := range channel {
		fmt.Println(item)
	}

	// select is used to listen to multiple channels
	//after one of the select statement is executed flow transfets out of the select
	select {
	case i, ok := <-chInts:
		fmt.Println(i)
	case str, ok := <-chString:
		fmt.Println(str)
	}
	// to check if the channel is closed or empty
	// _, ok := <-channel
	// if !ok {
	// 	fmt.Println("The Channel is closed or empty")
	// }
	//to close channel
	close(channel)
	// fmt.Println(([3]email{
	// 	{body: "abcd", date: time.Now()},
	// 	{body: "efgh", date: time.Now()},
	// 	{body: "ijkl", date: time.Now()},
	// }))
}

func addEmailsToQueue(emails []string) chan string {
	channels := make(chan string, len(emails))
	for _, email := range emails {
		channels <- email
	}
	return channels
}
