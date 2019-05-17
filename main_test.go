package sort

import (
	"testing"
)

func TestPrimitive(t *testing.T) {
	a := []int{2, 5, 1, 7, 6, 8, 1, 4}
	b := Primitive(a)
	if !isSorted(b) {
		t.Errorf("b is not sorted!\n%v", b)
	}
}

func TestBubble(t *testing.T) {
	a := []int{2, 5, 1, 7, 6, 8, 1, 4}
	Bubble(a)
	if !isSorted(a) {
		t.Errorf("Slice is not sorted!\n%v", a)
	}
}
func TestHeap(t *testing.T) {
	a := []int{2, 5, 1, 7, 6, 8, 1, 4}
	Heap(a)
	if !isSorted(a) {
		t.Errorf("Slice is not sorted!\n%v", a)
	}
}

func TestQuick(t *testing.T) {
	a := []int{2, 5, 1, 7, 6, 8, 1, 4}
	Quick(a)
	if !isSorted(a) {
		t.Errorf("Slice is not sorted!\n%v", a)
	}
}
