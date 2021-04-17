package main

import (
	"github.com/rprajapati0067/golang-coding-practice/sets"
)

func sockMerchant(n int32, ar []int32) int32 {
	s := sets.NewSet()
	var pairs int32
	for i := 0; int32(i) < n; i++ {

		if !s.Contains(ar[i]) {
			s.Add(ar[i])
		} else {
			pairs++
			s.Remove(ar[i])
		}
	}
	return pairs
}

