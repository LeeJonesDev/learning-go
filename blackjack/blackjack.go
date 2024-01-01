package blackjack

const (
	Stand string = "S"
	Hit   string = "H"
	Split string = "P"
	Win   string = "W"
)

func getCardsMap() map[string]int {
	return map[string]int{
		"other": 0,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
		"ace":   11,
	}
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardsMap := getCardsMap()
	return cardsMap[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardsMap := getCardsMap()

	cardSum := cardsMap[card1] + cardsMap[card2]

	switch {
	case card1 == "ace" && card2 == "ace":
		return Split
	case (cardSum == 21) &&
		(dealerCard == "ace" || cardsMap[dealerCard] == 10):
		return Stand
	case cardSum >= 17 && cardSum <= 20:
		return Stand
	case cardSum >= 12 && cardSum <= 16 && cardsMap[dealerCard] < 7:
		return Stand
	case cardSum >= 12 && cardSum <= 16 && cardsMap[dealerCard] >= 7:
		return Hit
	case cardSum <= 11:
		return Hit
	case (cardSum == 21) &&
		(dealerCard != "ace" || cardsMap[dealerCard] != 10):
		return Win
	default:
		panic("unexpected switch case")
	}
}
