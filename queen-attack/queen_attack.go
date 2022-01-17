package queenattack

import (
	"errors"
	"regexp"
)

type position struct {
	initial string
	letter  uint8
	number  uint8
}

func newPosition(initial string) (position, error) {
	re := regexp.MustCompile(`^[a-h]{1}[1-8]{1}$`)
	if !re.MatchString(initial) {
		return position{}, errors.New("Off Board")
	}
	return position{initial: initial, letter: initial[0], number: initial[1]}, nil
}

func (p *position) reset() {
	p.letter = p.initial[0]
	p.number = p.initial[1]
}

func (p *position) decrease() {
	if p.letter > 97 && p.number > 49 {
		p.letter--
		p.number--
		p.decrease()
	}
}

func (p *position) increase() {
	if p.letter > 97 && p.number < 56 {
		p.letter--
		p.number++
		p.increase()
	}
}

func (p position) equals(p1 position) bool {
	return p.letter == p1.letter && p.number == p1.number
}

//CanQueenAttack function
func CanQueenAttack(black, white string) (bool, error) {

	b, errB := newPosition(black)
	w, errW := newPosition(white)

	if errB != nil || errW != nil {
		return false, errors.New("off board")
	}

	if b == w {
		return false, errors.New("same square")
	}

	if b.letter == w.letter || b.number == w.number {
		return true, nil
	}

	b.decrease()
	w.decrease()

	if b.equals(w) {
		return true, nil
	}

	b.reset()
	w.reset()

	b.increase()
	w.increase()

	if b.equals(w) {
		return true, nil
	}

	return false, nil
}
