package game

import (
	"fmt"
)

const (
	THROTTLE uint8 = 0
	BREAK    uint8 = 1
	LEFT     uint8 = 2
	RIGHT    uint8 = 3
)

func InitControls(keyEvents <-chan uint8) {
	go func() {
		for e := range keyEvents {
			switch e {
			case THROTTLE:
				fmt.Print("↑")
				break
			case BREAK:
				fmt.Print("↓")
				break
			case LEFT:
				fmt.Print("←")
				break
			case RIGHT:
				fmt.Print("→")
				break
			}
		}
	}()
}
