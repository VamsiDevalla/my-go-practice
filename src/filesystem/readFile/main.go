package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := openFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(readFile(fp)))
	os.Exit(0)
}

func openFile(filename string) (*os.File, error) {
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error while opening the file with name", filename, err)
		return nil, err
	}
	return fp, nil
}

func readFile(fp *os.File) []byte {
	stats, _ := fp.Stat()
	data := make([]byte, stats.Size())
	bytes, err := fp.Read(data)
	if err != nil {
		fmt.Println("some thing went wrong while reading the file", err)
	} else {
		fmt.Println("successfully read", bytes, "from the file")
	}
	fp.Close()
	return data
}
