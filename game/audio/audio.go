package audio

import (
	"fmt"
	"os"

	b "github.com/morcmarc/gosteroids/game/broadcast"
	. "github.com/morcmarc/gosteroids/game/shared"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_mixer"
)

type AudioPlayer struct {
	Volume int
	Muted  bool
}

func NewAudioPlayer() *AudioPlayer {
	ap := &AudioPlayer{
		Volume: 75,
		Muted:  false,
	}
	return ap
}

func (a *AudioPlayer) Listen(cc b.Receiver) {
	for m := cc.Read(); m != nil; m = cc.Read() {
		if m == VolumeDown && a.Volume > -1 {
			a.Volume -= 1
			mix.SetMusicVolume(a.Volume)
		}
		if m == VolumeUp && a.Volume < 101 {
			a.Volume += 1
			mix.SetMusicVolume(a.Volume)
		}
		if m == Mute {
			a.Muted = !a.Muted
			if a.Muted {
				mix.SetMusicVolume(0)
			} else {
				mix.SetMusicVolume(a.Volume)
			}
		}
	}
}

func (a *AudioPlayer) Play(filename string, loop int) {
	_, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}

	mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096)
	if err := sdl.GetError(); err != nil {
		panic(err)
	}

	mix.SetMusicVolume(a.Volume)

	m := mix.LoadMUS(filename)
	if err := sdl.GetError(); err != nil {
		fmt.Printf("%s (Have you installed SDL2 with Ogg support?)\n", err)
	}

	m.Play(loop)
}
