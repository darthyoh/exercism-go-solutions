package tournament

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type Team struct {
	name  string
	play  int
	won   int
	drawn int
	lost  int
}

func (t Team) points() int {
	return t.won*3 + t.drawn
}

func Tally(reader io.Reader, buffer io.Writer) error {
	var teams = map[string]Team{}
	content := ""
	if b, err := ioutil.ReadAll(reader); err == nil {
		content = string(b)
	}
	lines := strings.Split(content, "\n")
	for _, match := range lines {

		if match != "" && match[0] != 35 {

			result := strings.Split(match, ";")

			if len(result) != 3 {
				return fmt.Errorf("Error")
			}

			if result[0] == result[1] {
				return fmt.Errorf("Error")
			}

			var team1, team2 Team

			if team, ok := teams[result[0]]; ok {
				team1 = team
			} else {
				team1 = Team{name: result[0]}
			}

			if team, ok := teams[result[1]]; ok {
				team2 = team
			} else {
				team2 = Team{name: result[1]}
			}

			team1.play++
			team2.play++

			switch result[2] {
			case "win":
				team1.won++
				team2.lost++
			case "draw":
				team1.drawn++
				team2.drawn++
			case "loss":
				team2.won++
				team1.lost++
			default:
				return fmt.Errorf("Error")
			}

			teams[result[0]] = team1
			teams[result[1]] = team2
		}
	}

	arr := make([]Team, 0, len(teams))
	for _, team := range teams {
		arr = append(arr, team)
	}
	sort.SliceStable(arr, func(i, j int) bool {
		if arr[i].points() != arr[j].points() {
			return arr[i].points() > arr[j].points()
		} else if arr[i].play != arr[j].play {
			return arr[i].play > arr[j].play
		}

		return arr[i].name < arr[j].name

	})

	result := fmt.Sprintf("%-31v| MP |  W |  D |  L |  P\n", "Team")
	for _, team := range arr {
		result += fmt.Sprintf("%-31v|%3v |%3v |%3v |%3v |%3v\n", team.name, team.play, team.won, team.drawn, team.lost, team.points())
	}

	if _, err := io.WriteString(buffer, result); err != nil {
		return fmt.Errorf("Error")
	}
	return nil
}
