package main

import "fmt"

func main() {
	const usd_to_euro = 0.89
	const usd_to_rub = 81.10

	convert := 1 / usd_to_euro * usd_to_rub

	fmt.Println(convert)
	
}
