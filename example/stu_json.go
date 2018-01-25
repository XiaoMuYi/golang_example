// json序列化与反序列化,实现简单高效;
package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"encoding/json"
	"io/ioutil"
)

type Stu struct {
	Id   int
	Name string
	Sex  string
	Tel  int
}
func main() {
	var id   int
	var name string
	var sex  string
	var tel  int

	var cmd  string
	var line string

	f := bufio.NewReader(os.Stdin)
	var stuinfo = make(map[string]Stu)
	for {
		fmt.Print(">")
		line,_ = f.ReadString('\n')
		fmt.Sscan(line,&cmd)

		switch cmd {
		case "list":
			for _,v := range stuinfo {
				fmt.Println(v.Id,v.Name,v.Sex,v.Tel)
			}
		case "add":
			fmt.Sscan(line,&cmd,&id,&name,&sex,&tel)
			if _,ok := stuinfo[name]; ok {
				fmt.Println("name is exits")
				continue
			}
			stuinfo[name] = Stu{Id:id,Name:name,Sex:sex,Tel:tel}
			fmt.Println("add done")
		case "save":
			w,err := json.Marshal(stuinfo)
			if err != nil {
				log.Fatal(err)
			}
			f, err := os.Create("/tmp/file.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			var d1 = []byte(w)
			if _, err := f.Write(d1); err != nil {
				log.Fatal(err)
			}
			f.Sync()
		case "load":
			f1,err := ioutil.ReadFile("/tmp/file.txt")
			if err != nil {
				log.Fatal(err)
			}
			json.Unmarshal(f1,&stuinfo)
			fmt.Println("load successful")
                case "exit":    
                        fmt.Println("Good bye!")
                        os.Exit(0)
                default:
                        fmt.Println("Usage: list | add | save | load ")
		}
	}
}
