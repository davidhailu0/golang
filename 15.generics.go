package main

import "fmt"

// Generics are like java that can be used by any type
func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func printOut[T any](g T) {
	fmt.Println(g)
}

func main(){
	var d [T any]
	fmt.Println(d)
}

// we can make it specific by adding interfaces
type stringer interface {
    String() string
}

func concat[T stringer](vals []T) string {
    result := ""
    for _, val := range vals {
        // this is where the .String() method
        // is used. That's why we need a more specific
        // constraint instead of the any constraint
        result += val.String()
    }
    return result
}

