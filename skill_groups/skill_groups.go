package skillgroups

func GetSkillGroup(BlackElo int, WhiteElo int) (string, error) {
	average := (BlackElo + WhiteElo) / 2
	if average < 1400 {
		return "Beginner", nil
	} else if 1400 <= average && average < 2000 {
		return "Intermediate", nil
	} else if 2000 <= average {
		return "Expert", nil
	}

	return "", nil
}
