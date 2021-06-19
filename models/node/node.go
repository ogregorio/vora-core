package models

import "time"

type Node struct {
	meta    Metadata
	data    string
	visited bool
}

type Metadata struct {
	_id          string
	pHash        string
	nHash        string
	filetype     string
	lastModified time.Time
	createDate   time.Time
}

func New(name string) Node {
	var node Node
	node.data = name
	return node
}

func GetData(node Node) string {
	return node.data
}

func SetData(node Node, ref string) Node {
	node.data = ref
	return node
}
