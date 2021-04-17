package main

import (
	"fmt"
)

var (
	reverse int64
)

func main() {
	//fmt.Println("Result: ", ReverseNum(123))
	//SwapNumber(10, 20)
	//fmt.Printf("Result: %v", IsPrime(11))
	//FibonacciSeries()
	//CountDigit(123)
	//FibonacciSeriesGoLangWay(10)
	vanEckSequence()
}

func vanEckSequence() {
	seq := make([]int, 100)
	for i := 0; i < 100; i++ {
		num := 0
		fmt.Printf("%d ", num)
		seq = append(seq, num)
		if len(seq)-2 >= 0 {
			for i := len(seq) - 2; i > 0; i-- {
				count := 1
				if seq[i] == num {
					num = count
					seq = append(seq, num)
				}else{
					seq = append(seq, 0)
				}
				count++

			}
		}else {
			seq = append(seq, num)
		}
	}
}

func FibonacciSeriesGoLangWay(num int) {
	n1 := 0
	n2 := 1
	//n3 := 0
	for i := 0; i < num; i++ {
		n1, n2 = n2, n2+n1
		fmt.Println(n1)
	}
}

func ReverseNum(number int64) int64 {
	for number != 0 {
		reverse *= 10
		reverse += number % 10
		number /= 10
	}
	return reverse
}

func SwapNumber(a int64, b int64) {
	b, a = a, b
	fmt.Printf("swaped number a %d and b %d", a, b)
}
func IsPrime(num int) bool {

	if num == 0 || num == 1 {
		return false
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
			break
		}

	}
	return true

}

func FibonacciSeries() {
	first := 0
	second := 0
	third := 1
	fmt.Printf("%d ", first)

	for i := 1; i <= 10; i++ {
		second = first
		first = third
		third = first + second
		fmt.Printf("%d ", third)
	}

}

func CountDigit(num int) {
	var count int
	if num == 0 {
		count++
	}

	for num != 0 {
		num /= 10
		count++
	}
	fmt.Println(count)
}
