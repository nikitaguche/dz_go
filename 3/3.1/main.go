package main

import (
	"fmt"
	"math"
)

func main() {
	
	useHeight := 1.8
	userKg :=  100.0
	IMT := userKg  / math.Pow(useHeight, 2)

fmt.Print(IMT)

}
