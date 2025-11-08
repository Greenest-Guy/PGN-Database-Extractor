package elodiffs

import (
	"fmt"
	"strconv"
)

func GetEloDiff(BlackEloTag string, WhiteEloTag string) (int, error) {
	BlackElo, err := strconv.Atoi(BlackEloTag)
	if err != nil {
		return 0, fmt.Errorf("invalid black elo tag")
	}

	WhiteElo, err := strconv.Atoi(WhiteEloTag)
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
