package sort

func Bubble(a []int) {
	for i, x := range a {
		for j := i + 1; j < len(a); j++ {
			y := a[j]
			if x > y {
				Swap(a, i, j)
				x = y
			}
		}
	}
}
