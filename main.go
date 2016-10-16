package main

import (
	"os"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "timely"
	app.Usage = "get an timely info"

	app.Commands = []cli.Command{
		{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "Show all timely info",
			Action:  func(c *cli.Context) error {
				Qiita("")
				Hatena("")
				Github("")
				return nil
			},
		},
		{
			Name:    "qiita",
			Aliases: []string{"q"},
			Usage:   "Get timely info from qiita",
			Action:  func(c *cli.Context) error {
				args := c.Args().First()
				Qiita(args)
				return nil
			},
		},
		{
			Name:    "github",
			Aliases: []string{"g"},
			Usage:   "Get timely info from github",
			Action:  func(c *cli.Context) error {
				args := c.Args().First()
				Github(args)
				return nil
			},
		},
		{
			Name:    "hatena",
			Aliases: []string{"h"},
			Usage:   "Get timely info from hatena bookmark",
			Action:  func(c *cli.Context) error {
				args := c.Args().First()
				Hatena(args)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
