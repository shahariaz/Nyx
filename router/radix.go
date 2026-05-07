package router

type Router struct {
	root *Node
}

type Node struct {
	path     string
	children []*Node
}
