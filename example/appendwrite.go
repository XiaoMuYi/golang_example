/*
o,打开文件,写入内容,关闭文件。如此重复多次,速度缓慢且稳定;
1,打开文件,写入内容,defer关闭文件。如此重复多次,速度一般;
2,打开文件,重复多次写入内容defer关闭文件,输入速度较快;
*/
package main

import (
	"fmt"
	"time"
	"os"
)

func benchmarkFileWrite(filename string,n int,index int) (d time.Duration) {
	v := "You say, I'm good boy!"
	os.Remove(filename)

	t0 := time.Now()
	switch index {

	case 0:
		for i :=0; i < n;i++  {
			file,err := os.OpenFile(filename,os.O_WRONLY|os.O_APPEND|os.O_CREATE,066)
			if err != nil {
				fmt.Println(index,i,"open file failed.",err.Error())
				break
			}
			file.WriteString(v)
			file.WriteString("\n")
			file.Close()
		}

	case 1:
		for i := 0; i < n; i++ {
			file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(index, i, "open file failed.", err.Error())
				break
			}
			defer file.Close()

			file.WriteString(v)
			file.WriteString("\n")
		}
	case 2:
		file,err := os.OpenFile(filename,os.O_WRONLY|os.O_APPEND|os.O_CREATE,066)
		if err != nil {
			fmt.Println(index,"open file faild.",err.Error())
			break
		}
		defer file.Close()

		for i := 0; i < n;i ++ {
			file.WriteString(v)
			file.WriteString("\n")
		}
	}

	t1 := time.Now()
	d = t1.Sub(t0)
	fmt.Printf("time of way(%d)=%v\n",index,d)
	return d

}
func main() {
	const k, n int = 3, 1000
	d := [k]time.Duration{}
	for i := 0; i < k; i++ {
		d[i] = benchmarkFileWrite("/tmp/benchmarkFile.txt", n, i)
	}

	for i := 0; i < k-1; i++ {
		fmt.Printf("way %d cost time is %6.1f times of way %d\n", i, float32(d[i])/float32(d[k-1]), k-1)
	}
}
