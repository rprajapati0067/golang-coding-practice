package main

import "fmt"

func convertingDiagonalArrayToOneD() {
	var res [5]int
	a := [5][5]int{
		{3, 0, 0, 0, 0},
		{0, 7, 0, 0, 0},
		{0, 0, 4, 0, 0},
		{0, 0, 0, 9, 0},
		{0, 0, 0, 0, 6},
	}
	for i, _ := range a {
		for j, _ := range a {
			if i == j {
				res[i] = a[i][j]
			}
		}
	}
	fmt.Println(res)
}

func (m *Matrix) set(i int, j int, x int) {
	if i == j {
		m.A[i-1] = x
	}
}
func (m *Matrix) get(i int, j int) int {
	if i == j {
		return m.A[i-1]
	} else {
		return 0
	}
}

func (m *Matrix) display() {
	for i := 0; i < m.n; i++ {
		for j := 0; j < m.n; j++ {
			if i == j {
				fmt.Printf("%d", m.A[i-1])
			} else {
				fmt.Printf("0")
			}
		}
	}
}

type Matrix struct {
	A [10]int
	n int
}

func main() {
	//convertingDiagonalArrayToOneD()
	var m Matrix
	m.n = 4
	m.set(1,1,5)
	m.set(2,2,8)
	m.set(3,3,8)
	m.set(4,4,8)

	//m.display()

}
