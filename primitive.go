package sort

func Primitive(arr []int) []int {
	var result []int
	for len(arr) > 0 {
		minIdx := 0
		min := arr[minIdx]
		for i, a := range arr {
			if a < min {
				min = a
				minIdx = i
			}
		}
		result = append(result, min)
		copy(arr[minIdx:], arr[minIdx+1:])
		arr = arr[:len(arr)-1]
	}
	return result
}
