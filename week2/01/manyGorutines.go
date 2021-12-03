package main

import (
	"fmt"
	//"time"
)

func main(){
	
	chan1 := make(chan int)
	/*
	chan2 := make(chan int)
	chan3 := make(chan int)
	chan4 := make(chan int)
	chan5 := make(chan int)
	chan6 := make(chan int)
	chan7 := make(chan int)
	
	/*
	fn1 :=  func(ch chan int){
		time.Sleep(6*time.Second)
		ch <- 1
	}
	go fn1(chan1)
	go fn1(chan2)
	go fn1(chan3)
	go fn1(chan4)
	go fn1(chan5)
	go fn1(chan6)
	go fn1(chan7)
	*/
	
	get(chan1)
	
	fmt.Println( <-chan1)
	
	
}

func get(ch chan int) int{
	ch <- 1
	return 1
	
}