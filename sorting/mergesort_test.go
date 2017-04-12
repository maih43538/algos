package sorting_test

import (
	"testing"

	"github.com/itsanna/algos/sorting"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeSort(t *testing.T) {
	Convey("Verify merge sort", t, func() {
		O := []int{2, 4, 1, 6, 8, 5, 3, 7}
		L := []int{1, 2, 4, 6}
		R := []int{3, 5, 7, 8}
		merged := sorting.Merge(O, L, R)
		So(merged, ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7, 8})

		O = []int{0, 0, 0, 0, 0}
		L = []int{67, 90}
		R = []int{8, 21, 79}
		merged = sorting.Merge(O, L, R)
		So(merged, ShouldResemble, []int{8, 21, 67, 79, 90})

		// merge
		O = []int{99, 4}
		L = []int{99}
		R = []int{4}
		merged = sorting.Merge(O, L, R)
		So(merged, ShouldResemble, []int{4, 99})
	})
	Convey("Verify merge sort", t, func() {
		in := []int{4}
		sorting.MergeSort(in)
		So(in, ShouldResemble, []int{4})

		in = []int{2, 4, 1, 6, 8, 5, 3}
		sorting.MergeSort(in)
		So(in, ShouldResemble, []int{1, 2, 3, 4, 5, 6, 8})

		in = []int{55, 6, 88, 55, -36, 36}
		sorting.MergeSort(in)
		So(in, ShouldResemble, []int{-36, 6, 36, 55, 55, 88})

		in = []int{100, 600, 999, 39999}
		sorting.MergeSort(in)
		So(in, ShouldResemble, []int{100, 600, 999, 39999})
	})
}
