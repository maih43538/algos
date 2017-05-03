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