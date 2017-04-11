package sorting_test

import (
	"testing"

	"fmt"

	"github.com/itsanna/algos/sorting"
	. "github.com/smartystreets/goconvey/convey"
)

func TestQuicksort(t *testing.T) {
	Convey("Verify swap", t, func() {
		nums := []int{5, 0, 1, 4, 2}
		sorting.Swap(nums, 1, 0)
		So(nums[0], ShouldEqual, 0)
	})
	Convey("Verify partition", t, func() {
		nums := []int{5, 1, 0, 4, 2}
		i := sorting.Partition(nums, 0, 3)
		So(i, ShouldEqual, 2)

		nums = []int{7, 9, 10}
		i = sorting.Partition(nums, 0, 2)
		So(i, ShouldEqual, 2)

		nums = []int{6}
		i = sorting.Partition(nums, 0, 0)
		So(i, ShouldEqual, 0)
	})
	Convey("Verify quicksort", t, func() {
		nums := []int{5, 1, 0, 4, 2}
		err := sorting.QuickSort(nums, 0, 4)
		So(err, ShouldBeNil)
		So(nums[0], ShouldEqual, 0)
		fmt.Println(nums)

		nums = []int{99, 56, 10000, 0, 0, 4, 2}
		err = sorting.QuickSort(nums, 0, len(nums)-1)
		So(err, ShouldBeNil)
		So(nums[0], ShouldEqual, 0)
		fmt.Println(nums)
	})
}
