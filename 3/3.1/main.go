package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	defer func(){
		r := recover() 
		if r !=nil{
			fmt.Println("Recover", r)
		}
	}()
	for{
		IMT, err :=calcIMT(getUserInput())
		if err != nil{
			//fmt.Println(err)
			//continue
			panic("Не указан вес или рост")
		}
		resOut(IMT)
		fmt.Println("Хотите повторить? y/n")
		var isContinie string
		fmt.Scan(&isContinie)
		if isContinie == "n"{
			break
		}
	}
}


func resOut(IMT float64){
	result := fmt.Sprintf("Ваш ИМТ =  %f\n", IMT)
	switch{
	case IMT <=16:
		fmt.Println("Выраженный дефицит массы тела")
	case IMT>16 && IMT <=18.5:
		fmt.Println("Дефицит массы тела")
	case  IMT>18.5 && IMT<=24.99:
		fmt.Println("Норма")
	case IMT>24.99 && IMT<=30:
		fmt.Println("Избыточная масса тела")
	case  IMT>30 && IMT<=35:
		fmt.Println("Ожирение первой степени")
	case  IMT>35 && IMT<40:
		fmt.Println("Ожирение второй степени")
	case  IMT>=40:
		fmt.Println("Ожирение третьей степени")
				
	}


	/* if IMT <=16{
		fmt.Println("Выраженный дефицит массы тела")
	} else if IMT>16 && IMT <=18.5{
		fmt.Println("Дефицит массы тела")
	}else if IMT>18.5 && IMT<=24.99{
		fmt.Println("Норма")
	}else if IMT>24.99 && IMT<=30{
		fmt.Println("Избыточная масса тела")
	}else if IMT>30 && IMT<=35{
		fmt.Println("Ожирение первой степени")
	}else if IMT>35 && IMT<40{
		fmt.Println("Ожирение второй степени")
	}else if IMT>=40 {
		fmt.Println("Ожирение третьей степени")
	}*/
	fmt.Print(result)
}

func calcIMT(useHeight, userKg float64) (float64, error){
	if useHeight <=0 || userKg <=0{
		return 0, errors.New("Не указан вес или рост")
		//panic("Не указан вес или рост")
	}	
	const IMTPow = 2
		IMT := userKg  / math.Pow(useHeight/100, IMTPow)
		return IMT, nil
}


func getUserInput () (useHeight, userKg float64){
	fmt.Print(`Калькулятор ИМТ
Введите ваш рост в сантиметрах: `)
	fmt.Scan(&useHeight)
	fmt.Print("Введите ваш вес: ")
	fmt.Scan(&userKg)
	return 
}