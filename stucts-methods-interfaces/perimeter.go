package perimeter

func Perimeter(width float64, height float64) float64 {
	return width *2 + height*2
}

func Area(width float64, height float64) float64 {
	return width * height
}

//Structs
//named collection of fields where you can store data

type Rectangle struct {
	width float64
	height float64
}


