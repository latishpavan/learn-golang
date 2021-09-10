package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type Square struct {
	Length float64
}

func (s *Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (int, error) {
	diff := byte('a' - 'A')

	out := make([]byte, len(p))

	for index, char := range p {
		if char > 'a' && char < 'z' {
			char -= diff
		}

		out[index] = char
	}

	return c.wtr.Write(out)
}

func main() {
	circle := &Circle{3.0}
	square := &Square{3.0}

	fmt.Println(sumAreas([]Shape{circle, square}))

	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
