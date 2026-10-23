package main

import "fmt"

func main() {
	usd := "usd"
	rub := "rub"
	val := userInput()
	res := rasch(val, usd, rub)
	fmt.Println(res)

}

func userInput() (val float64) {
	fmt.Scan(&val)
	return
}

func rasch(val float64, ish string, cel string) (res float64) {
	const usd_to_euro = 0.89
	const usd_to_rub = 81.10

	res = val / usd_to_euro * usd_to_rub
	return

}
