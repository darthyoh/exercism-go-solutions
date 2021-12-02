package pov

import "fmt"

//Arc type symbolize link between nodes
type Arc struct{ fr, to string }

//Graph struct
type Graph struct {
	nodes []string
	arcs  []Arc
}

//New generates a new pointer to a Graph
func New() *Graph {
	return new(Graph)
}

//AddNode adds a node to a Graph
func (g *Graph) AddNode(label string) {
	g.nodes = append(g.nodes, label)
}

//AddArc adds a relation between nodes
func (g *Graph) AddArc(fr, to string) {
	g.arcs = append(g.arcs, Arc{fr, to})
}

//ArcList lists all arcs in a Graph
func (g *Graph) ArcList() []string {
	list := make([]string, 0)
	for _, a := range g.arcs {
		list = append(list, fmt.Sprintf("%v -> %v", a.fr, a.to))
	}
	return list
}

//ChangeRoot generates a new Graph with a new rootNode
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	newGraph := &Graph{nodes: g.nodes}
	for _, a := range g.getChildArcs(newRoot) {
		newGraph.AddArc(a.fr, a.to)
	}
	for _, a := range g.getParentArcs(newRoot) {
		newGraph.AddArc(a.fr, a.to)
	}
	return newGraph
}

func (g *Graph) getChildArcs(node string, prevent ...string) []Arc {
	foundArcs := make([]Arc, 0)
	for _, a := range g.arcs {
		prev := true
		if len(prevent) == 1 {
			if prevent[0] == a.to {
				prev = false
			}
		}
		if a.fr == node && prev {
			foundArcs = append(foundArcs, a)
			foundArcs = append(foundArcs, g.getChildArcs(a.to)...)
		}
	}
	return foundArcs
}

func (g *Graph) getParentArcs(node string) []Arc {
	foundArcs := make([]Arc, 0)
	for _, a := range g.arcs {
		if a.to == node {
			foundArcs = append(foundArcs, Arc{fr: a.to, to: a.fr})
			foundArcs = append(foundArcs, g.getChildArcs(a.fr, node)...)
			foundArcs = append(foundArcs, g.getParentArcs(a.fr)...)
		}
	}
	return foundArcs
}
