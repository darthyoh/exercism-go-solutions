package bowling

import "errors"

//Game is a bowling game
type Game struct {
	rollID             int
	activeScore        int
	lastRollBeforeTurn bool
	bonus              []int
	score              int
	lastFrame          []int
	completed          bool
}

//NewGame initialize a bowling game
func NewGame() *Game {
	return &Game{rollID: 1}
}

func (g *Game) pendingSpare() {
	if len(g.bonus) == 0 {
		g.bonus = append(g.bonus, 1)
	} else {
		g.bonus[0]++
	}
}

func (g *Game) pendingStrike() {
	switch len(g.bonus) {
	case 0:
		g.bonus = []int{1, 1}
	case 1:
		g.bonus[0]++
		g.bonus = append(g.bonus, 1)
	default:
		g.bonus[0]++
		g.bonus[1]++
	}
}

func (g *Game) getBonus(p int) {
	if len(g.bonus) != 0 {
		g.score += p * g.bonus[0]
		g.bonus = g.bonus[1:]
	}
}

//Score get the score of a completed bowling Game
func (g *Game) Score() (int, error) {
	if !g.completed {
		return 0, errors.New("Game not over")
	}
	return g.score, nil
}

//Roll determines what to do with pins
func (g *Game) Roll(p int) error {
	if g.completed {
		return errors.New("Game over")
	}
	if p < 0 || p > 10 {
		return errors.New("Invalid roll")
	}

	//add the pending bonus from previous Strikes and Spares to the score
	g.getBonus(p)
	//add the pins to the score
	g.score += p

	// normal case (nineth first frames)
	if g.rollID <= 9 {
		g.activeScore += p

		if g.activeScore > 10 {
			return errors.New("Invalid sum (>10)")
		}
		if g.activeScore == 10 {
			if p == 10 {
				g.pendingStrike()
			} else {
				g.pendingSpare()
			}
			g.lastRollBeforeTurn = false
			g.rollID++
			g.activeScore = 0
			return nil
		}

		if g.lastRollBeforeTurn {
			g.rollID++
			g.activeScore = 0
		}
		g.lastRollBeforeTurn = !g.lastRollBeforeTurn

		return nil
	}

	// lastFrame
	if len(g.lastFrame) == 0 { //first roll simply add the score
		g.lastFrame = append(g.lastFrame, p)
		return nil
	}

	if len(g.lastFrame) == 1 { //second roll : detect spare or end of game
		if g.lastFrame[0] != 10 && g.lastFrame[0]+p > 10 { // normal roll so detect error of sum
			return errors.New("Invalid sum (>10)")
		}
		if g.lastFrame[0] != 10 && g.lastFrame[0]+p != 10 { // normal roll : end game

			g.lastFrame = append(g.lastFrame, p)
			g.completed = true
			return nil
		}
		g.lastFrame = append(g.lastFrame, p)
		return nil
	}

	if len(g.lastFrame) == 2 { //third roll
		if g.lastFrame[0] == 10 && g.lastFrame[1] != 10 && g.lastFrame[1]+p > 10 {
			return errors.New("Invalid sum (>10)")
		}

		g.completed = true
	}

	return nil
}
