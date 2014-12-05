package game

import (
	"github.com/morcmarc/gosteroids/game/audio"
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
	audioPlayer := audio.NewAudioPlayer()

	go objectManager.Listen(controlChanel)
	// TODO: fps killer go channel, we have to live without volume controls
	// for now
	// go audioPlayer.Listen(controlChanel)

	audioPlayer.Play("assets/audio/mass.ogg", -1)

	graphics.Init(Width, Height, Title, controlChanel, objectManager)
}
