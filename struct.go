package xml2json

import (
	"strings"
)

// Node is a data element on a tree
type Node struct {
	Children              map[string]Nodes
	Data                  string
	ChildrenAlwaysAsArray bool
	Prefix                string
	ChildrenKeys          []string
}

// Nodes is a list of nodes
type Nodes []*Node

// AddChild appends a node to the list of children
func (n *Node) AddChild(s string, c *Node) {
	// Lazy lazy
	if n.Children == nil {
		n.Children = map[string]Nodes{}
	}

	var exists bool
	for _, val := range n.ChildrenKeys {
		if val == s {
			exists = true
		}
	}
	if !exists {
		n.ChildrenKeys = append(n.ChildrenKeys, s)
	}
	n.Children[s] = append(n.Children[s], c)
}

// IsComplex returns whether it is a complex type (has children)
func (n *Node) IsComplex() bool {
	return len(n.Children) > 0
}

func (n *Node) AddNamespacePrefix(prefix string) {
	n.Prefix = prefix
}

// GetChild returns child by path if exists. Path looks like "grandparent.parent.child.grandchild"
func (n *Node) GetChild(path string) *Node {
	result := n
	names := strings.Split(path, ".")
	for _, name := range names {
		children, exists := result.Children[name]
		if !exists {
			return nil
		}
		if len(children) == 0 {
			return nil
		}
		result = children[0]
	}
	return result
}
