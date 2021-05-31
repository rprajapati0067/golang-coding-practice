package main

import (
	"fmt"
	"io"
	"os"
)

func manipulate(r io.Reader, w1, w2 io.Writer) {

}

func creatingFile(fileName string) (*os.File, error) {
	// Creating a file
	f, err := os.Create(fileName + ".txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return nil, err
	}
	return f, nil
}

func writingFile(file *os.File) (int, error) {
	// Writing string to file
	l, err := file.WriteString("123456")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return -1, err
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return l, nil
}

func main() {
	file, _ := creatingFile("hello")
	_, e := writingFile(file)

	if e != nil {
		fmt.Println(e)
	}

	helloFile, error := os.OpenFile("hello.txt", os.O_APPEND|os.O_RDONLY, 0644)
	if error != nil {
		fmt.Println(error)
	}
	defer helloFile.Close()

}
