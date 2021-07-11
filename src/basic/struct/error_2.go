package main

import (
	"fmt"
	"math"
)

type dualError struct {
	Num     float64
	problem string
}

func (e dualError) Error() string {
	return fmt.Sprintf("Wrong!!!, because \"%f\" is a %s", e.Num, e.problem)
}

func Sqrt2(f float64) (float64, error) {
	if f < 0 {
		return -1, dualError{Num: f, problem: "negative number"}
	}
	return math.Sqrt(f), nil
}

func main() {
	result, err := Sqrt2(-13)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
