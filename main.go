package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/morcmarc/gosteroids/game"
)

func main() {
	app := cli.NewApp()
	app.Name = "Gosteroids"
	app.Version = "0.1.0"
	app.Usage = "Asteroids clone built in Go"
	app.Action = runGame
	app.Run(os.Args)
}

func runGame(c *cli.Context) {
	game.Start()
}
