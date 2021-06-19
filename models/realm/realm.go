package realm

import (
	json "encoding/json"
	fmt "fmt"
	os "os"
	edge "vora/models/edge"
	node "vora/models/node"

	tableprinter "github.com/lensesio/tableprinter"
)

const X byte = 0
const Y byte = 1

type Realm struct {
	name  string
	space map[string][]edge.Edge
	nodes map[string]node.Node
	edges map[string]edge.Edge
}

func New(s string) Realm {
	var realm Realm
	realm.name = s
	realm.space = make(map[string][]edge.Edge)
	realm.nodes = make(map[string]node.Node)
	realm.edges = make(map[string]edge.Edge)
	return realm
}

func AddNode(r Realm, s string) {
	n := node.New(s)
	r.nodes[s] = n
	r.space[node.GetData(n)] = []edge.Edge{}
}

func AddEdge(r Realm, s1 string, s2 string) {
	var n [2]string
	n[0] = node.GetData(r.nodes[s1])
	n[1] = node.GetData(r.nodes[s2])
	e := edge.New(n[0], n[1])
	r.edges[edge.GetRef(e)] = e
	r.space[n[0]] = append(r.space[n[0]], e)
	r.space[n[1]] = append(r.space[n[0]], e)
}

func TotalNodes(r Realm) int {
	return len(r.nodes)
}

func AllNodes(r Realm) string {
	var allNodes string
	for k, _ := range r.nodes {
		allNodes += node.GetData(r.nodes[k])
	}
	return allNodes
}

func GetName(r Realm) string {
	return r.name
}

func Print(r Realm) {
	json, _ := json.Marshal(r.space)
	fmt.Println(string(json))
}

func Printf(r Realm) {
	json, _ := json.Marshal(r.space)
	tableprinter.PrintJSON(os.Stdout, json)
}
