package sublist

import (
	"reflect"
)

//Relation between two lists
type Relation string

//define the Relations between lists
const (
	EQUAL     = Relation("equal")
	SUBLIST   = Relation("sublist")
	SUPERLIST = Relation("superlist")
	UNEQUAL   = Relation("unequal")
)

//Sublist return relation between two lists
func Sublist(first, second []int) Relation {
	if len(first) == 0 {
		if len(second) == 0 {
			return EQUAL
		}
		return SUBLIST
	}
	if len(second) < len(first) {
		ret := Sublist(second, first)
		if ret == SUBLIST {
			return SUPERLIST
		}
		return ret
	}
	for i := 0; i <= len(second)-len(first); i++ {
		if reflect.DeepEqual(first[:], second[i:i+len(first)]) {
			if len(first) == len(second) {
				return EQUAL
			}
			return SUBLIST
		}
	}
	return UNEQUAL
}
