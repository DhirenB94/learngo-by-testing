package main


const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"
const frenchHelloPrefix = "Bonjour, "
const french = "French"


func Hello(name string, language string) string {
	if name == "" {
		return "Hello, world"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	} else if language == french {
		return  frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}
func main () {
	println(Hello("BOB", "French"))
}