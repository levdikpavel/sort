package sort

import (
	"testing"
)

func BenchmarkBubble(b *testing.B) {
	a := getSlice1()
	for i := 0; i < b.N; i++ {
		Bubble(a)
		if !isSorted(a) {
			panic("a is not sorted!")
		}
	}
}
func aBenchmarkPrimitive(b *testing.B) {
	a := getSlice1()
	for i := 0; i < b.N; i++ {
		Primitive(a)
		if !isSorted(a) {
			panic("a is not sorted!")
		}
	}
}
func BenchmarkHeap(b *testing.B) {
	a := getSlice1()
	for i := 0; i < b.N; i++ {
		Heap(a)
		if !isSorted(a) {
			panic("a is not sorted!")
		}
	}
}
func BenchmarkQuick(b *testing.B) {
	a := getSlice1()
	for i := 0; i < b.N; i++ {
		Quick(a)
		if !isSorted(a) {
			panic("a is not sorted!")
		}
	}
}

var slice1 []int
var n1 = 50000

func getSlice1() []int {
	if len(slice1) == 0 {
		slice1 = generateSlice(n1)
	}
	sliceCopy := make([]int, n1, n1)
	copy(sliceCopy, slice1)
	return sliceCopy
}

func generateSlice(n int) []int {
	a := make([]int, n, n)
	for i := range a {
		a[i] = random.Int()
	}
	return a
}
