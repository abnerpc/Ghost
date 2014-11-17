package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

const (
	file = "/etc/hosts"
)

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
			ShortName: "sa",
			Usage:     "show my hosts",
			Action: func(c *cli.Context) {

				if c.Args().First() == "active" {
					showActive()

				} else if c.Args().First() == "inactive" {
					showInative()

				} else {
					showAll()

				}
			},
		},
	}

	app.Run(os.Args)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func showAll() {
	lines, err := readLines(file)

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("===============================================================")
	for i, line := range lines {
		fmt.Println(i, line)

		if string(line[0]) == "#" {
			fmt.Println(i, line, "| INACTIVE")
		}
	}
	fmt.Println("===============================================================")
}

func showInative() {
	lines, err := readLines(file)

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("===============================================================")
	for i, line := range lines {
		if string(line[0]) == "#" {
			fmt.Println("|", i, line, "|")
		}
	}
	fmt.Println("===============================================================")

}

func showActive() {
	lines, err := readLines(file)

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("===============================================================")
	for i, line := range lines {
		if string(line[0]) != "#" {
			fmt.Println("|", i, line, "|")
		}
	}
	fmt.Println("===============================================================")

}
