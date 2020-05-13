package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := createFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	numBytes, err := writeTo(fp, os.Args[2])
	if err != nil {
		os.Exit(2)
	}
	fmt.Println("Wrote", numBytes, "into the file")
	os.Exit(0)
}

func createFile(fileName string) (*os.File, error) {
	fp, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error while creating a file with name ", fileName, err)
		return nil, err
	}
	return fp, nil
}

func writeTo(fp *os.File, content string) (int , error) {
	bytes := []byte(content)
	numBytes, err := fp.Write(bytes)
	if err != nil {
		fmt.Println("error while writing to the file")
		return 0, err
	}
	fp.Close()
	return numBytes, nil
}