package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck which is a slice of strings

func newDeck() deck {
	cards := deck{}

	// all 52 cards. let's loop it
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		// Option 1 - either log the error and return the call
		// Option 2, if we fail to load the deck, we will return a new Dec
		// Option 3 - log the error and entirely quit the program
		fmt.Println("Error ", err)
		os.Exit(1)
	}
	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {
	// rand.Intn would produce  a randome number, but it will be same everytime cuz of the same seed
	// thus we create a new source and seed all together to have nice random. Also, time.now().unixnano is 
	// just to use to get some seed of type int64 every time it will be different
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

type deck []string
