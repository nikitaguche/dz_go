package main

import (
	"fmt"
	"math/rand/v2"
	"demo/password/account"
	"github.com/fatih/color"
	"demo/password/files"
)

func main(){
	rand.IntN(10)
	login := promtData("Введите логин: ")
	//password := generatePassword(8)
	password := promtData("Введите пароль: ")
	url := promtData("Введите URL: ")
    myAccount , err := account.NewAccountWithStamp(login, password, url)
	if err!= nil{
		return
	}

files.ReadFile()
	myAccount.OutputPassword()
}

func promtData (promt string) string{
	color.Red(promt)
	var res string
	fmt.Scanln(&res)
	return res
}

