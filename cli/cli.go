package cli

import (
	fmt "fmt"
	access "graphgdb/access"
	read "graphgdb/io/read"
)

func Run() {
	var forceExit bool

	access.Login("arthur", "asdgiasghdasnda")

	for !forceExit {
		fmt.Print("[graphgdb #] ")
		forceExit = Interpreter(read.Read())
	}
}
