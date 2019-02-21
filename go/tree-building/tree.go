package tree

// Record is a struct used to declare what Nodes are created
// and also their relationships to other Nodes
type Record struct {
	ID     int
	Parent int
}

// Node is a struct used to create a tree data structure
type Node struct {
	ID       int
	Children []*Node
}

// Build takes in a Record array to build a tree of Nodes
func Build(records []Record) (*Node, error) {
	n := &Node{0, []*Node{}}
	return n, nil
}
