package main

import (
	"fmt"
	"sync"

)


func main(){
	ch := make(chan string, 100)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ch chan string) {
		for i := 0;i < 100;i++ {
			ch <- "a"
			fmt.Print("w")
		}
		wg.Done()
	}(ch)
	
	wg.Add(1)
	go func(ch chan string){
		for i := 0;i < 100;i++ {
			fmt.Print(<-ch)
		}
		wg.Done()
	}(ch)
	
	wg.Wait()
	
	fmt.Println("Arcada lalala!!!")
	
}


















