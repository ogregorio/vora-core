package models

type Node struct {
	ref string
}

func New(name string) Node {
	var node Node
	node.ref = name
	return node
}

func GetRef(node Node) string {
	return node.ref
}

func SetRef(node Node, ref string) Node {
	node.ref = ref
	return node
}
