package skillgroups

import (
	"fmt"
	"strconv"
)

func GetSkillGroup(blackEloTag string, whiteEloTag string) (string, error) {
	BlackElo, err := strconv.Atoi(blackEloTag)
	if err != nil {
		return "", fmt.Errorf("invalid black elo '%s': %w", blackEloTag, err)
	}

	WhiteElo, err := strconv.Atoi(whiteEloTag)
	if err != nil {
		return "", fmt.Errorf("invalid white elo '%s': %w", whiteEloTag, err)
	}

	average := (BlackElo + WhiteElo) / 2

	if average < 1400 {
		return "Beginner", nil
	} else if average < 2000 {
		return "Intermediate", nil
	} else {
		return "Expert", nil
	}
}
