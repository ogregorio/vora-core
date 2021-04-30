package io

import (
	"bufio"
	"os"
	"strings"
)

//Read a data and replace first \n with " "
func Read() string {
	phrase, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	strings.Replace(phrase, "\n", " ", 1)
	return phrase
}

//AppendLine concat string vector with another string vector
func AppendLine(vector []string, line string, regex string) []string {
	words := strings.Split(line, regex)
	for i := 0; i < len(words); i++ {
		vector = append(vector, words[i])
	}
	return vector
}
