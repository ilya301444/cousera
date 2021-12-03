package main

import (
	"fmt"
	"io"
	"os"
	//"path/filepath"
	"strings"
	"strconv"
)

var str string

func walkDir(out io.Writer, path string, printFiles bool, indentCod int) error {
	files, err := os.ReadDir(path)
	count := strings.Count(path, string(os.PathSeparator))
	indent := ""
	tempICod := indentCod
	
	for i := 0;i < count;i++ {
		if tempICod % 2 == 1 {
			
			indent =  "│	" + indent
		} else {
			indent =  "\t" + indent 
		}
		tempICod >>= 1
	}
	
	for i, f := range files {
		info, _ := f.Info()
		size := ""
		if info.Size() == 0 {
			size = " (empty)"
		} else {
			size = " (" + strconv.Itoa(int(info.Size())) + "b)"
		}
		 
		str = indent + "├───" + f.Name() + "\n"
		if f.IsDir() {
			if len(files)-1 == i {
				str = indent + "└───" + f.Name() + "\n"
				indentCod <<= 1
			} else {
				indentCod <<= 1
				indentCod += 1
			}
			fmt.Fprintf(out, str)
			walkDir(out, path + string(os.PathSeparator) + f.Name(), printFiles, indentCod)
			indentCod >>= 1
		} else {
			if len(files)-1 == i {
				str = indent + "└───" + f.Name() + size + "\n"
			} else {
				str = indent + "├───" + f.Name() + size + "\n"
			}
			fmt.Fprintf(out, str)
		}
	}
	
	return err
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	return walkDir(out, path, printFiles, 1)
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	fmt.Println(path)
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

