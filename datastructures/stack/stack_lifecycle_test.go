package stack_test

import (
	"testing"

	"github.com/itsanna/algos/datastructures/stack"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStackLifecycle(t *testing.T) {
	Convey("Stack lifecycle", t, func() {
		people := []string{
			"Norman",
			"Jax",
			"Arya",
			"Glen",
		}
		max := 4
		s := stack.New(max)

		for _, person := range people {
			err := s.Push(person)
			So(err, ShouldBeNil)
			So(s.Length(), ShouldBeGreaterThan, 0)
		}
		err := s.Push("Otto")
		So(err, ShouldEqual, stack.ErrMaxCapReached)

		for _, _ = range people {
			val, err := s.Pop()
			So(err, ShouldBeNil)
			So(val, ShouldNotEqual, "")
		}
		So(s.Length(), ShouldEqual, 0)
	})
}
