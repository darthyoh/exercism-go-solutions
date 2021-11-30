package tree

import (
	"fmt"
	"sort"
)

//Record struct
type Record struct {
	ID     int
	Parent int
}

//Node struct
type Node struct {
	ID       int
	Children []*Node
}

//Build function
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodeMap := make(map[int]int)

	for i, v := range records {
		if (v.ID == 0 && v.Parent != 0) || (v.ID < v.Parent) || (v.ID == v.Parent && v.ID != 0) || (i > 0 && v.ID != records[i-1].ID+1) {
			return nil, fmt.Errorf("Error")
		}
		nodeMap[v.ID]++
	}

	if root := nodeMap[0]; root != 1 {
		return nil, fmt.Errorf("Error")
	}

	for _, v := range nodeMap {
		if v > 1 {
			return nil, fmt.Errorf("Error")
		}
	}

	if rootChilds, err := getChildren(0, &records); err == nil {
		return &Node{ID: 0, Children: rootChilds}, nil
	}
	return nil, fmt.Errorf("Error")

}

func getChildren(id int, records *[]Record) ([]*Node, error) {

	children := make([]*Node, 0)

	for _, value := range *records {
		if value.Parent == id && value.ID != 0 {
			if childs, err := getChildren(value.ID, records); err == nil {
				children = append(children, &Node{ID: value.ID, Children: childs})
			} else {
				return nil, fmt.Errorf("Error")
			}
		}
	}

	if len(children) == 0 {
		return nil, nil
	}
	return children, nil
}
