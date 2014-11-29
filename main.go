package main

import (
	"runtime"

	"github.com/morcmarc/gosteroids/game"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	game.Start()
}
