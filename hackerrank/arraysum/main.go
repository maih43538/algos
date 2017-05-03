package main

import (
	"fmt"
	"log"
	"os"
)
// Input Format:
// The first line contains an integer, N, denoting the size of the array.
// The second line contains N space-separated integers representing the array's elements.

// Output Format:
// Print the sum of the array's elements as a single integer.

func main() {
	var n int
	total := 0
	_, err := fmt.Fscanln(os.Stdin, &n)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < n; i++ {
		var num int
		_, err := fmt.Fscan(os.Stdin, &num)
		if err != nil {
			log.Fatal(err)
		}
		total += num
	}
	fmt.Println(total)
}
