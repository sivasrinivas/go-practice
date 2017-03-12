package main

import (
	"fmt"
	"math"
)

type Calculator struct {
	acc float64
}

func (c *Calculator) Do(op func(float64) float64) float64 {
	c.acc = op(c.acc)
	return c.acc
}

func Add(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc + n
	}
}

func Sub(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc - n
	}
}

func Mul(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc * n
	}
}

func Sqrt() func(float64) float64 {
	return func(acc float64) float64 {
		return math.Sqrt(acc)
	}
}

func main() {
	var c Calculator
	c.Do(Add(1000))
	c.Do(Mul(2))
	c.Do(Sub(400))
	c.Do(math.Sqrt)
	c.Do(Sqrt())
	fmt.Println(c.acc)
}
