package main

import "fmt"

func main() {
	gameDeck := newDeck()
	fmt.Println(gameDeck.toString())
	gameDeck.shuffleDeck()
	fmt.Println(gameDeck.toString())
}
