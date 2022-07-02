package main

//write a program which counts down from 3, printing each number on a new line (with a 1-second pause)
//and when it reaches zero it will print "Go!" and exit

//We want our Countdown function to write data somewhere and io.Writer is the de-facto way of capturing that as an interface in Go.

//In main we will send to os.Stdout so our users see the countdown printed to the terminal.
//In test we will send to bytes.Buffer so our tests can capture what data is being generated

func main() {
	Countdown()
}