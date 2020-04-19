package quicksort

func quick(data []int) {
	length := len(data)
	if length <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, length-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	quick(data[:head])
	quick(data[head+1:])
}
