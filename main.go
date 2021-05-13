package main

import (
	cmd "vora/cmd"
	database "vora/models"
	system "vora/system"
)

func main() {
	system.System()
	database.NewDomain()
	cmd.Execute()
}
