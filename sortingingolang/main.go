package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	// sortInt()
	// sortString()
	// sortSliceOfStruct()
	// sortCustomStruct()
	sortMapOnKey()

}

func sortMapOnKey() {
	m := map[string]int{"Ravi": 6, "Amit": 3, "Sumit": 6, "Neha": 9}

	keys := make([]string, 0, len(m))

	for v := range m {
		keys = append(keys, v)
	}
	sort.Strings(keys)

	for _, v := range keys {
		fmt.Println(v, m[v])
	}
}

// Sort custom data structures
func sortCustomStruct() {
	persons := []Person{
		{"Ravi", 23},
		{"Amit", 30},
		{"Neha", 27},
		{"Sumit", 20},
	}
	sort.Sort(ByAge(persons))
	fmt.Println(persons)
}

// Sort with custom comparator
func sortSliceOfStruct() {
	family := []struct {
		Name string
		Age  int
	}{
		{"Ravi", 26},
		{"Neha", 19},
		{"Amit", 28},
		{"Sumeet", 12},
	}

	sort.Slice(family, func(i, j int) bool {
		return family[i].Age > family[j].Age
	})

	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family)
}

// Sort a slice of strings
func sortString() {
	slice := []string{"neha", "sumeet", "ravi", "amit"}

	sort.Strings(slice)
	fmt.Println(slice)
}

// Sort a slice of ints

func sortInt() {
	slice := []int{4, 2, 1, 4, 7, 3, 10}

	sort.Ints(slice)
	fmt.Println(slice)
}
