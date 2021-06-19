package main

import (
	cmd "vora/cmd"
	database "vora/models/database"
	system "vora/system"
)

func main() {
	var database database.Database
	database.Instantiate()
	system.System()
	cmd.Execute()
}
