package sort

func Quick(a []int) {
	Partitioning(a, 0, len(a)-1)
}

func Partitioning(a []int, i, j int) {
	intervalLength := j - i + 1
	if intervalLength <= 1 {
		return
	}
	p := random.Intn(intervalLength) + i
	pivot := a[p]
	left := i
	right := j
	for left < right {
		for a[left] < pivot {
			left++
		}
		for a[right] > pivot {
			right--
		}
		if left < right && a[left] == a[right] {
			left++
		} else {
			Swap(a, left, right)
		}
	}
	Partitioning(a, i, left-1)
	Partitioning(a, right+1, j)
}
