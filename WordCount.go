package main

import (
	"strings"
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"regexp"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	strs := strings.Fields(s)

	for _,str :=range strs {
		res[strings.ToLower(str)]++
	}
	return res
}

func main() {
        reg := regexp.MustCompile(`\PP+`)
	file,err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	fs,err :=ioutil.ReadAll(file)
	for k,v := range WordCount(string(fs)) {
//		fmt.Println(k,v)
//		fmt.Println( reg.FindAllString(k, -1),v)
		k := reg.FindAllString(k, -1)
		fmt.Println( k[0],v)
	}	
}
