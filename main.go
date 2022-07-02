package main

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type Command struct {
	cmd    string
	hasArg bool
	fn     func(string, int)
}

type Pen struct {
	size   int
	x      int
	y      int
	doDraw bool
}

type CommandList []Command

var img = image.NewRGBA(image.Rect(0, 0, 100, 100))
var col = color.RGBA{255, 0, 0, 255}
var p = Pen{}

func main() {

	cX, cY := getRectCenter(img.Rect)
	p = Pen{x: cX, y: cY}

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

	pf, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(pf, img)
}

func initCommands() CommandList {
	cmds := []Command{
		{"P", true, p.doSelectPen},
		{"U", false, p.doPenUp},
		{"D", false, p.doPenDown},
		{"N", true, p.doPenDir},
		{"S", true, p.doPenDir},
		{"E", true, p.doPenDir},
		{"W", true, p.doPenDir},
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

func (p *Pen) doSelectPen(_ string, size int) {
	fmt.Println("Pen size set to", size)
}

func (p *Pen) doPenUp(_ string, _ int) {
	fmt.Println("Pen is up")
	p.doDraw = false
}

func (p *Pen) doPenDown(_ string, _ int) {
	fmt.Println("Pen is down")
	p.doDraw = true
}

//this needs a pen receiver
func (p *Pen) doPenDir(dir string, steps int) {
	fmt.Println("Pen moved", steps, "to the ", dir)

	px := p.x
	py := p.y
	fmt.Println("X:", p.x, "Y:", p.y)
	i := 1
	for i <= steps {
		switch dir {
		case "N":
			py--
		case "S":
			py++
		case "E":
			px++
		case "W":
			px--
		}
		if p.doDraw {
			img.Set(px, py, col)
		}
		i++
	}
	(*p).x = px
	(*p).y = py

}

func getRectCenter(r image.Rectangle) (x, y int) {
	return r.Dx() / 2, r.Dy() / 2
}
