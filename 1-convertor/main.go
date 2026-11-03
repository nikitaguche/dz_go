package main

import "fmt"

func main() {
	
	ish := userInputOne()
	val := userInputVal()
	conv := userInputTwo(ish)
	res := rasch(val, ish, conv)
	fmt.Println("Результат конвертации:", res)

}
func userInputOne() (val string) {
	for{
		fmt.Print("Введите исходную валюту(USD, EUR, RUB): ")
		fmt.Scan(&val)
		if (val != "USD") && (val != "EUR") && (val != "RUB"){
			fmt.Println("Введено неккоректное значение!")
		}else{
			return
		}
	}
}

func userInputVal() (val float64) {
	for{
		fmt.Print("Введите количество валюты: ")
		fmt.Scan(&val)
		if val<=0{
			fmt.Println("Введено неккоректное значение!")
		}else{
			return
		}
	}
}

func userInputTwo(ish string) (val string) {
	for{
		switch{
		case ish == "USD":
			fmt.Print("Введите конечную валюту(EUR, RUB): ")
		case ish == "EUR":
			fmt.Print("Введите конечную валюту(USD, RUB): ")
		case ish == "RUB":
			fmt.Print("Введите конечную валюту(USD, EUR): ")
		default:
			fmt.Print("Введите конечную валюту(USD, EUR, RUB): ")
		}
		
		fmt.Scan(&val)
		if ((val != "USD") && (val != "EUR") && (val != "RUB")) || (val == ish){
			fmt.Println("Введено неккоректное значение!")
		}else{
			return
		}
	}
}

func rasch(val float64, ish string, cel string) (res float64) {
	const usd_to_euro = 0.89
	const usd_to_rub = 81.10
	const euro_to_rub = 91.44
	if(ish == "USD"){
		if(cel == "RUB"){
			res = val * usd_to_rub
		}else{
			res = val * usd_to_euro
		}
	}else if(ish == "EUR"){
		if(cel == "RUB"){
			res = val * euro_to_rub
		}else{
			res = val /  usd_to_euro
		}
	}else if ish == "RUB"{
		if(cel == "USD"){
			res = val / usd_to_rub
		}else{
			res = val /  euro_to_rub
		}
	}
	return

}
