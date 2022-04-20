package adder

import "fmt"

func Add(x, y int) (int, error) {
	if x >= 0 && y >= 0 {
		return x + y, nil
	}
	return 0, fmt.Errorf("x and y must be non-negative")
}
