package main

import (
	"fmt"
	
	
	
)

var x = I(T{}).ab()   // x has an undetected, hidden dependency on a and b
//var _ = sideEffect()  // unrelated to x, a, or b
var a = b
var b = 42

type I interface      { ab() []int }
type T struct{}
func (T) ab() []int   { return []int{a, b} }



func main(){
	var mas [120]int
	for i:= 0;i< 30;i++ {
		mas[i] = i
	}
	fmt.Println(mas[:15])
	
	
}
