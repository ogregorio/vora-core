package main

import (
	cmd "graphgdb/cmd"
	system "graphgdb/system"
)

func main() {
	system.System()
	cmd.Execute()
}
