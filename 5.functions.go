// package main

// import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

// // this is also valid
// func substract(x, y int) int {
// 	return x - y
// }

// // works by pass by value
// func getMonthlyPrice(tier string) int {
// 	if tier == "basic" {
// 		return 100
// 	} else if tier == "premium" {
// 		return 150.
// 	} else if tier == "enterprise" {
// 		return 500
// 	} else {
// 		return 0
// 	}
// }

// func getPoint() (x, y int) {
// 	return 3, 4
// }

// func main() {
// 	fmt.Println("basic")

// 	//x, _ = getPoint() //we can use the underscore to remove the unwanted value
// 	a, x := getCoords()
// 	fmt.Println(a, x)
// }

// // Should only be used in short function
// func getCoords() (x, y int) {
// 	// x and y are initialized with zero values
// 	x = 14
// 	return // automatically returns x and y
// }
