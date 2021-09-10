package main

import "fmt"

func main() {
	stocks := map[string]float64{
		"AMZN": 100.32,
		"TSLA": 200.3,
	}

	_, ok := stocks["AMZN"]
	if ok {
		fmt.Println("AMZN found in the map")
	} else {
		fmt.Println("AMZN not found in the map")
	}

	// add/update a value in the map
	stocks["LDN"] = 34.54

	// delete a value
	delete(stocks, "LDN")

	// iterate over the map
	for key, value := range stocks {
		fmt.Printf("%v: %v\n", key, value)
	}
}
