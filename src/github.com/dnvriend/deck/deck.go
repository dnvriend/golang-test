package deck

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

type deck []string

func NewDeck() deck {
	cards := deck{}
	cardSuits := []string { "Spades", "Diamonds", "Hearts", "Clubs" }
	cardValues := []string { "Ace", "Two", "Three", "Four" }

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) Deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) ToString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}

func NewDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	xs := strings.Split(string(bs), ",")
	return deck(xs)
}

func (d deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)
	for i := range d {
		np := rnd.Intn(len(d) - 1)
		// swap elements
		d[i], d[np] = d[np], d[i]
	}
}
