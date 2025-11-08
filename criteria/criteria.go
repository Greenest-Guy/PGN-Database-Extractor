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
	SkillGroup  = "Beginner"
	MaxEloDiff  = 150
	FileName    = fmt.Sprintf("%s_%s_%d.csv", TimeControl, SkillGroup, MaxEloDiff)
)
