package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~\*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+([\w]+)\s+`)
	returned := make([]string, 0)

	for _, line := range lines {
		groups := re.FindAllStringSubmatch(line, -1)
		if len(groups) == 1 && len(groups[0]) == 2 {
			returned = append(returned, fmt.Sprintf("[USR] %s %s", groups[0][1], line))
		} else {
			returned = append(returned, line)
		}

	}
	return returned
}
