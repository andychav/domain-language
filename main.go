package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type Command struct {
	cmd    string
	hasArg bool
	fn     func(string, int)
}

type CommandList []Command

func main() {
	//set up command table
	cmds := initCommands()

	//Read file line by line
	f, err := os.Open("dl3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//For each line look up command
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var cmd string
		var arg int = -1
		ln := sc.Text()
		fmt.Sscanf(ln, "%s %d", &cmd, &arg)
		foundCmd, err := cmds.findCommand(cmd)

		if err != nil {
			log.Fatal(err)
		}

		//Check if command takes an argument
		if foundCmd.hasArg && arg == -1 {
			log.Fatal(cmd, "needs an argument")
		}

		foundCmd.fn(cmd, arg)
	}
}

func initCommands() CommandList {
	cmds := []Command{
		{"P", true, doSelectPen},
		{"U", false, doPenUp},
		{"D", false, doPenDown},
		{"N", true, doPenDir},
		{"S", true, doPenDir},
		{"E", true, doPenDir},
		{"W", true, doPenDir},
	}

	return cmds
}

func (cl CommandList) findCommand(cmd string) (Command, error) {
	for _, c := range cl {
		if c.cmd == cmd {
			return c, nil
		}
	}
	return Command{}, errors.New("Unknown command:" + cmd)
}

func doSelectPen(_ string, size int) {
	fmt.Println("Pen size set to", size)
}

func doPenUp(_ string, _ int) {
	fmt.Println("Pen is up")
}

func doPenDown(_ string, _ int) {
	fmt.Println("Pen is down")
}

func doPenDir(dir string, steps int) {
	fmt.Println("Pen moved", steps, "to the ", dir)
}
