package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/rich1111/gobot"
)

func main() {
	app := cli.NewApp()
	app.Name = "gobot"
	app.Author = "The Gobot team"
	app.Email = "https://github.com/rich1111/gobot"
	app.Version = gobot.Version()
	app.Usage = "Command Line Utility for generating new Gobot adaptors, drivers, and platforms"
	app.Commands = []cli.Command{
		Generate(),
	}
	app.Run(os.Args)
}
