package graph

import (
	"fmt"
	edge "graphgdb/models/edge"
	node "graphgdb/models/node"
	strings "strings"
)

const X byte = 0
const Y byte = 1

type Graph struct {
	name  string
	space [][]string
	nodes []node.Node
}

func New(name string) Graph {
	var g Graph
	g.name = name
	g.space = make([][]string, 2)
	g.space[X] = []string{}
	g.space[Y] = []string{}
	return g
}

func AddNode(g Graph, name string) {
	n := node.New(name)
	n = node.SetRef(n, g.name+"@"+node.GetRef(n))
	addNode(g, n)
}

func addNode(g Graph, n node.Node) {
	g.space[0] = append(g.space[X], node.GetRef(n))
	g.space[1] = append(g.space[Y], node.GetRef(n))
	g.nodes = append(g.nodes, n)
}

func AddEdge(e edge.Edge) {

}

func ReturnNodePosition(g Graph, n node.Node) int {
	for i := 0; i < len(g.space[X]); i++ {
		if g.space[X][i] == node.GetRef(n) {
			return i
		}
	}
	return -1
}

func TotalNodes(g Graph) int {
	return len(g.nodes)
}

func AllNodes(g Graph) string {
	var allNodes strings.Builder
	for i := 0; i < len(g.nodes); i++ {
		allNodes.WriteString("node:" + node.GetRef(g.nodes[i]))
	}
	return allNodes.String()
}

func GetName(g Graph) string {
	return g.name
}

func Print(g Graph) {
	fmt.Println("Graph:" + g.name)
	fmt.Printf("\t %s \t", " ")
	for i := 0; i < len(g.space[0]); i++ {
		fmt.Printf("\t %s \t", g.space[X][i])
	}
	fmt.Println()
	for i := 0; i < len(g.space[X]); i++ {
		fmt.Printf("\t %s \t", g.space[i])
		for j := 0; j < len(g.space[Y]); j++ {
			fmt.Printf("\t %s \t", g.space[i][j])
		}
		fmt.Println()
	}
}
