//package src

// A type can implement many interfaces

// type Shape interface {
// 	area() float32
// 	perimeter() float32
// 	//withParams(d float32, e float32) (returnedVal float32) //with named value
// 	// withParams(float32,float32)(float32)
// }

// type Geometry interface {
// 	sine() float32
// 	cosine() float32
// 	tan() float32
// }

// type Rectangle struct {
// 	height float32
// 	width  float32
// }

// func (r Rectangle) area() float32 {
// 	return r.height * r.width
// }

// func (r Rectangle) perimeter() float32 {
// 	return 2 * (r.height + r.width)
// }

// func (r Rectangle) cosine() float32 {
// 	return r.height * r.width * 0.2
// }

// type Circle struct {
// 	radius float32
// }

// func (c Circle) area() float32 {
// 	return c.radius * c.radius * 3.1425
// }

// func (c Circle) perimeter() float32 {
// 	return 2 * 3.1425 * c.radius
// }
// func main() {
// 	c := Circle{
// 		radius: 5.214,
// 	}
// 	r := Rectangle{
// 		height: 4.5,
// 		width:  4.2,
// 	}
// 	printAreaAndPerimeter(c)
// 	printAreaAndPerimeter(r)
// }

// func printAreaAndPerimeter(s Shape) {
// 	_, ok := s.(Circle)
// 	if !ok {
// 		fmt.Println("The Input is not a circle")
// 	}
// 	_, ok = s.(Rectangle)
// 	if !ok {
// 		fmt.Println("The Input is not a rectangle")
// 	}
// 	fmt.Printf("The Area is %.2f\n", s.area())
// 	fmt.Printf("The Perimeter is %.2f\n", s.perimeter())

// 	switch typeOfVar := s.(type){
// 	case "Circle":
// 		fmt.Printf("Type of the Variable is %s",&typeOfVar)

// 	}
// }
