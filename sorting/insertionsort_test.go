package sorting_test

import (
	"testing"

	"github.com/itsanna/algos/sorting"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInsertionSort(t *testing.T) {
	Convey("Verify insertion sort", t, func() {
		in := []int{7}
		sorting.InsertionSort(in)
		So(in[0], ShouldEqual, 7)

		// best case
		in = []int{1, 2, 3, 4, 5, 6, 7}
		sorting.InsertionSort(in)
		So(in[0], ShouldEqual, 1)

		in = []int{6, 2, 4, 1, 5, 3}
		sorting.InsertionSort(in)
		for index, num := range in {
			So(num, ShouldEqual, index+1)
		}
	})
}
