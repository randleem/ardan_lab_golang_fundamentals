package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println(Relu(7))
	fmt.Println(Relu(-7))
	fmt.Println(Relu(7.7))
	fmt.Println(Relu2(time.February))
	m, err := NewMatric[float64](10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(m)
	fmt.Println(m.At(3, 2))

	fmt.Println(Max([]int{3, 1, 2}))     // 3 <nil>
	fmt.Println(Max([]float64{3, 1, 2})) // 3 <nil>
	fmt.Println(Max[int](nil))           // 0 Max of empty slice
}

func Max[T Number](s []T) (T, error) {
	if len(s) == 0 {
		var zero T
		return zero, errors.New("Max of empty slice")
	}
	m := s[0]
	for _, v := range s[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil
}

// Exercise:" Write Max Function for ints or floats"
// DOnt use inbuilt max function

func (m *Matrix[T]) At(row, col int) T {
	i := (row * m.Cols) + col
	return m.data[i]
}

func NewMatric[T Number](rows, cols int) (*Matrix[T], error) {
	if rows <= 0 || cols <= 0 {
		return nil, fmt.Errorf("bad dimenstions: %d/%d", rows, cols)
	}
	m := Matrix[T]{
		Rows: rows,
		Cols: cols,
		data: make([]T, rows*cols),
	}
	return &m, nil
}

type Matrix[T Number] struct {
	Rows int
	Cols int
	data []T
}

type Number interface {
	~float64 | ~int
}

func Relu[T int | float64](i T) T {
	if i < 0 {
		return 0
	}
	return i
}

func Relu2[T Number](i T) T {
	if i < 0 {
		return 0
	}
	return i
}
