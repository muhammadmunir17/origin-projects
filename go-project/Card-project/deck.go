package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"spades", "hearts", "diamonds", "clubs"}
	cardValues := []string{"ace", "two", "three", "four"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}
	

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) savetofile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func ReadFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {

		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		newposition := r.Intn(len(d) - 1)

		d[i], d[newposition] = d[newposition], d[i]
	}
}
