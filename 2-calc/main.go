package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){

	oper := inputOper()
	mas := inputMas()
	result :=  calcRes(oper, mas)
	fmt.Printf("Результат операции %s равен %f", oper, result)

}

func inputOper() (val string){
	for {
		fmt.Print("Введите желаемую операцию(AVG, SUM, MED): ")
		fmt.Scan(&val)
		if val == "AVG" || val == "SUM" || val == "MED"{
			return
		}else{
			fmt.Println("Введено неккоректное значение!")
		}
	}
}
func inputMas() (val []int){
	var str string
	fmt.Println("Введите числа одной строкой через запятую (БЕЗ пробелов): ")
	fmt.Scan(&str)

	parts := strings.Split(str, ",")
	for _, value := range parts {
		
		num, err := strconv.Atoi(value)
		if err != nil {
			continue
		}
		val = append(val, num)
	}
	return
}

func calcRes(oper string, mas []int) (res float64){
	switch{
	case oper == "AVG":
		if len(mas) == 0 {return 0}
		sum := 0
		for _, val := range(mas){
			sum += val
		}
		return float64(sum) / float64(len(mas))

	case oper == "SUM":
		if len(mas) == 0 {return 0}
		sum := 0
		for _, val := range(mas){
			sum += val
		}
		return float64(sum)

	case oper == "MED":
		if len(mas) % 2 == 1{
			return float64(mas[((len(mas) + 1) / 2) -1])
		}else if len(mas) % 2 == 0{
			return float64((mas[len(mas)/2] + mas[(len(mas) + 1)/2] )/ 2)
		}

	default:
		fmt.Println("Введено неккоректное значение операции!")
	}
	return
}