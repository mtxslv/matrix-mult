package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Sample struct {
	n uint8
	strassen float64
	regular float64
}
 	

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main(){
	// https://gobyexample.com/writing-files
	f, err := os.Create("/tmp/dat2")
    check(err)
}