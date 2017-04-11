package sorting

import "fmt"

func QuickSort(in []int, start, end int) error {
	if start < end {
		index := Partition(in, start, end)
		QuickSort(in, start, index-1)
		QuickSort(in, index+1, end)

	}
	return nil
}

func Swap(slice []int, a, b int) {
	temp := slice[a]
	slice[a] = slice[b]
	slice[b] = temp
}

func Partition(in []int, start, end int) int {
	if start > end {
		fmt.Println("invalid start index:", start, "; end index:", end)
		return 0
	}
	pivot := in[end]
	partitionIndex := start
	for i := start; i < end; i++ {
		if in[i] <= pivot {
			// swap if element is less than pivot
			Swap(in, i, partitionIndex)
			partitionIndex++
		}
	}

	Swap(in, partitionIndex, end) // swap pivot with element at partition index
	return partitionIndex
}
