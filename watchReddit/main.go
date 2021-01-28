package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type reminderBot struct {
	bot     reddit.Bot
	symbols []string
}

func MatchString(a []string, x string) int {
	idx := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if idx < len(a) && a[idx] == x {
		return idx
	}
	return len(a)
}

func (r *reminderBot) Comment(comment *reddit.Comment) error {

	numSymbols := len(r.symbols)
	commentText := comment.Body

	commentWords := strings.Split(commentText, " ")

	for _, word := range commentWords {
		if len(word) <= 5 {
			if strings.ToUpper(word) == word {
				searchResult := MatchString(r.symbols, word)
				if searchResult < numSymbols {
					fmt.Println(word)
				}
			}
		}
	}

	return nil
}

func main() {

	// Read the ticker symbols into a slice
	fileBytes, err := ioutil.ReadFile("./symbols")
	if err != nil {
		log.Fatalln("Couldn't read ticker symbols")
	}

	sliceData := strings.Split(string(fileBytes), "\n")

	bot, err := reddit.NewBotFromAgentFile("./redvars", 0)

	if err != nil {
		fmt.Println("could not create bot ", err)
		panic(1)
	}

	cfg := graw.Config{SubredditComments: []string{"wallstreetbets"}}
	handler := &reminderBot{bot: bot, symbols: sliceData}

	if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
		fmt.Println("failed to start graw run ", err)
	} else {
		fmt.Println("graw run failed ", wait())
	}

}
