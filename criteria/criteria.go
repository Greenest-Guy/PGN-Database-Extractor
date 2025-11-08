package criteria

import "fmt"

/*
Time Controls - UltraBullet, Bullet, Blitz, Rapid, Classical
Skill Groups - Beginner, Intermediate, Expert
Maximum Elo Difference - The Maximum allowed difference between black's elo and white's elo
FileName - Name of output file
*/

var (
	TimeControl = "Classical"
	SkillGroup  = "Expert"
	MaxEloDiff  = 100
	FileName    = fmt.Sprintf("%s_%s_%d.csv", TimeControl, SkillGroup, MaxEloDiff)
)
