package utils

func UpdateTeam(team []string) ([]string){

	for i := 0; i < len(team); i++ {
		team[i] = "user:" + team[i]
	}

	return team
}