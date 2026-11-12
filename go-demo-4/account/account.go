package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)
type Account struct{
	login string
	password string
	url string
}
type accountWithStamp struct{
	createAt time.Time
	updateAt time.Time
	Account
}
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!№%:,.;()")
func (acc Account) OutputPassword(){
	fmt.Println(acc.login, acc.password, acc.url)
}
func (acc *Account) generatePassword(n int){
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

func NewAccountWithStamp(login, password, urlString string) (*accountWithStamp, error){

	_, err := url.ParseRequestURI(urlString)
	newAcc := &accountWithStamp{
		createAt : time.Now(),
		updateAt : time.Now(),
		Account : Account{
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