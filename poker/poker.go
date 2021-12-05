package poker

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`^([2-9]{1}|10|[JQKA])([♢♤♡♧]{1})$`)

var ranks = map[string]int{
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
	"A":  14,
}

var suits = map[string]int{
	"♢": 1,
	"♤": 2,
	"♡": 3,
	"♧": 4,
}

const (
	highcard int = iota
	onepair
	twopair
	three
	straight
	flush
	full
	four
	straightflush
)

func (h *Hand) isFull() (bool, int, int) {
	cards := map[int]int{}

	for _, c := range h.cards {
		if _, ok := cards[c.rank]; !ok {
			cards[c.rank] = 0
		}
		cards[c.rank]++
	}
	if len(cards) != 2 {
		return false, 0, 0
	}
	for _, v := range cards {
		if v != 2 && v != 3 {
			return false, 0, 0
		}
	}
	max := 0
	min := 0

	for k, v := range cards {
		if v == 3 {
			max = k
		}
		if v == 2 {
			min = k
		}
	}
	return true, max, min
}

func (h *Hand) isFour() (bool, int, int) {
	cards := map[int]int{}

	for _, c := range h.cards {
		if _, ok := cards[c.rank]; !ok {
			cards[c.rank] = 0
		}
		cards[c.rank]++
	}
	if len(cards) != 2 {
		return false, 0, 0
	}
	for _, v := range cards {
		if v != 1 && v != 4 {
			return false, 0, 0
		}
	}
	min, max := 0, 0
	for k, v := range cards {
		if v == 4 {
			max = k
		}
		if v == 1 {
			min = k
		}
	}
	return true, max, min
}

func (h *Hand) isThree() bool {
	cards := map[int]int{}

	for _, c := range h.cards {
		if _, ok := cards[c.rank]; !ok {
			cards[c.rank] = 0
		}
		cards[c.rank]++
	}
	if len(cards) != 3 {
		return false
	}
	for _, v := range cards {
		if v != 1 && v != 3 {
			return false
		}
	}
	return true
}

func (h *Hand) isTwoPair() bool {
	cards := map[int]int{}

	for _, c := range h.cards {
		if _, ok := cards[c.rank]; !ok {
			cards[c.rank] = 0
		}
		cards[c.rank]++
	}
	if len(cards) != 3 {
		return false
	}
	for _, v := range cards {
		if v != 1 && v != 2 {
			return false
		}
	}
	return true
}

func (h *Hand) isPair() bool {
	cards := map[int]int{}

	for _, c := range h.cards {
		if _, ok := cards[c.rank]; !ok {
			cards[c.rank] = 0
		}
		cards[c.rank]++
	}
	if len(cards) != 4 {
		return false
	}
	for _, v := range cards {
		if v != 1 && v != 2 {
			return false
		}
	}
	return true
}

func (h *Hand) isStraight() bool {
	ranks := make([]int, 0)
	for _, c := range h.cards {
		ranks = append(ranks, c.rank)
	}
	sort.Ints(ranks)

	//test with high Ace
	highAce := true
	for i := 1; i < 5; i++ {
		if ranks[i] != ranks[i-1]+1 {
			highAce = false
		}
	}
	if highAce {
		return true
	}

	lowAces := make([]int, 0)
	for _, v := range ranks {
		if v == 14 {
			v = 1
		}
		lowAces = append(lowAces, v)
	}
	sort.Ints(lowAces)
	for i := 1; i < 5; i++ {
		if lowAces[i] != lowAces[i-1]+1 {
			return false
		}
	}

	for i, c := range h.cards {
		if c.rank == 14 {
			h.cards[i].rank = 1
		}
	}
	return true
}

func (h *Hand) isFlush() bool {
	suit := 0
	for i, c := range h.cards {
		if i == 0 {
			suit = c.suit
		} else {
			if suit != c.suit {
				return false
			}
		}
	}
	return true
}

func (h *Hand) getType() int {
	if h.isStraight() && h.isFlush() {
		return straightflush
	}

	if ok, _, _ := h.isFour(); ok {
		return four
	}

	if ok, _, _ := h.isFull(); ok {
		return full
	}

	if h.isFlush() {
		return flush
	}

	if h.isStraight() {
		return straight
	}

	if h.isThree() {
		return three
	}

	if h.isTwoPair() {
		return twopair
	}

	if h.isPair() {
		return onepair
	}

	return highcard
}

//Card type
type Card struct {
	rank, suit int
	str        string
}

//Hand type
type Hand struct {
	cards []Card
}

func newHand(hand string) (*Hand, error) {
	c := strings.Fields(hand)
	if len(c) != 5 {
		return nil, errors.New("Not enough cards")
	}
	cards := make([]Card, 5)
	for i, card := range c {
		if !re.MatchString(card) {
			return nil, errors.New("Invalid card")
		}
		found := re.FindAllStringSubmatch(card, -1)
		var rank, suit int
		v, ok := ranks[found[0][1]]
		if !ok {
			return nil, errors.New("Invalid rank")
		}
		rank = v
		v, ok = suits[found[0][2]]
		if !ok {
			return nil, errors.New("Invalid suit")
		}
		suit = v

		cards[i] = Card{rank, suit, card}
	}
	return &Hand{cards}, nil
}

func (h *Hand) maxCard() int {
	max := 0
	for _, c := range h.cards {
		if c.rank > max {
			max = c.rank
		}
	}
	return max
}

func (h *Hand) String() string {
	str := make([]string, 0)
	for _, c := range h.cards {
		str = append(str, c.str)
	}
	return strings.Join(str, " ")
}

func (h *Hand) compare(o *Hand) int {

	if ok, vMax, vMin := h.isFour(); ok {
		_, oMax, oMin := o.isFour()
		if vMax > oMax {
			return 1
		}
		if vMax < oMax {
			return -1
		}
		if vMin > oMin {
			return 1
		}
		if vMin < oMin {
			return -1
		}
		return 0
	}

	if ok, vMax, vMin := h.isFull(); ok {
		_, oMax, oMin := o.isFull()
		if vMax > oMax {
			return 1
		}
		if vMax < oMax {
			return -1
		}
		if vMin > oMin {
			return 1
		}
		if vMin < oMin {
			return -1
		}
		return 0
	}

	sortedHand := make([]int, 5)
	sortedOther := make([]int, 5)

	for i := 0; i < 5; i++ {
		sortedHand[i], sortedOther[i] = h.cards[i].rank, o.cards[i].rank
	}
	sort.Ints(sortedHand)
	sort.Ints(sortedOther)

	for i := 4; i >= 0; i-- {
		if sortedHand[i] > sortedOther[i] {
			return 1
		}
		if sortedHand[i] < sortedOther[i] {
			return -1
		}
	}
	return 0

}

//BestHand returns best poker hands passed
func BestHand(hs []string) ([]string, error) {
	hands := map[*Hand]int{}

	for _, h := range hs {
		hand, err := newHand(h)
		if err != nil {
			return nil, errors.New("Invalid")
		}
		hands[hand] = hand.getType()
	}

	if len(hands) == 1 {
		for h := range hands {
			return []string{h.String()}, nil
		}

	}

	maxHand := 0

	for _, v := range hands {
		if v > maxHand {
			maxHand = v
		}
	}

	for k, v := range hands {
		if v < maxHand {
			delete(hands, k)
		}
	}

	maxHands := make([]*Hand, 0)

	for k := range hands {
		if len(maxHands) == 0 {
			maxHands = append(maxHands, k)
			continue
		}

		switch maxHands[0].compare(k) {
		case -1:

			maxHands = []*Hand{k}
		case 0:

			maxHands = append(maxHands, k)
		}
	}

	strs := make([]string, 0)

	for _, h := range hs {
		for _, m := range maxHands {
			if h == m.String() {
				strs = append(strs, m.String())
			}
		}
	}

	return strs, nil
}
