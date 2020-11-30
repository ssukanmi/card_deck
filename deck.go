package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) printDeck() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() (cardDeck deck) {
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}

	for _, cardSuite := range cardSuites {
		for _, cardValue := range cardValues {
			cardDeck = append(cardDeck, cardValue+" of "+cardSuite)
		}
	}
	return
}

func dealHand(cardDeck deck, handSize int) (deck, deck) {
	return cardDeck[:handSize], cardDeck[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) toByte() []byte {
	return []byte(d.toString())
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, d.toByte(), 0666)
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(byteSlice), ","))
}

func (d deck) shuffleDeck() {
	rand.Seed(time.Now().UnixNano())
	randIndex := len(d) - 1
	for i := range d {
		newPosition := rand.Intn(randIndex)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
