package sort

func Swap(a []int, i, j int) {
	b := a[i]
	a[i] = a[j]
	a[j] = b
}
