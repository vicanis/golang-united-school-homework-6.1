package golang_united_school_homework

import "fmt"

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (r Rectangle) CalcPerimeter() float64 {
	return 2*r.Weight + 2*r.Height
}

func (r Rectangle) CalcArea() float64 {
	return r.Weight * r.Height
}

func (r Rectangle) GetType() string {
	return "rectangle"
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle (width %.2f, height %.2f): perimeter %.3f, area %.3f", r.Weight, r.Height, r.CalcPerimeter(), r.CalcArea())
}
