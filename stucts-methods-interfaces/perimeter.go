package perimeter

func Perimeter(rectanglePerim Rectangle) float64 {
	return 2 * (rectanglePerim.width + rectanglePerim.height) //access the structs fields
}

func Area(rectangleArea Rectangle) float64 {
	return rectangleArea.width * rectangleArea.height
}





//Structs
//named collection of fields where you can store data

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}


