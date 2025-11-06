package timecontrols

import (
	"fmt"
	"strconv"
	"strings"
)

/*
https://lichess.org/faq#time-controls
Lichess time controls are based on estimated game duration = (clock initial time in seconds) + 40 × (clock increment).

≤ 29s = UltraBullet
≤ 179s = Bullet
≤ 479s = Blitz
≤ 1499s = Rapid
≥ 1500s = Classical
*/

func estimatedGameDuration(timecontrol string) (int, error) {
	if timecontrol == "-" {
		return -1, nil
	}

	parts := strings.Split(timecontrol, "+")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid time control format: %s", timecontrol)
	}

	initialTime, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid initial time: %s", parts[0])
	}

	increment, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid increment: %s", parts[1])
	}

	return initialTime*60 + 40*increment, nil
}

func GetTimeControl(timecontrol string) (string, error) {
	/*
		≤ 29s = UltraBullet
		≤ 179s = Bullet
		≤ 479s = Blitz
		≤ 1499s = Rapid
		≥ 1500s = Classical
	*/

	est_duration, err := estimatedGameDuration(timecontrol)
	if err != nil {
		return "", err
	}

	if est_duration == -1 {
		return "-", nil
	}

	if est_duration <= 29 {
		return "UltraBullet", nil
	} else if 29 < est_duration && est_duration <= 179 {
		return "Bullet", nil
	} else if 179 < est_duration && est_duration <= 479 {
		return "Blitz", nil
	} else if 479 < est_duration && est_duration <= 1499 {
		return "Rapid", nil
	} else if est_duration >= 1500 {
		return "Classical", nil
	} else {
		return "", nil
	}
}
