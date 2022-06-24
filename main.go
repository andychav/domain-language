package main

import "fmt"

type Pen struct {
	size int
	down bool
}

func main() {
	p := Pen{}
	p.selectPen(2)
	p.putPenDown(true)
	fmt.Println(p)
}

func (p *Pen) selectPen(size int) {
	(*p).size = size
}

func (p *Pen) putPenDown(down bool) {
	(*p).down = down
}

func penMove(dir String, dist int) {
	//todo: logic for moving pen
}
