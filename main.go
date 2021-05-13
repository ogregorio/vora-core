package main

import (
	cmd "vora-core/cmd"
	database "vora-core/models"
	system "vora-core/system"
)

func main() {
	system.System()
	database.NewDatabase()
	cmd.Execute()
}
