package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "timely"
	app.Usage = "get an timely info"

	app.Commands = []cli.Command{
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
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
	}

	app.Run(os.Args)
}
