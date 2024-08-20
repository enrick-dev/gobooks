package main

import (
	"errors"
	"fmt"
)

func sum(x, y int) (int, error) {
	if x+y > 5 {
		return x + y, nil
	}
	return 0, errors.New("x+y is less than 5")
}

func main() {
	x, err := sum(1, 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(x)
}
