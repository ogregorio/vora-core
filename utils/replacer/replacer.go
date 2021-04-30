package replacer

import (
	"strings"
)

//Replace replace in phrase where has a regex
func Replace(phrase string, word string, regex string) string {
	return strings.ReplaceAll(phrase, regex, word)
}

//VectorReplace replace in an array phrase where has a regex match
func VectorReplace(phrases []string, word string, regex string) []string {
	var newPhrases []string = phrases
	for i := 0; i < len(phrases); i++ {
		newPhrases[0] = Replace(phrases[0], word, regex)
	}
	return newPhrases
}
