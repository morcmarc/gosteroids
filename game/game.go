package game

import (
	"fmt"

	"github.com/morcmarc/gosteroids/game/graphics"
	. "github.com/morcmarc/gosteroids/game/shared"
)

const (
	Width  int    = 600
	Height int    = 600
	Title  string = "Gosteroids"
)

func Start() {
	ctrlChnl := make(chan uint8)
	InitControls(ctrlChnl)
	graphics.Init(Width, Height, Title, ctrlChnl)
}

func InitControls(keyEvents <-chan uint8) {
	go func() {
		for e := range keyEvents {
			switch e {
			case Throttle:
				fmt.Print("↑")
				break
			case Break:
				fmt.Print("↓")
				break
			case Left:
				fmt.Print("←")
				break
			case Right:
				fmt.Print("→")
				break
			}
		}
	}()
}
