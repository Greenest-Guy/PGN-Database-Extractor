// O(n)

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"PGN-Database-Extractor/config"

	"github.com/corentings/chess/v2"
	"github.com/schollz/progressbar/v3"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil)) // http://localhost:6060/debug/pprof/
	}()

	count := 0
	bar := progressbar.Default(100)
	lastPercent := 0
	const num_games = 1000000
	var games []string

	start := time.Now() // Start counting time

	f, err := os.Open(config.PgnPath())
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := chess.NewScanner(f)

	for scanner.HasNext() {
		game, _ := scanner.ScanGame()

		games = append(games, game.Raw)

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
