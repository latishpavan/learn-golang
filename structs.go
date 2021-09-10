package main

import "fmt"

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (t *Trade) Value() float64 {
	val := t.Price * float64(t.Volume)

	if t.Buy {
		return -val
	} else {
		return val
	}
}

func main() {
	t1 := Trade{"MSFT", 200, 45.3, false}
	fmt.Printf("%+v\n", t1)

	t2 := Trade{
		Symbol: "AAPL",
		Price:  234.3,
		Volume: 344,
		Buy:    true,
	}

	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Printf("%+v\n", t3)
	fmt.Println(t2.Value())
}
