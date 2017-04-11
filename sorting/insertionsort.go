package sorting

func InsertionSort(in []int) {
	// iterate from i=1 to i<=n-1
	for i := 1; i <= len(in)-1; i++ {
		key := i
		value := in[i]
		for {
			// if key <= 0 OR in[key-1] is less than value
			if key <= 0 || in[key-1] < value {
				// break out of inner loop
				break
			}
			// shift key - 1 over to key
			in[key] = in[key-1]
			// reassign key to key - 1
			key = key - 1
		}
		// place value in array at key
		in[key] = value
	}
}
