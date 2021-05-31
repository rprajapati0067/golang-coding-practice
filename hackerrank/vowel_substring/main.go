package main

import "fmt"

var exists = struct{}{}

type Set struct {
	m    map[interface{}]struct{}
	size int
}

func newSet(len int) *Set {
	s := &Set{}
	s.m = make(map[interface{}]struct{}, len)
	return s
}

func (s *Set) add(value interface{}) {
	s.m[value] = exists
	s.size++
}

func (s *Set) contains(value interface{}) bool {
	_, ok := s.m[value]
	return ok
}

func (s *Set) delete(value interface{}) {
	delete(s.m, s.m[value])
	s.size--
}

func isVowel(st string) bool {

	if st != "" && (st == "a" || st == "e" || st == "i" || st == "o" || st == "u") {
		return true
	}
	return false
}

func main() {
	res := vowelSubStr("aeioaexaaeuiou")
	fmt.Println(res)

}

// vowel substring hacker rank (question asked in Junglee)

func vowelSubStr(st string) int {
	count := 0
	s := newSet(10)

	for i := 0; i < len(st); i++ {
		for j := i; j < len(st); j++ {
			if isVowel(string(st[j])) == false {
				break
			}

			s.add(string(st[j]))

			if s.size == 5 {
				count++
			}
		}
		s = newSet(10)
	}

	return count

}
