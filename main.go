package main

import (
	"runtime"

	"github.com/morcmarc/gosteroids/game"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	game.Start()
}
