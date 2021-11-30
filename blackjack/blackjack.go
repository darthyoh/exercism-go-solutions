package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace":
    		return 11
        case "king", "queen", "jack", "ten":
    		return 10
        case "nine":
    		return 9
        case "eight":
    		return 8
        case "seven":
    		return 7
        case "six":
    		return 6
        case "five":
    		return 5
        case "four":
    		return 4
        case "three":
    		return 3
        case "two":
    		return 2
        default:
    		return 0
    }
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2)==21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
    if !isBlackjack {
        return "P"
    }
	if isBlackjack && dealerScore<10 {
        return "W"
    }
	return "S"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
    if handScore >= 17 {
        return "S"
    } else if handScore <= 11 {
    	return "H"
    }
	if dealerScore >= 7 {
        return "H"
    }
	return "S"
}
