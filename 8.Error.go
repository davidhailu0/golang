package src

// import (
// 	"errors"
// 	"fmt"
// )

// type user struct {
// 	name string
// }

// func (u user) Error() string {
// 	return fmt.Sprintf("%v has a problem with his account", u.name)
// }

// func sendSMS() error {
// 	return user{name: "Dave"}
// }
// func main() {
// 	er := sendSMS()
// 	fmt.Println(er.Error())
// 	_, ok := divide(14, 0)
// 	if ok != nil {
// 		fmt.Printf(ok.Error())
// 	}

// }

// func divide(a int32, b int32) (int32, error) {
// 	if b == 0 {
// 		return 0, errors.New("Can not divide by Zero")
// 	}
// 	return a / b, nil
// }
