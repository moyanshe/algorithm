package classics

// BubbleAsort 冒泡排序（升序）
func BubbleAsort(values []int) []int {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	return values
}

//BubbleZsort 冒泡排序（降序）
func BubbleZsort(values []int) []int {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	return values
}
