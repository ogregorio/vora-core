package cli

import (
	cli "graphgdb/cli/commands"
	"strings"
)

func Interpreter(input string) bool {
	params := strings.Split(input, " ")
	switch len(params) {
	case 2:
		return decode(params[0], params[1])
	}

	return false
}

func decode(param string, arguments string) bool {
	switch param {
	case "TEST":
		cli.Test(arguments)
		return true
	}
	return false
}
