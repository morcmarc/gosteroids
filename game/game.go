package game

import (
	"github.com/morcmarc/gosteroids/game/audio"
	"github.com/morcmarc/gosteroids/game/broadcast"
	"github.com/morcmarc/gosteroids/game/graphics"
	"github.com/morcmarc/gosteroids/game/objects"
)

const (
	title string = "Gosteroids"
)

func Start(w, h, bgQuality int, noMusic bool) {
	controlChannel := broadcast.NewBroadcaster()
	defer controlChannel.Write(nil)
	controlChannelListener := controlChannel.Listen()

	objectManager := objects.NewObjectManager()
	audioPlayer := audio.NewAudioPlayer()

	go objectManager.Listen(controlChannelListener)
	go audioPlayer.Listen(controlChannelListener)

	if !noMusic {
		audioPlayer.Play("assets/audio/mass.ogg", -1)
	}

	graphics.Init(w, h, title, bgQuality, controlChannel, objectManager)
}
