package sorting

import "fmt"

func MergeSort(in []int) {
	// find the midpoint
	n := len(in)
	mid := n / 2

	// handle base case of single element
	if n < 2 {
		return
	}

	// create new slices L and R
	L := make([]int, mid)
	R := make([]int, n-mid)

	// copy content from original into slices
	for i := 0; i <= mid-1; i++ {
		L[i] = in[i]
	}
	for i := mid; i <= n-1; i++ {
		R[i-mid] = in[i]
	}

	fmt.Println("Left sub:", L, "Right sub:", R)

	// call MergeSort on L
	MergeSort(L)
	// call MergeSort on  R
	MergeSort(R)
	// merge the two sorted slices
	Merge(in, L, R)
}

func Merge(O, L, R []int) []int {
	nL := len(L)
	nR := len(R)
	i, j, k := 0, 0, 0

	for {
		if i < nL && j < nR {
			if L[i] <= R[j] {
				// the unpicked in L is smaller than unpicked in R
				// over write original
				O[k] = L[i]

				k++ // go to the next position
				i++ // go to the next unpicked in L
				continue
			} else {
				// else the unpicked in R is smaller than unpicked in L
				// over write original
				O[k] = R[j]
				k++ // go to the next position
				j++ // go to the next unpicked in L
				continue
			}
		}
		// one of the subarray has exhausted, break
		break
	}

	// if the L subarray is not exhausted
	for {
		if i < nL {
			// copy the rest of L into Original
			O[k] = L[i]
			i++
			k++
			continue
		}
		break
	}

	// if the R subarray is not exhausted
	for {
		if j < nR {
			// copy the rest of R into Original
			O[k] = R[j]
			j++
			k++
			continue
		}
		break
	}
	return O
}
