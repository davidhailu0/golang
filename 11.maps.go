// //package src

// import (
// 	"errors"
// 	"fmt"
// )

// //hash map or dictionaries
// func main(){
//    keyValueTable := make(map[string]int)
//    keyValueTable["a"] = 21
//    fmt.Println(keyValueTable["a"])
//    //---------------------------------//
//    age := map[int]string{
// 	//1:"newBorn"
//    }

//    age[1] = "newBorn"
//    //to delete
//    delete(age,1)
//    //to check if it exists
//    elem, ok := age[1]
//    fmt.Println(age[1])
//    fmt.Println(len(age))
//    fmt.Println(getUserMap([]string{"Abebe","ddjdjd","ddjdjdjd"},[]int{1,2,3}))
//    test := map[string]map[string]int{}
// }

// func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
// 	if len(names)!=len(phoneNumbers){
// 		return map[string]user{}, errors.New("cfcfffcfcf")
// 	}
// 	shshs := map[string]user{}
// 	for i:=0;i<len(names);i++{
// 		shshs[names[i]] = user{
// 			name:names[i],
// 			phoneNumber:phoneNumbers[i],
// 		}
// 	}
// 	return shshs,nil
// }

type user struct {
	name        string
	phoneNumber int
}