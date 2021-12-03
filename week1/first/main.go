package main

import (
    "fmt"
    "io/ioutil"
     "log"
	 "os"
)

func main() {
	os.ReadDir(".")
    files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }
 
    for _, f := range files {
            fmt.Println(f.Name())
    }
}