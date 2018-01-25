package main

import (
	"os"
	"io"
	"io/ioutil"
	"bufio"
	"log"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileExist(filename string)  (bool) {
	var exist= true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func main()  {
	var writeString = "测试"
	var filename = "/tmp/outfile01.txt"
	var f *os.File
	var err1 error
	var x int = 1000
//	fmt.Println("\t第一种方式：使用 io.writestring写入文件")
	if checkFileExist(filename) {
		f,err1 = os.OpenFile(filename,os.O_APPEND|os.O_WRONLY,0664)
//		fmt.Println("文件已存在")
	} else {
		f,err1 = os.Create(filename)
//		fmt.Println("文件不存在,已创建文件")
	}
	check(err1)
	for i := 0; i < x; i++ {
		_,err1 := io.WriteString(f,writeString)
		check(err1)
	}

//	fmt.Println("\t第二种方式：使用 ioutil.WriteFile 写入文件")
	for i := 0; i < x; i++ {
		var d1 = []byte(writeString)
		err2 := ioutil.WriteFile("/tmp/outfile02.txt",d1,0664) //ioutil.WriteFile每次执行都会清空原有内容
		check(err2)
	}

//	fmt.Println("\t第三种方式：使用 File(Write,WriteString) 写入文件")
	f, err3 := os.Create("/tmp/outfile03.txt")  //创建文件
	check(err3)
	defer f.Close()

	for i := 0; i < x; i++ {
		var d1 = []byte(writeString)
		_, err3 := f.Write(d1)  //写入文件(字节数组)
		check(err3)
		f.Sync()
	}

//	fmt.Println("\t第四种方式：使用  bufio.NewWriter 写入文件")
	for i := 0; i < x; i++ {
		f, err := os.Create("/tmp/outfile04.txt")  //创建文件
		w := bufio.NewWriter(f)  //创建新的 Writer 对象
		_, err3 := w.WriteString(writeString)
		if err != nil {
			log.Fatal(err3)
		}
		w.Flush()
		f.Close()
	}
}
 
