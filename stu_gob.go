//gob是Golang包自带的一个数据结构序列化的编码/解码工具,编码使用Encoder,解码使用Decoder;
//Encoder序列化数据是乱码,不能直接查看,gob另一个特性是原生支持序列化struct;
package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"encoding/gob"
)

type Stu struct {
	Id   int
	Name string
	Sex  string
	Tel  int
}


func main() {
	var id int
	var name string
	var sex string
	var tel int

	var cmd string
	var line string

	f := bufio.NewReader(os.Stdin)
	var stuinfo= make(map[string]Stu)
	for {
		fmt.Print(">")
		line, _ = f.ReadString('\n')
		fmt.Sscan(line, &cmd)

		switch cmd {
		case "list":
			for _, v := range stuinfo {
				fmt.Println(v.Id, v.Name, v.Sex, v.Tel)
			}
		case "add":
			fmt.Sscan(line, &cmd, &id, &name, &sex, &tel)
			if _, ok := stuinfo[name]; ok {
				fmt.Println("name is exits")
				continue
			}
			stuinfo[name] = Stu{Id: id, Name: name, Sex: sex, Tel: tel}
			fmt.Println("add done!")
		case "save":
			fi, err := os.Create("/tmp/db.txt")
			if err != nil {
				log.Fatal(err)
			}
			enc := gob.NewEncoder(fi)
			if err := enc.Encode(stuinfo); err != nil {
				log.Fatal(err)
			}
			fmt.Println("save successful!")
		case "load":
			fo,err := os.Open("/tmp/db.txt")
			if err != nil {
				log.Fatal(err)
			}
			dec := gob.NewDecoder(fo)
			if err := dec.Decode(&stuinfo); err != nil { //反序列化数据一定是到指针,否则会报错(json也一样);
				log.Fatal(err)
			}
			fmt.Println("load successful!")
		case "exit":
			fmt.Println("Good bye!")
			os.Exit(0)
		default:
			fmt.Println("Usage: list | add | save | load ")
		}
	}
}
