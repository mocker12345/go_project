package main

import (
	"fmt"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}
type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}
func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}

	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

type Stringer interface {
	String() string
}
type S struct {
	i int
}

func (s *S) String() string {
	return fmt.Sprintf("%d", s.i)
}
func Print(s Stringer) {
	println(s.String())
}
func DynamicPrint(any interface{}) {
	if s, ok := any.(Stringer); ok {
		Print(s)
	}
}

func main() {

	var s S
	s.i = 123456789
	Print(&s)
	DynamicPrint(&s)
	boxs := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}
	fmt.Printf("We have %d boxes in our set\n", len(boxs))
	fmt.Println("The volume of the first one is", boxs[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxs[len(boxs)-1].color.String())
	fmt.Println("The biggest one is", boxs.BiggestColor().String())
	fmt.Println("Let's paint them all black")
	boxs.PaintItBlack()
	fmt.Println("The color of the second one is", boxs[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxs.BiggestColor().String())
}
