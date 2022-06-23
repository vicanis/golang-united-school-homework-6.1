package golang_united_school_homework

import (
	"fmt"
	"math"
)

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcPerimeter() float64 {
	return 3 * t.Side
}

func (t Triangle) CalcArea() float64 {
	return (math.Pow(t.Side, 2) * math.Sin(math.Pi/3)) / 2
}

func (t Triangle) GetType() string {
	return "triangle"
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle (side %.3f): perimeter %.3f, area %.3f", t.Side, t.CalcPerimeter(), t.CalcArea())
}
