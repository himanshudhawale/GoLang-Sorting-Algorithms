package merge

func mergeSort(array []int) []int {
	length:= len(array)
	if length< 2 {
		return array
	}

	key := length / 2
	left := mergeSort(array[0:key])
	right := mergeSort(array[key:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	tmp := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}
	tmp = append(tmp, left[i:]...)
	tmp = append(tmp, right[j:]...)
	return tmp
}
