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

const (
	UltraBulletThreshold = 29
	BulletThreshold      = 179
	BlitzThreshold       = 479
	RapidThreshold       = 1499
)

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

	// estimated game duration = (clock initial time in seconds) + 40 × (clock increment)
	return initialTime + 40*increment, nil
}

func GetTimeControl(timecontrol string) (string, error) {
	/*
		≤ 29s = UltraBullet
		≤ 179s = Bullet
		≤ 479s = Blitz
		≤ 1499s = Rapid
		≥ 1500s = Classical
	*/

	estDuration, err := estimatedGameDuration(timecontrol)
	if err != nil {
		return "", err
	}

	if estDuration == -1 {
		return "-", nil
	}

	if estDuration <= UltraBulletThreshold {
		return "UltraBullet", nil
	} else if estDuration <= BulletThreshold {
		return "Bullet", nil
	} else if estDuration <= BlitzThreshold {
		return "Blitz", nil
	} else if estDuration <= RapidThreshold {
		return "Rapid", nil
	} else {
		return "Classical", nil
	}
}
