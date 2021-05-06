package commands

import (
	"graphgdb/gui"
	database "graphgdb/models"
	"graphgdb/models/graph"
)

func Print() string {
	gui.P.Println("Type what you want to print:")
	objectType := Default()
	gui.P.Println("Insert object name:")
	name := Default()

	switch objectType {
	case "graph":
		printGraph(name)
	}

	return ""
}

func printGraph(name string) {
	graph.Print(database.GetGraphByName(name))
}
