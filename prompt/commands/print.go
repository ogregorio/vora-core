package commands

import (
	"vora-core/gui"
	database "vora-core/models"
	"vora-core/models/graph"
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
	addGraph("teste")
	addNode("teste", "one")
	addNode("teste", "two")
	addNode("teste", "three")
	addEdge("teste", "one", "two")
	addEdge("teste", "two", "three")
	graph.Print(database.GetGraphByName(name))
}
