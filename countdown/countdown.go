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
	const delay = 1000000000
	for i := countdownStart; i > 0; i-- {
		time.Sleep(delay)
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
