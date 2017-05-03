package shuffle_test

import (
	"testing"

	"github.com/itsanna/algos/interviewcake/shuffle"
	. "github.com/smartystreets/goconvey/convey"
)

func TestShuffle(t *testing.T) {
	Convey("Verify shuffle works", t, func() {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8}
		shuffle.Shuffle(in)
		So(in, ShouldNotResemble, []int{1, 2, 3, 4, 5, 6, 7, 8})

		in = []int{1}
		shuffle.Shuffle(in)
		So(in, ShouldResemble, []int{1})
	})
}
