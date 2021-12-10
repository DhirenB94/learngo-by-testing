package perimeter

import "math"


func Perimeter(rectanglePerim Rectangle) float64 {
	return 2 * (rectanglePerim.width + rectanglePerim.height) //access the structs fields
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base float64
	height float64
}

type Shape interface {
	 Area() float64
}

//Rectangle and Circle have methods called Area() that return a float64, so it satisfies the Shape interface


//Structs
//named collection of fields where you can store data


//Methods

//When we call t.Errorf() we are calling the method Errorf on the instance of our t (testing.T).
//A method is a function with a receiver.
//Where you can just call functions wherever you like, such as Area(rectangle) you can only call methods on "things".

//Syntax --> func (receiverName ReceiverType) MethodName(args) -
//convention in Go to have the receiver variable be the first letter of the type.*/



