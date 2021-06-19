package run

import "vora/cmd/cmdstructs"

var add cmdstructs.Add

func Add(params cmdstructs.Add) bool {
	add = params
	switch add.Model {
	case "graph":
		//addGraph(add.Domain, add.Name)
	}
	return true
}

// func addGraph(d string, g string) int {

// }

// func addNode(d string, g string, n string) int {
// 	return 0
// }

// func addEdge(d string, g string, n string, n2 string, r string) int {
// 	return 0
// }
