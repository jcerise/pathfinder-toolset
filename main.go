package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	version = "0.0.1"
	app = cli.NewApp()
)

func init() {
	app.Name = "Pathfinder Toolset"
	app.Usage = "A collection of usefule Pathfinder Player Character and Gamemaster tools"
	author := &cli.Author{Name:"Jeremy Cerise", Email:"mail@jeremyceri.se"}
	app.Authors = []*cli.Author{author}
	app.Version = version
}

func commands() {
	app.Commands = []*cli.Command {
		{
			Name: "about",
			Aliases: []string{},
			Action: func(c *cli.Context) error {
				fmt.Println("   ___       __  __   ____         __          ______          __        __ ")
				fmt.Println("  / _ \\___ _/ /_/ /  / _(_)__  ___/ /__ ____  /_  __/__  ___  / /__ ___ / /_")
				fmt.Println(" / ___/ _ `/ __/ _ \\/ _/ / _ \\/ _  / -_) __/   / / / _ \\/ _ \\/ (_-</ -_) __/")
				fmt.Println("/_/   \\_,_/\\__/_//_/_//_/_//_/\\_,_/\\__/_/     /_/  \\___/\\___/_/___/\\__/\\__/ ")
				fmt.Println("")
				fmt.Println("Conceived and Coded by Jeremy Cerise, February 2020")
				fmt.Println("Version " + version)
				fmt.Println("All Pathfinder Content used under the OGL")
				return nil
			},
		},
		{
			Name: "new_character",
			Aliases: []string{"nc", "new_char"},
			Action: func(c *cli.Context) error {

				return nil
			},
		},
	}
}

func main() {
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
