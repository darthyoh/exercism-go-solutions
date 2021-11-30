package zebra

import "fmt"

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

type Property struct {
	category, value string
}

type LinkedProperties struct {
	p1, p2 Property
}

func (l *LinkedProperties) getLinked(category, value string) (Property, bool) {
	if l.p1.category == category && l.p1.value == value {
		return Property{category: l.p2.category, value: l.p2.value}, true
	}
	if l.p2.category == category && l.p2.value == value {
		return Property{category: l.p1.category, value: l.p1.value}, true
	}
	return Property{}, false
}

type House struct {
	//nationality, cigarette, drink, animal, color string
	founds map[string]string
	lefts  map[string][]string
}

func (h House) String() string {
	nationality := h.founds["nationality"]
	if nationality == "" {
		nationality = fmt.Sprintf("%v", h.lefts["nationality"])
	}
	color := h.founds["color"]
	if color == "" {
		color = fmt.Sprintf("%v", h.lefts["color"])
	}
	animal := h.founds["animal"]
	if animal == "" {
		animal = fmt.Sprintf("%v", h.lefts["animal"])
	}
	drink := h.founds["drink"]
	if drink == "" {
		drink = fmt.Sprintf("%v", h.lefts["drink"])
	}
	cigarette := h.founds["cigarette"]
	if cigarette == "" {
		cigarette = fmt.Sprintf("%v", h.lefts["cigarette"])
	}
	return fmt.Sprintf("Nationality:%v\nColor:%v\nAnimal:%v\nDrink:%v\nCigarette:%v\n", nationality, color, animal, drink, cigarette)
}

func (h House) isPresent(category, value string) bool {
	for _, v := range h.lefts[category] {
		if v == value {
			return true
		}
	}
	return false
}

func newHouse() House {
	return House{
		founds: map[string]string{
			"nationality": "",
			"cigarette":   "",
			"animal":      "",
			"drink":       "",
			"color":       "",
		},
		lefts: map[string][]string{
			"nationality": []string{"Japanese", "Norwegian", "Englishman", "Ukrainian", "Spaniard"},
			"cigarette":   []string{"Kools", "Lucky Strike", "Parliaments", "Chesterfields", "Old Gold"},
			"animal":      []string{"Horse", "Fox", "Snails", "Zebra", "Dog"},
			"drink":       []string{"Water", "Milk", "Coffee", "Tea", "Orange Juice"},
			"color":       []string{"Blue", "Red", "Green", "Yellow", "Ivory"},
		}}
}

type Hypothesis struct {
	houses           []House
	linkedProperties []LinkedProperties
}

func (h *Hypothesis) getSolution() Solution {
	zebra := ""
	water := ""

	for _, house := range h.houses {
		if house.founds["animal"] == "Zebra" {
			zebra = house.founds["nationality"]
		}
		if house.founds["drink"] == "Water" {
			water = house.founds["nationality"]
		}
	}

	return Solution{DrinksWater: water, OwnsZebra: zebra}

}

func (h *Hypothesis) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n#################", h.houses[0], h.houses[1], h.houses[2], h.houses[3], h.houses[4])
}

func (h *Hypothesis) getLinked(category, value string) (Property, bool) {
	for _, p := range h.linkedProperties {
		if l, ok := p.getLinked(category, value); ok {
			return Property{category: l.category, value: l.value}, ok
		}
	}
	return Property{}, false
}

func (h *Hypothesis) getIndirectLinked(category, value string) []Property {
	indirect := make([]Property, 0)
	for _, linked := range h.linkedProperties {
		if (linked.p1.category == category && linked.p1.value != value) || (linked.p2.category == category && linked.p2.value != value) {
			indirect = append(indirect, linked.p1)
		}
	}
	return indirect
}

func newHypothesis() *Hypothesis {
	hypo := &Hypothesis{}
	for i := 0; i < 5; i++ {
		hypo.houses = append(hypo.houses, newHouse())
	}
	hypo.linkedProperties = []LinkedProperties{
		LinkedProperties{
			p1: Property{category: "nationality", value: "Englishman"},
			p2: Property{category: "color", value: "Red"},
		},
		LinkedProperties{
			p1: Property{category: "nationality", value: "Spaniard"},
			p2: Property{category: "animal", value: "Dog"},
		},
		LinkedProperties{
			p1: Property{category: "drink", value: "Coffee"},
			p2: Property{category: "color", value: "Green"},
		},
		LinkedProperties{
			p1: Property{category: "nationality", value: "Ukrainian"},
			p2: Property{category: "drink", value: "Tea"},
		},
		LinkedProperties{
			p1: Property{category: "cigarette", value: "Old Gold"},
			p2: Property{category: "animal", value: "Snails"},
		},
		LinkedProperties{
			p1: Property{category: "cigarette", value: "Kools"},
			p2: Property{category: "color", value: "Yellow"},
		},
		LinkedProperties{
			p1: Property{category: "cigarette", value: "Lucky Strike"},
			p2: Property{category: "drink", value: "Orange Juice"},
		},
		LinkedProperties{
			p1: Property{category: "nationality", value: "Japanese"},
			p2: Property{category: "cigarette", value: "Parliaments"},
		},
	}

	return hypo
}

func (h *Hypothesis) set(category string, houseID int, value string) {
	if h.houses[houseID].founds[category] == "" { //prevent infiny loops
		//set "category" for house
		h.houses[houseID].founds[category] = value
		h.houses[houseID].lefts[category] = []string{}

		//removes probabilities for others houses
		for i := range h.houses {
			if i != houseID {
				h.removeProbability(category, i, value)
			}
		}
		//set "linked category" for the same house
		if p, ok := h.getLinked(category, value); ok {
			h.set(p.category, houseID, p.value)
		}
		//removes probabilities for linked properties to same category
		for _, prop := range h.getIndirectLinked(category, value) {
			h.removeProbability(prop.category, houseID, prop.value)
		}
	}
}

func (h *Hypothesis) removeProbability(category string, houseID int, value string) {
	for i, v := range h.houses[houseID].lefts[category] {
		if v == value {
			h.houses[houseID].lefts[category] = append(h.houses[houseID].lefts[category][:i], h.houses[houseID].lefts[category][i+1:]...)
			break
		}
	}

	if p, ok := h.getLinked(category, value); ok {
		if h.houses[houseID].isPresent(p.category, p.value) {
			h.removeProbability(p.category, houseID, p.value)
		}
	}

	if len(h.houses[houseID].lefts[category]) == 1 {
		h.set(category, houseID, h.houses[houseID].lefts[category][0])
	}
}

//SolvePuzzle resolves the Zebra Puzzle
func SolvePuzzle() Solution {
	hypo := newHypothesis()
	//10
	hypo.set("nationality", 0, "Norwegian")

	//15
	hypo.set("color", 1, "Blue")

	//9
	hypo.set("drink", 2, "Milk")

	//6
	hypo.removeProbability("color", 0, "Ivory")
	hypo.removeProbability("color", 1, "Ivory")
	hypo.removeProbability("color", 4, "Ivory")
	hypo.removeProbability("color", 0, "Green")
	hypo.removeProbability("color", 1, "Green")
	hypo.removeProbability("color", 3, "Red")

	//Kools only remains on house 0
	hypo.set("cigarette", 0, "Kools")

	//12
	hypo.set("animal", 1, "Horse")

	//In house 1, if the Japanese lives here, it smokes Parliaments, so no Lucky Strike, so no Orange Juice, leaving only Tea...
	//But Tea is drunk by Ukrainian... So House 1 is Ukrainian !
	hypo.set("nationality", 1, "Ukrainian")

	//11
	hypo.removeProbability("animal", 3, "Fox")
	hypo.removeProbability("animal", 4, "Fox")

	//Englishman can't live in House 4 : if it did, the house would be red, and he wouldn't drink coffee, so drink Orange Juice and smoke Lucky Strike
	//But this is incompatible because Spaniard would live in house 2 or 3, having the Dog (so not the Snails) so don't smoke Old Gold, so smoke Parliaments
	//And this is incompatible because this is the Japanese who smokes Parliaments
	hypo.removeProbability("nationality", 4, "Englishman")

	//fmt.Println(hypo)
	return hypo.getSolution()
}
