package piglatin_test

import (
	"testing"

	"github.com/itsanna/algos/piglatin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPigLatin(t *testing.T) {
	Convey("Verify piglatin", t, func() {
		// General rule: take the first letter of a word, move it to the end, and add "ay".
		str := piglatin.PigLatin("hello")
		So(str, ShouldEqual, "ellohay")

		// A phrase with multiple words should translate each word.
		str = piglatin.PigLatin("hello world")
		So(str, ShouldEqual, "ellohay orldway")

		// A word which begins with a vowel keeps its first letter, and just adds "way" to the end of the word.
		str = piglatin.PigLatin("eat apples")
		So(str, ShouldEqual, "eatway applesway")

		// A word beginning with multiple consonants should move all of them together to the end.
		str = piglatin.PigLatin("drunk strangers")
		So(str, ShouldEqual, "unkdray angersstray")

		// A word which is capitalized should remain capitalized after translation.
		str = piglatin.PigLatin("Hello World")
		So(str, ShouldEqual, "Ellohay Orldway")

		/*
		   // todo
		   // A phrase with punctuation should maintain the position of the punctuation.
		   // punctuationByte := []byte{",", "!", "?", "."}
		   str := piglatin.PigLatin("Hello, world!")
		   So(str, ShouldEqual, "Ellohay, orldway!")

		   // The letters "qu" should stay together when moved to the end of a word.
		   str = piglatin.PigLatin("quickly and quietly")
		   So(str, ShouldEqual, "icklyquay andway ietlyquay")
		*/
	})
}
