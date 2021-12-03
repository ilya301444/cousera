package main

import (
	"fmt"
	"time"
	"contex"
	"math/rand"
)
// вывод при пережаче значения через канал её получит только какая то одна горутина
/*
func main(){
	cancel := make(chan struct{} )
	
	fn1 := func(cancel chan struct{}, str string){ 
		//defer fmt.Println("Gorutine 1 done")
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
				case <-cancel : 
					fmt.Println(str) 
					return 
				case tm := <-ticker.C : fmt.Println(tm)
			}
		}
	}
	
	go fn1(cancel, "Gorutint 1 is done")
	go fn1(cancel, "Gorutint 2 is done")
	go fn1(cancel, "Gorutint 3 is done")
	go fn1(cancel, "Gorutint 4 is done")
	fmt.Scanln()
	cancel <- struct{}{}
	fmt.Scanln()
}
*/

func worker(ctx context.Context, workerNum int){
	rand.Seed(time.Now().Unix())
	timer := time.NewTimer(rand.Intd(100)+ 10)
	
	
}

func main (){
	ctx, finish := context.WithCancel(context.Backgroung())
	for i := 0;i < 10;i++ {
		go worker(ctx, i)
	}
	
	
	
}


























