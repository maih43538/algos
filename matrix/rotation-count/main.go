package main

import "fmt"

/*
How many times is a sorted list rotated?
eg. [11, 12, 15, 18, 2, 5, 6, 8]

Find the index of the lowest value, it will equal the number of times the list has been rotated.

Return the rotation count, return -1 if invalid input.

Property of the lowest value in a circularly sorted list:
* The values before and after it are larger

*/
func FindRotationCount(list []int) int {
	var n int = len(list)
	var low, high int = 0, n - 1
	// return -1 if there are duplicates in list

	mid := (low + high) / 2
	prev := (mid - 1) % n // modulo n in case mid is on the end
	next := (mid + 1) % n
	for {
		// case 1: no rotation
		if list[low] <= list[high] {
			return low
		}
		// case 2: mid is the lowest value
		if list[mid] <= list[prev] && list[mid] <= list[next] {
			fmt.Println("wee")
			return mid
		}

		// case 3: eliminate higher half
		if list[mid] <= list[high] {
			high = mid - 1
		}

		// case 4: eliminate lower half
		if list[mid] >= list[low] {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	l := []int{11, 12, 15, 18, 2, 5, 6, 8} // 4 rotations
	num := FindRotationCount(l)
	fmt.Printf("number of rotations for l list: %v \n", num)

	m := []int{1, 2, 3, 4, 5} // 0 rotations
	num = FindRotationCount(m)
	fmt.Printf("number of rotations for m list: %v \n", num)
}
