// //package src

import "fmt"

// import (
// 	"fmt"
// )

type Message struct {
	Recipient string
	Text      string
}

func getMessageText(m Message) string {
	return fmt.Sprintf(`
To: %v
Message: %v
`, m.Recipient, m.Text)
}

func main() {
	//to declare a pointer
	// var a *int
	// myString := "Hello"
	// stringPtr := &myString
	//to get the value of the pointer's memory location we use
	//fmt.Println(*stringPtr)
	// strings.ReplaceAll()

	fmt.Println(getMessageText(Message{Recipient: "Hello", Text: "DBhdhvbvhf"}))
}

// type car struct {
// 	color string
// }

// // to change the value of car color
// func (c *car) setColor(color string) {
// 	c.color = color
// }