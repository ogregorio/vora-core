package cli

import (
	"fmt"
	commands "graphgdb/cli/commands"
	consts "graphgdb/cli/commands"
)

func Run() {
	consts.DEBUG = true
	fmt.Println("[graphgdb #]")
	commands.Test("teste")
}
