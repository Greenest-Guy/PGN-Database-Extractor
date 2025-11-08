// O(n)

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

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil)) // http://localhost:6060/debug/pprof/profile?seconds=5
	}()

	count := 0
	bar := progressbar.Default(100)
	lastPercent := 0
	const num_games = 91549148

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

	// Scan each game (NO MOVE VALIDATION)
	scanner := chess.NewScanner(f)

	for scanner.HasNext() {
		game, _ := scanner.ScanGame()
		rawpgn := game.Raw

		if meetsCriteria(rawpgn) {
			chessGame := csvwriter.ChessGame{
				WhiteElo:    parseInt32(getTag(rawpgn, "WhiteElo")),
				BlackElo:    parseInt32(getTag(rawpgn, "BlackElo")),
				TimeControl: getTag(rawpgn, "TimeControl"),
				Moves:       getMoves(rawpgn),
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

	elapsed := time.Since(start)

	// Displays info
	fmt.Printf("\n\n")
	fmt.Printf("Number of games processed: %d\n", count)
	fmt.Printf("Number of games extracted: %d\n", games_written)
	fmt.Printf("Time elapsed: %s\n", elapsed)
}

func getTag(rawPGN string, tag string) string {
	scanner := bufio.NewScanner(strings.NewReader(rawPGN))
	re := regexp.MustCompile(`"(.*?)"`)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "["+tag) { // checks if tag exists
			if match := re.FindStringSubmatch(line); len(match) > 1 { // checks if tag has value
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
	return "" // returns empty string if moves not found
}

func meetsCriteria(rawPGN string) bool {
	timecontrol, err := timecontrols.GetTimeControl(getTag(rawPGN, "TimeControl"))
	if err != nil {
		fmt.Println(err)
		return false
	}

	skillgroup, err := skillgroups.GetSkillGroup(getTag(rawPGN, "BlackElo"), getTag(rawPGN, "WhiteElo"))
	if err != nil {
		fmt.Println(err)
		return false
	}

	elodiff, err := elodiffs.GetEloDiff(getTag(rawPGN, "BlackElo"), getTag(rawPGN, "WhiteElo"))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return timecontrol == criteria.TimeControl && skillgroup == criteria.SkillGroup && elodiff <= criteria.MaxEloDiff && getMoves(rawPGN) != ""
}

func parseInt32(s string) int32 {
	var result int32
	fmt.Sscanf(s, "%d", &result)
	return result
}
