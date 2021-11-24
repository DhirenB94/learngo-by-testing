package main


const englishHelloPrefix = "Hello, "


func Hello(name string) string {
	if name == "" {
		return "Hello, world"
	}
	return englishHelloPrefix + name
}
func main () {
	println(Hello(""))
}