package main

import (
	cmd "graphgdb/cmd"
	database "graphgdb/models"
	system "graphgdb/system"
)

func main() {
	system.System()
	database.NewDatabase()
	cmd.Execute()
}
