package main

import (
	"flag"
	"runtime"

	"github.com/morcmarc/gosteroids/game"
)

var (
	width             int
	height            int
	animateBackground bool
)

func init() {
	runtime.LockOSThread()
	flag.IntVar(&width, "w", 600, "Width")
	flag.IntVar(&height, "h", 600, "Height")
	flag.BoolVar(&animateBackground, "bg", true, "Animate background (requires a decent graphics card)")
}

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	flag.Parse()

	game.Start(width, height, animateBackground)
}
