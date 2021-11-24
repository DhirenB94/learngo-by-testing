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
	return greetingPrefix(language) + name

}


func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
//here we get back a named return value (prefix string) -- a string assigned to variable prefix -- so we just have to return, dont have to return prefix


func main () {
	println(Hello("BOB", "French"))
}