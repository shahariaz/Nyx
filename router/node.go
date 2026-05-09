package router

type Node struct {
	path string

	handler HandlerFunc

	children []*Node

	isParam    bool
	paramName  string
	isWildCard bool
	indices    string
}

func NewNode(path string) *Node {
	return &Node{
		path:     path,
		children: make([]*Node, 0, 2),
		indices:  "",
	}
}

func (n *Node) addChild(child *Node) {
	insertIndex := len(n.children)
	for i, c := range n.children {
		if !child.isParam && c.isParam {
			insertIndex = i
			break
		}
		if child.isParam && c.isWildCard {
			insertIndex = i
			break
		}
	}
	if insertIndex == len(n.children) {
		n.children = append(n.children, child)
	} else {
		n.children = append(n.children[:insertIndex], append([]*Node{child}, n.children[insertIndex:]...)...)
	}
}
