package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Go")

	readFile()
	readDir()

}

func readDir() {
	fmt.Println("readDir")
	dir, err := os.Open("/")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfo, err := dir.ReadDir(-1)
	if err != nil {
		return
	}

	for _, fi := range fileInfo {
		fmt.Println(fi.Name())
	}
}

func readFile() {
	fmt.Println("readFile")
	file, err := os.Open("tst.txt")
	if err != nil {
		return
	}
	defer file.Close()

	state, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, state.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}
