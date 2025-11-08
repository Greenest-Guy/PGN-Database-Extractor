package criteria

import "fmt"

/*
Skill Groups - Beginner, Intermediate, Expert
Time Controls - UltraBullet, Bullet, Blitz, Rapid, Classical
Maximum Elo Difference - The Maximum allowed difference between black's elo and white's elo
FileName - Name of output file
*/

var (
	TimeControl = "Rapid"
	SkillGroup  = "Intermediate"
	MaxEloDiff  = 150
	FileName    = fmt.Sprintf("%s_%s_%d.csv", TimeControl, SkillGroup, MaxEloDiff)
	NumGames    = 91549148 // Lichess monthly dump - 2025 October 29.9 GB
)
