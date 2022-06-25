package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pen struct {
	size int
	down bool
}

func main() {
	file := readDLFile("dl.txt")

	// Split string line by line
	spFile := splitDLFile(file)
	print(spFile)

	// Check lines for something that is not part of map and throw error if there is
	// Map each command to function
}

func readDLFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

//I think I should rename this for clarity
func splitDLFile(content []byte) []string {
	var commandLines []string
	for {
		advance, token, err := bufio.ScanLines(content, true)
		if advance == 0 {
			break
		}
		cl := strings.Split(string(token), "#")
		commandLines = append(commandLines, cl[0])
		fmt.Println(advance, string(token), err)

		if advance <= len(content) {
			content = content[advance:]
		}
	}
	return commandLines
}

func (p *Pen) selectPen(size int) {
	(*p).size = size
}

func (p *Pen) putPenDown(down bool) {
	(*p).down = down
}

func penMove(dir string, dist int) {
	fmt.Println("Pen moved", dist, "to the", dir)
}

func print(slice []string) {
	for i, s := range slice {
		fmt.Println(i, s)
	}
}
