package sort

func Heap(a []int) {
	count := len(a)
	heapify(a, count)
	end := count - 1
	for ; end > 0; end-- {
		Swap(a, 0, end)
		heapify(a, end)
	}
}

func heapify(a []int, count int) {
	end := 1
	for ; end < count; end++ {
		siftUp(a, 0, end)
	}
}

func siftUp(a []int, start, end int) {
	childIdx := end
	for childIdx > start {
		parentIdx := (childIdx - 1) / 2
		if a[parentIdx] < a[childIdx] {
			Swap(a, parentIdx, childIdx)
			childIdx = parentIdx
		} else {
			return
		}
	}
}
