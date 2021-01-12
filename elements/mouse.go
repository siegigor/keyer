package elements

import (
	"github.com/go-vgo/robotgo"
	"math/rand"
)

const minX = 650
const maxX = 1200

const minY = 350
const maxY = 600

const mouseScrollTimesIn10Clicks = 6
const downDirectionTimesIn10Clicks = 7

const minScroll = 30
const maxScroll= 100

func MouseAction() {
	if rand.Intn(10) <= mouseScrollTimesIn10Clicks {
		mouseScroll()
		return
	}
	mouseClick()
}

func mouseClick() {
	robotgo.MoveClick(randIntFromRange(minX, maxX), randIntFromRange(minY, maxY), "left")
}

func mouseScroll() {
	direction := "up";
	if rand.Intn(10) <= downDirectionTimesIn10Clicks {
		direction = "down"
	}
	robotgo.ScrollMouse(randIntFromRange(minScroll, maxScroll), direction);
}

func randIntFromRange(min int, max int) int {
	return rand.Intn(max - min) + min
}