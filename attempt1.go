package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CommandMap map[string]interface{}

type Pen struct {
	size int
	down bool
}

func main() {
	p := Pen{}
	dlFile := readDLFile("dl3.txt")
	cmdLines := removeComments(dlFile)
	m := initDLMapping()
	runCommands(cmdLines, m, &p)
}

func readDLFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func removeComments(fContent string) []string {
	var cmdOnly []string
	splitln := strings.Split(fContent, "\n")
	for _, line := range splitln {
		noCommentln := strings.Split(line, "#")[0]
		fmt.Println([]byte(noCommentln))
		cmdOnly = append(cmdOnly, noCommentln)
	}
	return cmdOnly
}

func runCommands(lines []string, m CommandMap, p *Pen) {
	for _, l := range lines {
		cmd := l[0:1]
		fmt.Println(l)
		switch cmd {
		case "P":
			i, _ := strconv.Atoi(string(l[2]))
			m[cmd].(func(*Pen, int))(p, i)
		case "D", "U":
			m[cmd].(func(*Pen))(p)
		case "N", "S", "E", "W":
			i, _ := strconv.Atoi(string(l[2]))
			m[cmd].(func(string, int))(cmd, i)
		default:
			fmt.Println("Command does not exist")
			os.Exit(1)
		}

	}
}

func (p *Pen) setPenSize(size int) {
	(*p).size = size
	fmt.Println("Pen size set to", size)
}

func (p *Pen) penDown() {
	(*p).down = true
	fmt.Println("Pen is down")
}

func (p *Pen) penUp() {
	(*p).down = false
	fmt.Println("Pen is up")
}

func penMove(dir string, dist int) {
	fmt.Println("Pen moved", dist, "to the", dir)
}

//Map each command to a function
func initDLMapping() CommandMap {
	m := CommandMap{
		"P": (*Pen).setPenSize,
		"U": (*Pen).penUp,
		"D": (*Pen).penDown,
		"N": penMove,
		"S": penMove,
		"E": penMove,
		"W": penMove,
	}

	return m
}
