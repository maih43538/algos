package sorting

func BubbleSort(in []int) {
	if len(in) <= 1 {
		return
	}

	for k := 1; k <= len(in); k++ {
		// k keeps track of the sorted portion
		swapped := false
		for i := 0; i < len(in)-k; i++ {
			// iterate up to but not including sorted portion
			if in[i] > in[i+1] {
				Swap(in, i, i+1)
				swapped = true
			}
		}
		if !swapped {
			// if there were no swaps, do not need to conduct anymore passes
			// break to avoid redundant passes once the array is sorted
			break
		}
	}
}

func BubbleSortII(in []int) {
	// shorter version
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(in)-1; i++ {
			if in[i+1] < in[i] {
				Swap(in, i, i+1)
				swapped = true
			}
		}
	}
}
