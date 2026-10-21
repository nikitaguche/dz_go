package main

import (
	"fmt"
	"math"
)

func main() {
	var useHeight = 1.8
var userKg float64= 100
var IMT = userKg  / math.Pow(useHeight, 2)

fmt.Print(IMT)

}
