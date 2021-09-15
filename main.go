package main

import (
	"log"
	"os"

	"github.com/linzhengen/bookmark-generator-go/cmd"
	"github.com/urfave/cli/v2"
)

var VERSION = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "bookmark-generator-go"
	app.Version = VERSION
	app.Usage = "This tool generates chrome bookmarks from yaml file"
	app.Commands = []*cli.Command{
		cmd.GenerateCmd(),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
