package main

import (
	"fmt"
	"log"
	"os"
	"math"
)

// Input Format:
// The first line contains a single integer, N. The next N lines denote the matrix's rows,
// with each line containing space-separated integers describing the columns.

// Output Format:
//  Print the absolute difference between the two sums of the matrix's diagonals as a single integer.

func main() {
	var n int
	_, err := fmt.Fscanln(os.Stdin, &n)
	if err != nil {
		log.Fatal(err)
	}

	array2d := [][]float64{}
	for i := 0; i < n; i++ {
		array1d := []float64{}
		for j := 0; j < n; j++ {
			var num float64
			_, err := fmt.Fscan(os.Stdin, &num)
			if err != nil {
				log.Fatal(err)
			}
			array1d = append(array1d, num)
		}
		array2d = append(array2d, array1d)
	}

	var diag1, diag2 float64
	diag1 = 0
	diag2 = 0
	for i := 0; i < n; i++ {
		diag1 += array2d[i][i]
		diag2 += array2d[i][n-(i+1)]

	}
	total := math.Abs(diag1 - diag2)
	fmt.Println(total)
}
