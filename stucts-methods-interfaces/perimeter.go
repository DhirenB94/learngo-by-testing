package perimeter

func Perimeter(rectanglePerim Rectangle) float64 {
	return 2 * (rectanglePerim.width + rectanglePerim.height) //access the structs fields
}

func Area(rectangleArea Rectangle) float64 {
	return rectangleArea.width * rectangleArea.height
}
// current error : cannot use circle (type Circle) as type Rectangle in argument to Area


//Structs
//named collection of fields where you can store data

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

//Methods

/*
So far we have only been writing functions but we have been using some methods. When we call t.Errorf we are calling the method Errorf on the instance of our t (testing.T).
A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.
Methods are very similar to functions but they are called by invoking them on an instance of a particular type. Where you can just call functions wherever you like, such as Area(rectangle) you can only call methods on "things".
*/

