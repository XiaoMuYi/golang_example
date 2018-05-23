package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Student struct {
	Id   int
	Name string
}
var students = make(map[string]Student)

func main() {
	var id int
	var cmd string
	var name string
	var file = "D:\\WeGame\\stu_manager.txt"
	//NewScanner创建并返回一个从r读取数据的Scanner，默认的分割函数是ScanLines(行);
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		scanner.Scan()
		line := scanner.Text() //Text 将最后一次扫描出的"匹配部分"作为字符串返回(返回副本);
		fmt.Sscan(line, &cmd)

		switch cmd {
		case "add":
			fmt.Sscan(line, &cmd, &id, &name)
			add(id, name)
		case "list":
			list()
		case "save":
			save(file)
		case "reload":
			reload(file)
		case "exit":
			os.Exit(0)
		case "del":
			fmt.Sscan(line,&cmd,&id,&name)
			del(name)
		case "update":
			fmt.Sscan(line,&cmd,&id,&name)
			updete(name,id)
		default:
			fmt.Println("Please enter add/list/save/reload!")

		}
	}
}

func add(id int, name string) {
	if _,ok := students[name];ok {
		fmt.Printf("name already existed",name)
	} else {
		students[name] = Student{Id:id,Name:name}
	}
}

func list() {
	for _, v := range students {
		fmt.Println(v.Id, v.Name)
	}
}

func save(filepath string ) {
	file, _ := os.Create(filepath)
	defer file.Close()

	buf01, err := json.Marshal(students)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(string(buf01))
	file.Sync()
}

func reload(filename string)  {
	f, err := ioutil.ReadFile("D:\\WeGame\\stu_manager.txt")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(f, &students)
}

func del(name string) {
	if _,ok := students[name];ok {
		delete(students,name)
		fmt.Printf("Delete name is %s",name)
		for _,v := range students {
			fmt.Println(v.Id,v.Name)
		}
	} else {
		fmt.Printf("name not exist")
	}
}

func updete(name string,id int) {
	if _,ok := students[name]; ok {
		students[name] = Student{id,name}
		fmt.Printf("update %s to %d",name,id)
	} else {
		fmt.Printf("%s don't exist",name)
	}
}
