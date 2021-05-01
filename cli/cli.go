package cli

import (
	"fmt"
	consts "graphgdb/cli/commands"
	read "graphgdb/io/read"
)

func Run() {
	consts.DEBUG = true
	var forceExit bool
	for !forceExit {
		fmt.Print("[graphgdb #] ")
		forceExit = Interpreter(read.Read())
	}
}
