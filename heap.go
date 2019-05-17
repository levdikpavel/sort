package sort

func Heap(a []int) {
	count := len(a)
	heapify(a, count)
	end := count - 1
	for ; end > 0; end-- {
		Swap(a, 0, end)

		siftDown(a, 0, end)
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
func siftDown(a []int, start, end int) {
	parentIdx := start
	for parentIdx < end {
		parent := a[parentIdx]
		childIdx1 := parentIdx*2 + 1
		childIdx2 := parentIdx*2 + 2
		swapIdx := -1
		swapIdx = checkChildSiftDown(a, parent, childIdx1, swapIdx, end)
		swapIdx = checkChildSiftDown(a, parent, childIdx2, swapIdx, end)
		if swapIdx >= 0 {
			Swap(a, parentIdx, swapIdx)
			parentIdx = swapIdx
		} else {
			return
		}
	}
}

func checkChildSiftDown(a []int, parent, childIdx, swapIdx, end int) int {
	if childIdx < end {
		child := a[childIdx]
		if child > parent {
			if swapIdx >= 0 && child < a[swapIdx] {
				return swapIdx
			}
			return childIdx
		}
	}
	return swapIdx
}
