package main

import (
	"fmt"
	"math"
)

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func double(val *float64) {
	*val *= 2
}

func squareRoot(val float64) (float64, error) {
	if val < 0 {
		return 0.0, fmt.Errorf("Cannot find square root of a negative value %f", val)
	}

	return math.Sqrt(val), nil
}

func main() {
	val := 2.2
	double(&val)
	fmt.Println(val)

	root, err := squareRoot(val)
	if err == nil {
		fmt.Println(root)
	} else {
		fmt.Printf("ERROR: %s\n", err)
	}
}
