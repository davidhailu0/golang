package main

// import "fmt"

// type Car struct {
// 	maker      string
// 	model      string
// 	doors      int
// 	mileage    int
// 	frontwheel Wheel
// 	backWheel  Wheel
// }

// func (c Car) drive() {
// 	fmt.Println("Pip Pip")
// }

// type Wheel struct {
// 	radius   int
// 	material string
// }

// type Truck struct {
// 	Car
// 	length float32
// }

// func main() {
// 	a := Car{} //this create an instance of a Car Type

// 	truck := Truck{
// 		length: 12,
// 		Car: Car{
// 			maker: "BMW",
// 			model: "AL21",
// 		},
// 	}
// 	fmt.Println(truck.length)
// 	fmt.Println(truck.maker)
// 	a.backWheel.radius = 3
// 	truck.drive()
// 	fmt.Println(a.backWheel.radius)

// 	bike := struct {
// 		numberOfWheels int
// 		radius         int
// 	}{
// 		numberOfWheels: 2,
// 		radius:         1,
// 	}
// 	fmt.Println(bike.numberOfWheels)
// }
