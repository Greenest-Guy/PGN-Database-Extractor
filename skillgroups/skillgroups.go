package skillgroups

import (
	"fmt"
	"strconv"
)

func GetSkillGroup(BlackEloTag string, WhiteEloTag string) (string, error) {
	BlackElo, _ := strconv.Atoi(BlackEloTag)
	WhiteElo, _ := strconv.Atoi(WhiteEloTag)

	average := (BlackElo + WhiteElo) / 2

	if average < 1400 {
		return "Beginner", nil
	} else if 1400 <= average && average < 2000 {
		return "Intermediate", nil
	} else if 2000 <= average {
		return "Expert", nil
	}

	return "", fmt.Errorf("error calculating skill group")
}
