package main

import (
	"fmt"
	//"io"
	"os"
	//"time"
	"strings"
)



func worker() {
	for i := 0;i < 100;i++ {
		go func(){
			for {
				fmt.Fprintf(os.Stdout, "1")
			}
		}()
	}
	
}


func worker1(){
	fmt.Println("Start worker 1")
	for i:=0;i<100;i++ {
		fmt.Fprintf(os.Stdout, "q")		
	}
	fmt.Fprintf(os.Stdin, "\n")	
}

func worker2(){
	fmt.Println("Start worker 2 ")
	s := ""
	fmt.Fscan(os.Stdin, &s)
	fmt.Println(strings.Count(s, "q"))
	
}

func main(){
	
	
	go worker2()
	go worker1()
	
	s := ""

	fmt.Scan(&s)
	//worker()
	/*
	time.Sleep(time.Millisecond)
	fmt.Printf("%T", os.Stdin)
	var s string
	fmt.Scanln(os.Stdin, &s)
	if s == "" {
		fmt.Print("string is", s)
		return
	}
	io.Copy(os.Stdout, os.Stdin)	
	fmt.Println(1234)
	*/
	
}

