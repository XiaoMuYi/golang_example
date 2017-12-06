//usage: go run ReadFile.go /file/path

package main

import (
	"fmt"
	"os"
	"log"
//	"bufio"
	"io/ioutil"
)

/*
func main() {

	fi,err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
*/

func main() {
	file,err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fline,err := ioutil.ReadAll(file)
	fmt.Println(string(fline))
}
