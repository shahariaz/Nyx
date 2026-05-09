package router

type HandlerFunc func()

type Router struct {
	root *Node
}
