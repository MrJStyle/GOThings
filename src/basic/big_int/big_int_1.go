package main

import (
	"fmt"
	"math/big"
)

////An Int represents a signed multi-precision integer.
////The zero value for an Int represents the value 0.
//type Int struct {
//	neg bool // sign
//	abs nat  // absolute value of the integer
//}

//// NewInt allocates and returns a new Int set to x.
//func NewInt(x int64) *Int {
//	return new(Int).SetInt64(x)
//}

func main() {
	big1 := new(big.Int).SetUint64(uint64(1000))
	fmt.Println("big1 is: ", big1)
	big2 := big1.Uint64()
	fmt.Println("big2 is: ", big2)
}
