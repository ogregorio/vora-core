package commands

import (
	"vora-core/gui"
	database "vora-core/models"
)

func Add() string {
	var graphName, objectType, objectName string
	gui.P.Println("Type what you want to add:")
	objectType = Default()
	gui.P.Println("Type a name to object:")
	objectName = Default()

	if objectType == "node" {
		gui.P.Println("You type 'node', where you add this?")
		graphName = Default()
	}

	switch objectType {
	case "graph":
		addGraph(objectName)
	case "node":
		addNode(graphName, objectName)
	}

	return ""
}

func addGraph(name string) {
	database.AddGraph(name)
}

func addNode(graph string, node string) {
	database.AddNode(graph, node)
}

func addEdge(graph string, node string, node2 string) {
	database.AddEdge(graph, node, node2)
}
