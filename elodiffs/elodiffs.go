package elodiffs

import (
	"fmt"
	"strconv"
)

func GetEloDiff(blackEloTag string, whiteEloTag string) (int, error) {
	BlackElo, err := strconv.Atoi(blackEloTag)
	if err != nil {
		return 0, fmt.Errorf("invalid black elo tag")
	}

	WhiteElo, err := strconv.Atoi(whiteEloTag)
	if err != nil {
		return 0, fmt.Errorf("invalid white elo tag")
	}

	return AbsInt(BlackElo - WhiteElo), nil
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
