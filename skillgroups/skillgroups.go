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

	lower_bound := 1400
	upper_bound := 2000

	if average < lower_bound {
		return "Beginner", nil
	} else if average < upper_bound {
		return "Intermediate", nil
	}
	return "Expert", nil
}
