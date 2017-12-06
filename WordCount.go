package main

import (
	"strings"
	"os"
	"log"
	"io/ioutil"
	"fmt"
)

func WordCount(s string) map[string]int {
	strs := strings.Fields(s)
	res := make(map[string]int)

	for _,str :=range strs {
		res[strings.ToLower(str)]++
	}
	return res
}

func main() {
	file,err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	fs,err :=ioutil.ReadAll(file)
	for k,v := range WordCount(string(fs)) {
		fmt.Println(k,v)
	}	
}
