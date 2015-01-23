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
	ALL = 0
	ACTIVE = 1
	INACTIVE = 2
)

func main() {

	app := cli.NewApp()

	app.Name = "Ghost"
	app.Usage = "Manipulate your hosts in the simplest way"
	app.Author = "Vinicius Souza - http://github.com/vsouza"
	app.Email = "hi@vsouza.com"
	app.Action = cli.ShowAppHelp

	app.Commands = getCommands()

	app.Run(os.Args)
}

func getCommands() []cli.Command {

	return []cli.Command{
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
			Action:	execShow,
		},
	}

}

func execShow(c *cli.Context) {

	lines, success := readLines(file)
	if ! success {
		return
	}
	
	
	fmt.Println("===============================================================")

	if c.Args().First() == "active" {
		show(lines, ACTIVE)
	} else if c.Args().First() == "inactive" {
		show(lines, INACTIVE)
	} else {
		show(lines, ALL)
	}

	fmt.Println("===============================================================")
}

func readLines(path string) ([]string, bool) {
	
	file, err := os.Open(path)
	
	if err != nil {
		log.Fatalf("readLines: %s", err)
		return nil, false
	}
	
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatalf("readLines: %s", scanner.Err())
		return nil, false
	}
	
	return lines, true
}

func show(lines []string, mode int) {
	
	first_letter := ""

	for i, line := range lines {
		
		first_letter = string(line[0])
		
		print := 
			(mode == ALL) ||
			(mode == ACTIVE && first_letter != "#") ||
			(mode == INACTIVE && first_letter == "#")

		if ! print { continue }

		fmt.Println(i, line)
	}

}