/*
How To Use

	Create a .env file setting variables for both PGN_PATH and OUTPUT_PATH
	Edit the criteria/criteria.go file with the games you'd like to extract from the pgn

Time Complexity

	linear scaling based off of length of pgn file
	O(n)

CPU Usage

	http://localhost:6060/debug/pprof/profile?seconds={num}
	go tool pprof profile
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"regexp"
	"strings"
	"time"

	"PGN-Database-Extractor/config"
	"PGN-Database-Extractor/criteria"
	"PGN-Database-Extractor/csvwriter"
	"PGN-Database-Extractor/elodiffs"
	"PGN-Database-Extractor/skillgroups"
	"PGN-Database-Extractor/timecontrols"

	"github.com/corentings/chess/v2"
	"github.com/schollz/progressbar/v3"
)

var tagValueRegex = regexp.MustCompile(`"(.*?)"`)

func main() {
	count := 0
	bar := progressbar.Default(100)
	lastPercent := 0
	num_games := criteria.NumGames

	start := time.Now() // Start counting time

	f, err := os.Open(config.PgnPath())
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	writer, err := csvwriter.New(config.OutputPath() + "/" + criteria.FileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := chess.NewScanner(f) // Scanner for games (NO MOVE VALIDATION)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for scanner.HasNext() { // Scans num_games in pgn
		game, err := scanner.ScanGame()
		if err != nil {
			log.Printf("Error scanning game at count %d: %v", count, err)
			continue
		}

		rawpgn := game.Raw

		whiteElo := getTag(rawpgn, "WhiteElo")
		blackElo := getTag(rawpgn, "BlackElo")
		timeControl := getTag(rawpgn, "TimeControl")
		moves := getMoves(rawpgn)

		if meetsCriteria(blackElo, whiteElo, timeControl, moves) {
			chessGame := csvwriter.ChessGame{
				WhiteElo:    parseInt32(whiteElo),
				BlackElo:    parseInt32(blackElo),
				TimeControl: timeControl,
				Moves:       moves,
			}
			writer.Append(chessGame)
		}

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

	games_written := writer.GetRowCount()
	writer.Close()

	// Displays info
	fmt.Printf("\n\n")
	fmt.Printf("Number of games processed: %d\n", count)
	fmt.Printf("Number of games extracted: %d\n", games_written)
	fmt.Printf("Time elapsed: %s\n", time.Since(start))
}

func getTag(rawPGN string, tag string) string {
	scanner := bufio.NewScanner(strings.NewReader(rawPGN))

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "["+tag) { // checks if tag exists
			if match := tagValueRegex.FindStringSubmatch(line); len(match) > 1 { // checks if tag has value
				return match[1]
			}
		}
	}
	return "" // returns empty string if tag not found
}

func getMoves(rawPGN string) string {
	scanner := bufio.NewScanner(strings.NewReader(rawPGN))

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "1. ") {
			return line
		}
	}
	return "" // returns empty string if moves not found or no moves
}

func meetsCriteria(blackElo string, whiteElo string, timeControl string, moves string) bool {
	skillgroup, err := skillgroups.GetSkillGroup(blackElo, whiteElo)
	if err != nil {
		return false
	}

	elodiff, err := elodiffs.GetEloDiff(blackElo, whiteElo)
	if err != nil {
		return false
	}

	timecontrol, err := timecontrols.GetTimeControl(timeControl)
	if err != nil {
		return false
	}

	return timecontrol == criteria.TimeControl && skillgroup == criteria.SkillGroup && elodiff <= criteria.MaxEloDiff && moves != ""
}

func parseInt32(s string) int32 {
	var result int32
	fmt.Sscanf(s, "%d", &result)
	return result
}
