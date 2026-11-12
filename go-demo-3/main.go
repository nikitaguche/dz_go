package main

import "fmt"

func main(){
	var menu int8
	 m := map[string] string{}
	var key, value string
	for{
		fmt.Println("Меню:")
		fmt.Println("1 - Посмотреть закладки")
		fmt.Println("2 - Добавить закладку")
		fmt.Println("3 - Удалить закладку")
		fmt.Println("4 - Выход")
		
		fmt.Scan(&menu)
		switch menu{
		case 1:
			fmt.Println(m)
		case 2:
			fmt.Print("Введите ключ закладки: ")
			fmt.Scan(&key)
			fmt.Print("Введите значение закладки: ")
			fmt.Scan(&value)
			m[key] =value
			fmt.Println("Закладка добавлена!")
		case 3:
			fmt.Print("Введите ключ закладки: ")
			fmt.Scan(&key)
			delete(m, key)
			fmt.Println("Закладка удалена!")
		case 4:
			return
		default:
			fmt.Println("Введено неккоректное значение! Попробуйте еще раз")
		}
	}
}