package sort

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func Swap(a []int, i, j int) {
	b := a[i]
	a[i] = a[j]
	a[j] = b
}

func isSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}
