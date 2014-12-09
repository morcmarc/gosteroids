package main

import (
	"flag"
	"runtime"

	"github.com/morcmarc/gosteroids/game"
)

var (
	width             int
	height            int
	animateBackground int
	noMusic           bool
)

func init() {
	runtime.LockOSThread()
	flag.IntVar(&width, "w", 512, "Width")
	flag.IntVar(&height, "h", 512, "Height")
	flag.IntVar(&animateBackground, "bg", 1, "Background quality (0: low, 1: med, 2: high)")
	flag.BoolVar(&noMusic, "m", false, "No music")
}

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	flag.Parse()

	game.Start(width, height, animateBackground, noMusic)
}
