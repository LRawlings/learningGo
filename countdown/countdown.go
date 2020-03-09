package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Countdown(out io.Writer) {
	const finalWord = "Go!"
	const countdownStart = 3
	const delay = 1
	for i := countdownStart; i > 0; i-- {
		time.Sleep(delay * time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(delay * time.Second)
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
