package main

import "fmt"

func main() {
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Println(len(loons))

	fmt.Println(loons[1:3])

	// append values to the slice
	loons = append(loons, "latish")

	// double value range
	for index, value := range loons {
		fmt.Printf("Value %v at index %d\n", value, index)
	}

	nums := []int{16, 3, 34, 5, 12, -90}

	_max := nums[0]

	for _, value := range nums[1:] {
		if _max < value {
			_max = value
		}
	}

	fmt.Println(_max)
}
