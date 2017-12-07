package main

import (
	"strings"
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"regexp"
	"sort"
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

	type kv struct {
		key string
		value int
	}
	
	//关于map排序,可以查阅相关资料,由于这里需要根据 value 进行排序,所以用到结构体;
	var ss  []kv
	fs,err :=ioutil.ReadAll(file)
	for k,v := range WordCount(string(fs)) {
		k := reg.FindAllString(k, -1)
		ss = append(ss,kv{k[0],v})
	}
	sort.Slice(ss, func(i, j int) bool { //sort.Slice在go 1.8 提供的功能;
		return ss[i].value > ss[j].value
	})

	fmt.Println("-------Frequency statistics and sorting of text words-----")	
	for _,kv := range ss{
		fmt.Printf("\t%d  %s\n",kv.value,kv.key)
	}
}
