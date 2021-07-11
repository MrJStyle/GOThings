package main

import (
	"fmt"
)

type Alipay struct {
}

func (a *Alipay) CanUseFaceID() {

}

type Cash struct {
}

func (a *Cash) Stolen() {

}

type ContainCanUseFaceID interface {
	CanUseFaceID()
}

type ContainStolen interface {
	Stolen()
}

func printFeature(payMethod interface{}) {
	switch payMethod.(type) {
	// 可以刷脸
	case ContainCanUseFaceID:
		fmt.Printf("%T can use faceid\n", payMethod)
	// 可能被偷
	case ContainStolen:
		fmt.Printf("%T may be stolen\n", payMethod)
	}
}

func main() {
	printFeature(new(Alipay))

	printFeature(new(Cash))
}
