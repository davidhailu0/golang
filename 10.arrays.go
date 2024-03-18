package src

// func main(){
//  fruits := []string{"Apple","Pineapple","Papaya"}

//  for ind, el := range fruits{
// 	fmt.Println(ind+1," is ",el)
//  }
// }
package main

import "fmt"

func main() {
	// to declare array of 10 integers
	// var intNum [10]int
	// intNum[0] = 1
	// fmt.Println(intNum[0])
	// intNum2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(&intNum2[0]) // prints the location address of element
	// fmt.Println(intNum2[2:])
	// fmt.Println(intNum2[:5])
	// fmt.Println(intNum2[2:9]) // The last item won't get added
	// fmt.Println(intNum2[0])
	// strArr, intArr := getMessageWithRetries("Hello", "World", "123")
	// fmt.Println(strArr)
	// fmt.Println(intArr)

	// //Slices use dynamic allocation of size based on the previous size of array
	// mySlice := make([]int, 2, 5) // the first argument specifies the type, the second specifies the length and the third specifies the capacity

	// mySlice = make([]int, 5) // this default the capacity equal to the length

	// fmt.Println(len(mySlice)) // len prints out the number of elements the array is currently holding
	// fmt.Println(cap(mySlice)) // cap prints out the capacity

	// strArr2 := []string{"Abebe", "Beso", "Bella"}
	// strArr2 = append(strArr2, "Kebede","Killo","Chemere")
	// stringAdd(strArr2...)
	// fmt.Println(createMatrix(10, 10))
	i := make([]int, 3, 8)
	fmt.Println("len of i:", len(i))
	// len of i: 3
	fmt.Println("cap of i:", cap(i))
	// cap of i: 8
	fmt.Println("appending 4 to j from i")
	// appending 4 to j from i
	j := append(i, 4)
	fmt.Println("j:", j)
	// j: [0 0 0 4]
	fmt.Println("addr of j:", &j[0])
	// addr of j: 0x454000
	fmt.Println("appending 5 to g from i")
	// appending 5 to g from i
	g := append(i, 5)
	fmt.Println("addr of g:", &g[0])
	// addr of g: 0x454000
	fmt.Println("addr of i:", &i[0])
	fmt.Println("i:", i)
	// i: [0 0 0]
	fmt.Println("j:", j)
	// j: [0 0 0 5]
	fmt.Println("g:", g)

}

// I can pass many arguments to function using ...

func stringAdd(str ...string) string {
	final := ""
	for i := 0; i < len(str); i++ {
		final += str[i]
	}
	return final
}

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	stringArr := [3]string{primary, secondary, tertiary}
	costArr := [3]int{len(primary), len(primary) + len(secondary), len(primary) + len(secondary) + len(tertiary)}
	return stringArr, costArr
}
func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}
	for i := 0; i < cols; i++ {
		arr := []int{}
		for j := 0; j < rows; j++ {
			arr = append(arr, i*j)
		}
		matrix = append(matrix, arr)
	}
	return matrix
}
