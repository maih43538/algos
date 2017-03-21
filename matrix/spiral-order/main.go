package main

import "fmt"

/*

Print 2-D array in spiral order

numbers := [
  [2, 4, 6, 8],
  [5, 9, 12, 16],
  [1, 11, 5, 9],
  [3, 2, 1, 8],
]

=> 2, 4, 6, 8, 5, 9, 12, 16, 1, 11, 5, 9, 3, 2, 1, 8

*/

func PrintSpiral(list [][]int, rows, cols int) {
	// initialize variables
	var T, B, L, R, dir int = 0, rows - 1, 0, cols - 1, 0

	// the top-most must be less than or equal to bottom-most AND the left-most must be less than or equal to right-most
	for {
		if T >= B || L >= R {
			break
		}
		// 0 - traverse right (going Left to Right)
		if dir == 0 {
			fmt.Println("going right")
			for i := 0; i <= R; i++ {
				fmt.Println(list[T][i])
			}
			T++
			dir = 1
		}
		// 1 - traverse down (going Top to Bottom)
		if dir == 1 {
			fmt.Println("going down")
			for i := T; i <= B; i++ {
				fmt.Println(list[i][R])
			}
			R--
			dir = 2
		}
		// 2 - traverse left
		if dir == 2 {
			fmt.Println("going left")
			for i := R; i >= L; i-- {
				//fmt.Println("---> ", R, i)
				fmt.Println(list[B][i])
			}
			B--
			dir = 3
		}
		// 3 - traverse up
		if dir == 3 {
			for i := B; i >= T; i-- {
				fmt.Println("going up")
				fmt.Println(list[i][L])
			}
			L++
			dir = 0
		}
	}

	fmt.Printf("wooo")
}

func main() {
	l := [][]int{
		{2, 4, 6, 8},
		{5, 9, 12, 16},
		{1, 11, 5, 9},
		{3, 2, 1, 8},
	}

	m := 4
	n := 4
	PrintSpiral(l, m, n)
}
