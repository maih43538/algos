package set_test

import (
	"testing"

	"fmt"

	"github.com/itsanna/algos/datastructures/set"
	. "github.com/smartystreets/goconvey/convey"
)

type Input struct {
	Key   string
	Value string
}

func TestSetLifecycle(t *testing.T) {
	Convey("Verify set insertion works", t, func() {
		size := 1
		ht := set.New(size)
		So(len(ht), ShouldEqual, size)

		in :=
			&Input{
				Key:   "Norman",
				Value: "(352)778-4433",
			}
		err := ht.Insert(in.Key, in.Value)
		So(err, ShouldBeNil)

		bool := ht.ContainsKey(in.Key)
		So(bool, ShouldBeTrue)

		found, err := ht.FindByKey(in.Key)
		So(err, ShouldBeNil)
		So(found, ShouldEqual, in.Value)

		in = &Input{
			Key:   "Norman",
			Value: "(988)460-0090",
		}
		err = ht.Insert(in.Key, in.Value)
		So(err, ShouldEqual, set.ErrKeyAlreadyExists)
	})

	Convey("Verify set handles collisions", t, func() {
		size := 1 // setting up single bucket that will cause collision
		ht := set.New(size)
		So(len(ht), ShouldEqual, size)

		inputs := []*Input{
			&Input{
				Key:   "Jax",
				Value: "(909)990-1254",
			},
			&Input{
				Key:   "Otto",
				Value: "(505)909-7788",
			},
			&Input{
				Key:   "Norman",
				Value: "(352)778-4433",
			},
		}

		for _, in := range inputs {
			err := ht.Insert(in.Key, in.Value)
			So(err, ShouldBeNil)

			bool := ht.ContainsKey(in.Key)
			So(bool, ShouldBeTrue)

			found, err := ht.FindByKey(in.Key)
			So(err, ShouldBeNil)
			So(found, ShouldEqual, in.Value)
		}
	})

	Convey("Verify set deletion works", t, func() {
		size := 2
		ht := set.New(size)

		in := &Input{
			Key:   "Arya",
			Value: "(707)988-1190",
		}
		err := ht.Insert(in.Key, in.Value)
		So(err, ShouldBeNil)

		in = &Input{
			Key:   "Carl",
			Value: "(414)333-1000",
		}
		err = ht.Insert(in.Key, in.Value)
		So(err, ShouldBeNil)

		err = ht.Delete("Arya")
		So(err, ShouldBeNil)

		_, err = ht.FindByKey("Arya")
		So(err, ShouldEqual, set.ErrElementNotFound)
	})

	Convey("Verify set resizing works", t, func() {
		size := 3
		ht := set.New(size)
		in := &Input{
			Key:   "Arya",
			Value: "(707)988-1190",
		}
		err := ht.Insert(in.Key, in.Value)
		So(err, ShouldBeNil)

		err = ht.Resize(6)
		So(err, ShouldBeNil)

		// todo: get resize to work...
		fmt.Println(ht)
	})
}
