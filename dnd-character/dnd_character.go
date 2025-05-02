package dndcharacter

import (
	"math"
	"math/rand"
	"time"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rand.Seed(time.Now().UnixNano())
	dices := []int{}
	for i := 0; i < 4; i++ {
		dices = append(dices, rand.Intn(6)+1)
	}
	minValue := 0
	minIndice := 0
	for i, v := range dices {
		if minValue == 0 || v <= minValue {
			minValue = v
			minIndice = i
		}
	}

	sum := 0
	for i, v := range dices {
		if i != minIndice {
			sum += v
		}
	}

	return sum
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	c.Hitpoints = 10 + Modifier(c.Constitution)
	return c
}
