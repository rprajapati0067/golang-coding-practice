package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Employee struct {
	Name string
	Age  int
	City string
}

func main() {

	// f, _ := creatingFile("sample")

	// bytes, _ := writingFile(f)

	// fmt.Println(bytes, "bytes written successfully")
	// appending()
	creatingAndWritingCSVFile()
	readingCSVFile()

}

func creatingAndWritingCSVFile() {
	// Creating CSV file
	file, err := os.Create("employee.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := []string{"name,age,city", "Garry,30,Newyork", "Jhone,40,Paris", "Adam,50,London", "William,20,Sydney"}

	for _, v := range data {
		// putting data into CSV file
		fmt.Fprintln(file, v)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func readingCSVFile() {
	empSlice := []Employee{}
	// Checking if file is present
	_, err := os.Stat("employee.csv")
	if err != nil {
		log.Fatalln(err)
	}
	// Opening CSV File
	csvFile, error := os.OpenFile("employee.csv", os.O_APPEND|os.O_RDONLY, 0644)
	if error != nil {
		fmt.Println(error)
	}
	defer csvFile.Close()

	// Reading CSV file
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for i, line := range csvLines {
		if i == 0 {
			continue
		}
		val, err := strconv.Atoi(line[1])

		if err != nil {
			fmt.Println(err)
		}
		// putting data into struct
		emp := Employee{
			Name: line[0],
			Age:  val,
			City: line[2],
		}
		// putting data into slice of employee
		empSlice = append(empSlice, emp)
		fmt.Println("Name: ", emp.Name, "Age: ", emp.Age, "City: ", emp.City)
		fmt.Println(empSlice)
	}
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
	l, err := file.WriteString("Hello World")
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

func writingLineByLine() {
	// Creating a file
	f, err := os.Create("lines")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}

	for _, v := range d {
		// Writing line by line to a file
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}

func appending() {
	// Opening a file for appending data to it
	f, err := os.OpenFile("lines", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handling is easy."
	// appending data to file
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")

}
