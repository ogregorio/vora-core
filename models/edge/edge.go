package edge

type Edge struct {
	nodeA    string
	nodeB    string
	relation string
}

func New(nA string, nB string) Edge {
	var edge Edge
	edge.nodeA = nA
	edge.nodeB = nB
	edge.relation = "@"
	return edge
}

func GetRef(e Edge) string {
	return e.nodeA + e.relation + e.nodeB
}
