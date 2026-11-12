package main

import (
	"errors"
	"fmt"
	"time"
)
type Bin struct{
	id string
	private  bool
	createdAt time.Time
	name string
}
func newBin (id string, private  bool, name string) (*Bin, error){
	if id == "" && name == ""{
		fmt.Println("Неверно введены все данные")
		return nil, errors.New("Invalid Data")
	}
	return &Bin{
		id : id,
		private : private,
		createdAt : time.Now(),
		name : name,
	}, nil

}

func main(){

}