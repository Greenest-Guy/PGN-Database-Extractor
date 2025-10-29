// O(n)

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"PGN-Database-Extractor/config"

	"github.com/corentings/chess/v2"
	"github.com/schollz/progressbar/v3"
)

func main() {
	count := 0
	bar := progressbar.Default(100)
	lastPercent := 0
	const num_games = 10000
	var TimeControls []string

	start := time.Now() // Start counting time

	f, err := os.Open(config.PgnPath())
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := chess.NewScanner(f)

	for scanner.HasNext() {
		game, _ := scanner.ParseNext()
		opening := game.GetTagPair("TimeControl")

		_ = append(TimeControls, opening) // TimeControls

		count++
		if count == num_games {
			bar.Set(100)
			break
		}

		percent := (count * 100) / num_games
		if percent != lastPercent {
			bar.Set(percent)
			lastPercent = percent
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("\n\n")
	fmt.Printf("Number of games processed: %d\n", num_games)
	fmt.Printf("Time taken: %s\n", elapsed)
}
