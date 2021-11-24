package main


const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"


func Hello(name string, language string) string {
	if name == "" {
		return "Hello, world"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}
func main () {
	println(Hello("BOB", "Spanish"))
}