package tree

// Record is a struct used to declare what Nodes are created
// and also their relationships to other Nodes
type Record struct {
	ID, Parent int
}

// Node is a struct used to create a tree data structure
type Node struct {
	ID       int
	Children []*Node
}

func delete(slice []Record, index int) []Record {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// Build takes in a Record array to build a tree of Nodes
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	root := &Node{0, []*Node{}}
	for i, r := range records {
		if r.Parent == r.ID {
			root = &Node{r.ID, []*Node{}}
			records = delete(records, i)
		}
	}
	current := root
	for len(records) > 0 {
		for i2, r2 := range records {
			if r2.Parent == current.ID {
				current.Children = append(current.Children, &Node{r2.ID, []*Node{}})
				records = delete(records, i2)
			}
		}
	}
	return root, nil
}
