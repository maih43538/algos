package sorting_test

import (
	"testing"

	"github.com/itsanna/algos/sorting"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBubblesort(t *testing.T) {
	Convey("Verify bubblesort", t, func() {
		in := []int{4}
		sorting.BubbleSort(in)
		So(in[0], ShouldEqual, 4)

		in = []int{1, 2, 3, 4, 5, 6, 7}
		sorting.BubbleSort(in)
		So(in, ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7})

		in = []int{4, 3, 1}
		sorting.BubbleSort(in)
		So(in[0], ShouldEqual, 1)
		So(in[1], ShouldEqual, 3)
		So(in[2], ShouldEqual, 4)

		in = []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
		sorting.BubbleSort(in)
		for index, val := range in {
			So(val, ShouldEqual, index)
		}
	})

	Convey("Verify bubblesort II", t, func() {
		in := []int{4}
		sorting.BubbleSortII(in)
		So(in[0], ShouldEqual, 4)

		in = []int{1, 2, 3, 4, 5, 6, 7}
		sorting.BubbleSortII(in)
		So(in, ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7})

		in = []int{4, 3, 1}
		sorting.BubbleSortII(in)
		So(in[0], ShouldEqual, 1)
		So(in[1], ShouldEqual, 3)
		So(in[2], ShouldEqual, 4)

		in = []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
		sorting.BubbleSortII(in)
		for index, val := range in {
			So(val, ShouldEqual, index)
		}
	})
}
