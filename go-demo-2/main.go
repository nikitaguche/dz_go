package main

import "fmt"

func main(){

	transacrions := []float64{}
	for{
		fmt.Print("Введите транзакцию (n для выхода): ")
		temp := Inp()
		if temp == 0{
			break
		}
		transacrions = append(transacrions, temp)
	}
	var sum float64
	for _, val := range transacrions{
		sum += val
	}
	fmt.Println(transacrions)
	fmt.Println("Сумма баланса: ", sum)
}

func Inp()(float64){
	var temp float64
	fmt.Scan(&temp)
	return temp
}