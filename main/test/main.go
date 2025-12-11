package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T\n", fileObj)
	fileInfo, err := fileObj.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
}
