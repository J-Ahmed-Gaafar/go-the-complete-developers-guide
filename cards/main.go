package main

func main() {
	//var card = string = "Ace of Spades"
	//card := "Ace of Spades"
	//card = "Five of Diamonds"

	//cards := []string{"Ace of Diamonds", newCard()}
	//cards := deck{"Ace of Diamonds", newCard()}

	//cards = append(cards, "Six of Spades")

	cards := newDeck()

	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}
	cards.print()
}

//func newCard() string {
//	return "Five of Diamonds"
//}
