package game

import (
	"github.com/morcmarc/gosteroids/game/graphics"
	"github.com/morcmarc/gosteroids/game/objects"
)

const (
	Width  int    = 600
	Height int    = 600
	Title  string = "Gosteroids"
)

func Start() {
	controlChanel := make(chan uint8)
	objectManager := objects.NewObjectManager()

	go objectManager.Listen(controlChanel)

	graphics.Init(Width, Height, Title, controlChanel, objectManager)
}
