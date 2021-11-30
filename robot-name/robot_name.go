package robotname

import (
	"fmt"
	"math/rand"
)

var robotsList = map[string]int{}

type Robot struct {
	name string
}

func generateName() string {
	for {
		name := ""
		for len(name) < 2 {
			name += fmt.Sprintf("%c", rand.Int31n(26)+65)
		}
		name = name + fmt.Sprintf("%03v", rand.Int31n(1000))
		if robotsList[name] == 0 {
			robotsList[name] = 1
			return name
		}
	}
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = generateName()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = generateName()
}
