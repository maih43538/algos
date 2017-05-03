package shuffle

import "math/rand"

// in-place shuffle

func Shuffle(in []int) {
	if len(in) <= 1 {
		return
	}

	// walk through beginning to end
	for indexWereChoosingFor := 0; indexWereChoosingFor <= len(in)-1; indexWereChoosingFor++ {
		// generate random number from 0 through length - 1
		rand := rand.Intn(len(in) - 1)

		if (rand != indexWereChoosingFor) {
			// swap
			temp := in[indexWereChoosingFor]
			in[indexWereChoosingFor] = in[rand]
			in[rand] = temp
		}
	}
}
