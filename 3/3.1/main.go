package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPow = 2

	useHeight := 1.8
	userKg :=  100.0
	IMT := userKg  / math.Pow(useHeight, IMTPow)

fmt.Print(IMT)

}
