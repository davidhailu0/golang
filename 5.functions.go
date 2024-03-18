package src

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// // func add(x int, y int) int {
// // 	return x + y
// // }

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
// 	aggregate(3,2,1,func(x ,y int)int{
// 		return x*y
// 	})
// }

// // Should only be used in short function
// func getCoords() (x, y int) {
// 	// x and y are initialized with zero values
// 	x = 14
// 	return // automatically returns x and y
// }

// func aggregate(a,b,c int, arthemetic func(int,int) int)int{
// 	return arthemetic( a, arthemetic(b,c))
// }

// func add(a,b int)int{
// 	return a+b
// }

// // Closure
// // doc will not be removed because it is being used by the return func
// func concatter() func(string) string {
// 	doc := ""
// 	return func(word string) string {
// 		doc += word + " "
// 		return doc
// 	}
// }

// func CopyFile(dstName, srcName string) (written int64, err error) {

// 	// Open the source file
// 	src, err := os.Open(srcName)
// 	if err != nil {
// 	  return
// 	}
// 	// Close the source file when the CopyFile function returns
// 	defer src.Close()

// 	// Create the destination file
// 	dst, err := os.Create(dstName)
// 	if err != nil {
// 	  return
// 	}
// 	// Close the destination file when the CopyFile function returns
// 	defer dst.Close()

// 	return io.Copy(dst, src)
//   }
