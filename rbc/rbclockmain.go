package main

import (
    "fmt"
	"github.com/cchrysostom/RollingBallClock/rbclock"
)


func main() {
    fmt.Println("Rolling ball clock simulator.")
    rbclock.BallOrderRepeat(27)
    rbclock.BallOrderRepeat(45)
    rbclock.BallOrderRepeat(127)
    rbclock.DisplayClockAfterMinutes(30, 325)
}

