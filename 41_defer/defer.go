package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("41_defer/info.log")
	defer closeFile(f)
	writeFile(f, "Hello")
}

func createFile(dir string) *os.File {
	fmt.Println("Creating file", dir)
	f, err := os.Create(dir)
	if err != nil {
		panic("createFile error:" + err.Error())
	}
	return f
}

func writeFile(f *os.File, s string) {
	fmt.Println("Writing a file")
	fmt.Fprintln(f, s)
}

func closeFile(f *os.File) {
	fmt.Println("Closing a file")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:%v\n", err)
		os.Exit(1)
	}
}
