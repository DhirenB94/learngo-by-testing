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

type Shape interface {
	 Area() float64
}





//Structs
//named collection of fields where you can store data


//Methods
/*
So far we have only been writing functions but we have been using some methods.
When we call t.Errorf we are calling the method Errorf on the instance of our t (testing.T).
A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.
Methods are very similar to functions but they are called by invoking them on an instance of a particular type. Where you can just call functions wherever you like, such as Area(rectangle) you can only call methods on "things".

The syntax for declaring methods is almost the same as functions and that's because they're so similar. The only difference is the syntax of the method receiver - func (receiverName ReceiverType) MethodName(args) -
When your method is called on a variable of that type, you get your reference to its data via the receiverName variable. In many other programming languages this is done implicitly and you access the receiver via this.
It is a convention in Go to have the receiver variable be the first letter of the type.*/

