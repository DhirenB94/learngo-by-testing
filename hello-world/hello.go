package main


const englishHelloPrefix = "Hello, "


func Hello(name string) string {
	return englishHelloPrefix + name
}
func main () {
	println(Hello("world"))
}
