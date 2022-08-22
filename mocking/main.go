package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

func Countdown(out io.Writer, sleeper Sleeper) {
	// 誤ったコード (デバッグ用)
	// for i := countdownStart; i > 0; i-- {
	// 	sleeper.Sleep()
	// }

	// for i := countdownStart; i > 0; i-- {
	// 	fmt.Fprintln(out, i)
	// }

	// 正しいコード
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
