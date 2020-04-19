package insertionsort

func InsertSort(array []int) {
	length:= len(array)
	if length<= 1 {
		return
	}
	for i := 1; i < length; i++ {
		value := array[i]
		j := i - 1
		for ; j >= 0; j-- {
			if array[j] > value {
				array[j+1] = array[j]
			} else {
				break
			}
		}
		array[j+1] = value
	}
}
