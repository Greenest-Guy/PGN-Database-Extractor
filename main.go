package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/corentings/chess/v2"
)

func main() {
	start := time.Now()

	f, err := os.Open("first_million.pgn")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := chess.NewScanner(f)

	count := 0
	num_games := 10000

	for scanner.HasNext() {
		game, _ := scanner.ParseNext()
		game.GetTagPair("Opening")

		count++
		if count == num_games {
			break
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("Number of games processed: %d\n", num_games)
	fmt.Printf("Time taken: %s\n", elapsed)
}
