package piglatin

import "strings"

/*
Pig Latin
=========

Write a function in any language of your choice that translates a phrase to
Pig Latin using the following rules:


1. General rule: take the first letter of a word, move it to the end, and add "ay".
   Example: "hello" becomes "ellohay".

2. A phrase with multiple words should translate each word.
   Example: "hello world" becomes "ellohay orldway"

3. A word which begins with a vowel keeps its first letter, and just adds "way"
   to the end of the word.
   Example: "eat apples" becomes "eatway applesway"

4. A word which is capitalized should remain capitalized after translation.
   Example: "Hello World" becomes "Ellohay Orldway"

5. A phrase with punctuation should maintain the position of the punctuation.
   Example: "Hello, world!" becomes "Ellohay, orldway!"

6. A word beginning with multiple consonants should move all of them together to the end.
   Example: "drunk strangers" becomes "unkdray angersstray"

7. The letters "qu" should stay together when moved to the end of a word.
   Example: "quickly and quietly" becomes "icklyquay andway ietlyquay"
*/

type Vowels []string

func (v Vowels) Contains(letter string) bool {
	for _, char := range v {
		if char == letter {
			return true
		}
	}
	return false
}

func PigLatin(word string) string {
	// separate multiple words by space
	words := strings.Split(word, " ")

	// for each word, run the logic below
	var strArray []string
	for _, w := range words {
		s := Rearrange(w)
		strArray = append(strArray, s)
	}

	// combine
	result := strings.Join(strArray, " ")
	return result
}

func Rearrange(word string) string {
	var result string
	letters := []byte(word)
	firstCapitalized := false

	// if the first letter is capitalized, set flag
	if strings.ToTitle(string(letters[0])) == string(letters[0]) {
		firstCapitalized = true
	}

	vowels := Vowels{"a", "e", "i", "o", "u"}
	// if first letter is not a vowel
	if !vowels.Contains(string(letters[0])) {
		// check if next letters are consonants
		// move it to the end
		cons := []byte{
			letters[0],
		}

		for i := 1; i < len(letters)-1; i++ {
			// if consonant, add to consonants
			if !vowels.Contains(string(letters[i])) {
				cons = append(cons, letters[i])
			} else {
				break
			}
		}

		for i := 0; i < len(letters)-len(cons); i++ {
			letters[i] = letters[i+len(cons)]
		}
		// copy cons into last part
		copy(letters[len(letters)-len(cons):], cons)

		// append byte value of "ay"
		str := []byte("ay")
		letters = append(letters, str...)
	} else {
		// add "way" to the end
		str := []byte("way")
		letters = append(letters, str...)
	}
	result = string(letters)
	if firstCapitalized {
		// keep capitalization in the first letter
		lowered := strings.ToLower(result)
		titled := strings.Title(lowered)
		result = titled
		return result
	}
	return result
}
