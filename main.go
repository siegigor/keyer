package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"keyer/elements"
	"math/rand"
	"time"
)

const maxIntervalSeconds = 8
const keyboardTimesIn10CLicks = 7
var isRun bool = false

func main() {
	fmt.Print("'y' to start, 'esc' to pause, 'esc' 2 times to exit\n")

	robotgo.EventHook(hook.KeyDown, []string{"y"}, func(e hook.Event) {
		fmt.Println("\nstarted")
		if isRun {
			return
		}
		isRun = true
		go initiateClicks()
	})
	robotgo.EventHook(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		fmt.Println("\npause")
		if !isRun {
			robotgo.EventEnd()
		}
		isRun = false
	})
	<- robotgo.EventProcess(robotgo.EventStart())
}

func initiateClicks() {
	for {
		if !isRun {
			return
		}
		randSleep()
		if rand.Intn(10) <= keyboardTimesIn10CLicks {
			elements.KeyboardAction()
		} else {
			elements.MouseAction()
		}
	}
}

func randSleep()  {
	duration := time.Duration(rand.Int31n(maxIntervalSeconds))
	time.Sleep(duration * time.Second)
}