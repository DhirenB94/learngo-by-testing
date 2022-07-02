package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

//write a program which counts down from 3, printing each number on a new line (with a 1-second pause)
//and when it reaches zero it will print "Go!" and exit

//We want our Countdown function to write data somewhere and io.Writer is the de-facto way of capturing that as an interface in Go.

//In main we will send to os.Stdout so our users see the countdown printed to the terminal.
//In test we will send to bytes.Buffer so our tests can capture what data is being generated

const countdownStart = 3
const finalWord = "Go!"

//define our dependency as an interface.
//This lets us then use a real Sleeper in main and a spy sleeper in our tests

type Sleeper interface {
	Sleep()
}

//create a real sleeper that implements the interface we need and pass to func main

type RealSleeper struct{}

func (rs *RealSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &RealSleeper{}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
		//tests will now pass, because now we call the injected in dependency rather than time.sleep()
	}
	fmt.Fprintf(out, finalWord)
}
