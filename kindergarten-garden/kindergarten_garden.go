package kindergarten

import (
	"errors"
	"strings"
)

var grasses = map[rune]string{
	'V': "violets",
	'G': "grass",
	'R': "radishes",
	'C': "clover",
}

var errNotUniqueChild error = errors.New("Not Unique children")
var errInvalidPattern error = errors.New("Invalid pattern")

//Child of the classroom
type Child struct {
	name  string
	seeds []string
	left  *Child
	right *Child
}

func (c *Child) addFriend(childName string) error {
	switch {
	case childName == c.name:
		return errNotUniqueChild
	case childName < c.name:
		if c.left == nil {
			c.left = &Child{name: childName}
		} else {
			if err := c.left.addFriend(childName); err != nil {
				return err
			}
		}
	case childName > c.name:
		if c.right == nil {
			c.right = &Child{name: childName}
		} else {
			if err := c.right.addFriend(childName); err != nil {
				return err
			}
		}
	}
	return nil
}

//Garden is a classroom
type Garden struct {
	rootChild *Child
	children  []*Child
}

func (g *Garden) addChild(childName string) error {
	if g.rootChild == nil {
		g.rootChild = &Child{name: childName}
		return nil
	}

	return g.rootChild.addFriend(childName)
}

func (c *Child) getFriends() []*Child {
	var leftVals, rightVals []*Child
	if c.left == nil {
		leftVals = []*Child{}
	} else {
		leftVals = c.left.getFriends()
	}
	if c.right == nil {
		rightVals = []*Child{}
	} else {
		rightVals = c.right.getFriends()
	}
	leftVals = append(leftVals, c)
	return append(leftVals, rightVals...)
}

func (g *Garden) getChildren() {
	g.children = g.rootChild.getFriends()
}

//Plants gives for a child the Plants he plants
func (g *Garden) Plants(child string) ([]string, bool) {
	for _, c := range g.children {
		if c.name == child {
			return c.seeds, true
		}
	}
	return nil, false
}

//NewGarden generates Garden from patterne and children list
func NewGarden(pattern string, children []string) (*Garden, error) {
	g := Garden{}
	for _, c := range children {
		if err := g.addChild(c); err != nil {
			return nil, errNotUniqueChild
		}
	}
	g.getChildren()
	rows := strings.Split(pattern, "\n")
	if len(rows) != 3 {
		return nil, errInvalidPattern
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) != len(children)*2 {
			return nil, errInvalidPattern
		}
		for j, r := range row {
			if _, ok := grasses[r]; !ok {
				return nil, errInvalidPattern
			}
			g.children[j/2].seeds = append(g.children[j/2].seeds, grasses[r])
		}
	}
	return &g, nil
}
