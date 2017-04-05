package set_test

import (
	"fmt"
	"testing"

	"github.com/itsanna/algos/datastructures/set"
	. "github.com/smartystreets/goconvey/convey"
)

func TestKeyMap(t *testing.T) {

	Convey("Verify mapping functions work", t, func() {
		key := "hamburger"
		buckets := 10
		index, err := set.KeyToIndex(key, buckets)
		So(err, ShouldBeNil)
		So(index, ShouldBeGreaterThanOrEqualTo, 0)

		hash, err := set.GenerateHash(key)
		So(err, ShouldBeNil)
		So(hash, ShouldBeGreaterThanOrEqualTo, 0)

		fmt.Println("\n index : ", index, "\n hash :", hash)
	})
}
