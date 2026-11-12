package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)
type account struct{
	login string
	password string
	url string
}
type accountWithStamp struct{
	createAt time.Time
	updateAt time.Time
	account
}
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!№%:,.;()")
func (acc account) outputPassword(){
	fmt.Println(acc.login, acc.password, acc.url)
}
func (acc *account) generatePassword(n int){
res := make([]rune, n)
for i:= range(res){
	res[i] = letterRunes[rand.IntN(len(letterRunes))]
}
acc.password = string(res)
}

/*
func newAccount(login, password, urlString string) (*account, error){

	_, err := url.ParseRequestURI(urlString)
	newAcc := &account{
		password : password,
		url : urlString,
		login : login,
	}
	switch{
	case (login == "")  && (err != nil):
		fmt.Println("Неверно введены все данные")
		return nil, errors.New("Invalid Data")
	case (login == ""):
		fmt.Println("Неверно введен логин")
		return nil, errors.New("Invalid login ")
	case (err != nil):
		fmt.Println("Неверный формат URL")
		return nil, errors.New("Invalid UR")
	case password == "":
		newAcc.generatePassword(12)
		return newAcc, nil
	default:
		return newAcc, nil
	}

}*/

func newAccountWithStamp(login, password, urlString string) (*accountWithStamp, error){

	_, err := url.ParseRequestURI(urlString)
	newAcc := &accountWithStamp{
		createAt : time.Now(),
		updateAt : time.Now(),
		account : account{
			password : password,
		    url : urlString,
			login : login,
		},
	}
	switch{
	case (login == "")  && (err != nil):
		fmt.Println("Неверно введены все данные")
		return nil, errors.New("Invalid Data")
	case (login == ""):
		fmt.Println("Неверно введен логин")
		return nil, errors.New("Invalid login ")
	case (err != nil):
		fmt.Println("Неверный формат URL")
		return nil, errors.New("Invalid UR")
	case password == "":
		newAcc.generatePassword(12)
		return newAcc, nil
	default:
		return newAcc, nil
	}

}

func main(){
	rand.IntN(10)
	login := promtData("Введите логин: ")
	//password := generatePassword(8)
	password := promtData("Введите пароль: ")
	url := promtData("Введите URL: ")
    myAccount , err := newAccountWithStamp(login, password, url)
	if err!= nil{
		return
	}

	

	myAccount.generatePassword(8)
	myAccount.outputPassword()
}

func promtData (promt string) string{
	fmt.Print(promt)
	var res string
	fmt.Scanln(&res)
	return res
}

