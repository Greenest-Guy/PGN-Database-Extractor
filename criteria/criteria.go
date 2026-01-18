package criteria

import "fmt"

/*
TimeControl - UltraBullet, Bullet, Blitz, Rapid, Classical

SkillGroup - Beginner, Intermediate, Expert

MaxEloDiff - The Maximum allowed elo difference between black and white

FileName - Name of output file (csv)

NumGames - Number of games in the PGN to be parsed
*/

var (
	TimeControl = "Rapid"
	SkillGroup  = "Intermediate"
	MaxEloDiff  = 150
	FileName    = fmt.Sprintf("%s_%s_%d.csv", TimeControl, SkillGroup, MaxEloDiff)
	NumGames    = 91549148 // Lichess monthly dump - 2025 October 29.9 GB 91,549,148
)
