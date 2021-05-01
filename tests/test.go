package consts

import (
	cli "graphgdb/cmd/cli"
)

func Test(argv ...string) {
	cli.G.Println("PASS")
}
