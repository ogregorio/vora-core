package models

import graph "graphgdb/models/graph"

var graphs []graph.Graph

func NewDatabase() {
	graphs = []graph.Graph{}
}

func AddNode(g string, n string) {
	gr := GetGraphByName(g)
	graph.AddNode(gr, n)
}

func AddGraph(name string) {
	addGraph(graph.New(name))
}

func addGraph(g graph.Graph) {
	graphs = append(graphs, g)
}

func TotalGraphs() int {
	return len(graphs)
}

func GetGraphByName(name string) graph.Graph {
	for i := 0; i < len(graphs); i++ {
		if graph.GetName(graphs[i]) == name {
			return graphs[i]
		}
	}
	panic("graph not found")
}
