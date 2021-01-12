package elements

import (
	"github.com/go-vgo/robotgo"
	"math/rand"
)

const buttonClickTimesIn10Clicks = 7
const ideTabChangingTimesIn10Clicks = 7

func KeyboardAction() {
	if rand.Intn(10) <= buttonClickTimesIn10Clicks {
		keyboardTap()
		return
	}
	changeTab()
}

func changeTab() {
	if rand.Intn(10) <= ideTabChangingTimesIn10Clicks {
		robotgo.KeyTap("left", "alt")
		return
	}
	robotgo.KeyTap("tab", "ctrl")
}

func keyboardTap() {
	robotgo.KeyTap(randButton())
}

func randButton() string {
	in := []string{"up", "down", "right", "down", "left", "lctrl", "down", "rctrl", "lshift", "down", "rshift"}
	randomIndex := rand.Intn(len(in))
	return in[randomIndex]
}