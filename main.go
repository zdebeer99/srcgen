package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/zdebeer99/srcgen/generator"
	"os"
)

/*
Generate code from data and templates.
  init - create config file and folders.
*/

func main() {
	app := cli.NewApp()
	app.Name = "srcgen"
	app.Version = generator.Version
	app.Usage = "Source Code Generator"
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}
	app.Commands = commands
	app.Run(os.Args)
}

var commands = []cli.Command{
	{
		Name:        "init",
		Usage:       "Initialize a code gen project.",
		Description: "init creates a 'srcgen.yaml' config file and creates the templatedata and template folders.",
		Action: func(c *cli.Context) {
			err := generator.InitializeGenProject()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("srcgen initialized.")
			}
		},
	},
	{
		Name:        "gen",
		Usage:       "Generate code.",
		Description: "Generate code based on the settings provided in the 'srcgen.yaml' file.",
		Action: func(c *cli.Context) {
			err := generator.Generate()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("done.")
			}
		},
	},
}
