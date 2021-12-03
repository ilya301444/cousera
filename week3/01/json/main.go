package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int
	Username string
	Phone    string
}



func init() {
	stuff = "ladder"
}

func main() {
	fn1()
	fn2()

}

var jsonStr2 = `[
				{"id": 17, "username":"iivan", "phone": 0},
				{"id": "17", "addres": "none", "company": "Mail.ru"}
]`
func fn2(){
	fmt.Println("Start function 2 ")
	data := []byte(jsonStr2)
	var user1 interface{}
	json.Unmarshal(data , &user1)
	fmt.Println("unpaced in empty interface: \n", user1)
	
}

var jsonStr = `{"id": 42, "username": "rvasily", "phone": "123"}`
var stuff string = "illusion"

func fn1(){
	fmt.Println(stuff)
	data := []byte(jsonStr)
	user := &User{}
	json.Unmarshal(data, user)
	fmt.Printf("struct: \n %v\n", user)
	user.Phone = "98765431"
	result, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json string: \n %s\n", result)
}


/*
когда можно восстанавливаться из функции
recover()
типа имеет ли место уровень вложенности

типа если defer с обработкой ошибок указать на несколько уровней больше
*/





