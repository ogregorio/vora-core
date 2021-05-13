package graph

import (
	json "encoding/json"
	fmt "fmt"
	os "os"
	edge "vora-core/models/edge"
	node "vora-core/models/node"

	tableprinter "github.com/lensesio/tableprinter"
)

const X byte = 0
const Y byte = 1

type Graph struct {
	name  string
	space map[string][]edge.Edge
	nodes map[string]node.Node
	edges map[string]edge.Edge
}

func New(s string) Graph {
	var graph Graph
	graph.name = s
	graph.space = make(map[string][]edge.Edge)
	graph.nodes = make(map[string]node.Node)
	graph.edges = make(map[string]edge.Edge)
	return graph
}

func AddNode(g Graph, s string) {
	n := node.New(s)
	g.nodes[s] = n
	g.space[node.GetRef(n)] = []edge.Edge{}
}

func AddEdge(g Graph, s1 string, s2 string) {
	var n [2]string
	n[0] = node.GetRef(g.nodes[s1])
	n[1] = node.GetRef(g.nodes[s2])
	e := edge.New(n[0], n[1])
	g.edges[edge.GetRef(e)] = e
	g.space[n[0]] = append(g.space[n[0]], e)
	g.space[n[1]] = append(g.space[n[0]], e)
}

func TotalNodes(g Graph) int {
	return len(g.nodes)
}

func AllNodes(g Graph) string {
	var allNodes string
	for k, _ := range g.nodes {
		allNodes += node.GetRef(g.nodes[k])
	}
	return allNodes
}

func GetName(g Graph) string {
	return g.name
}

func Print(g Graph) {
	json, _ := json.Marshal(g.space)
	fmt.Println(string(json))
}

func Printf(g Graph) {
	json, _ := json.Marshal(g.space)
	tableprinter.PrintJSON(os.Stdout, json)
}
