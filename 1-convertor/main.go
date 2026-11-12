package main

import "fmt"

func main() {
	convert := map[string] map[string] float64{
		"USD": {
        	"EUR": 0.92,
        	"RUB": 90.5,
    	},
    	"EUR": {
        	"USD": 1.09,
        	"RUB": 98.2,
    	},
		"RUB": {
        	"USD": 0.012,
        	"EUR": 0.011,
    	},	
	}

	ish := userInputOne()
	val := userInputVal()
	conv := userInputTwo(ish)
	res := rasch(val, ish, conv, &convert)
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

func rasch(val float64, ish string, cel string, convert *map[string]map[string]float64) (res float64) {
	
	res = val * (*convert)[ish][cel]
	return

}
