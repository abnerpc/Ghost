package main

import "os"
import "github.com/codegangsta/cli"

func main() {
	app := cli.NewApp()

	app.Name = "Ghost"
	app.Usage = "Manipulate your hosts in the simplest way"
	app.Author = "Vinicius Souza - http://github.com/vsouza"
	app.Email = "hi@vsouza.com"
	app.Action = func(c *cli.Context) {
		println("Ghost")
	}

	app.Commands = []cli.Command{
		{
			Name:      "add",
			ShortName: "a",
			Usage:     "add a host",
			Action: func(c *cli.Context) {
				println("host added!  name: ", c.Args().First(), "   ip: ", c.Args().Get(1))

				if c.Args().First() == "-h" {
					println("FIND")
				}
			},
		},
		{
			Name:      "rm",
			ShortName: "r",
			Usage:     "remove host",
			Action: func(c *cli.Context) {
				println("host removed!: ", c.Args().First())
			},
		},
		{
			Name:      "show",
			ShortName: "r",
			Usage:     "show my hosts",
			Action: func(c *cli.Context) {
				openFile()
				isExists()
				println("see hosts file!: ", c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}

func openFile() {
	file, err := os.Open("/etc/hosts")
	if err != nil {
		println("unable to find your hosts file")
	}

	println(file)
}

func isExists() {
	file, err := os.IsExist("/etc/hosts")
	if err != nil {
		println("not exits")
	}

	println("EXISTS")
}
