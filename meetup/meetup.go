package meetup

import "time"
import "fmt"

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	Teenth WeekSchedule = iota
	First
	Second
	Third
	Fourth
	Last
)

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	format := "2006 1 2"
	parsed, _ := time.Parse(format, fmt.Sprintf("%d %d 1", year, month))

	for parsed.Weekday() != weekday {
		parsed = parsed.Add(time.Hour * 24)
	}
	if week == First {
		return parsed.Day()
	}
	days := make([]int, 0)

	for parsed.Month() == month {
		days = append(days, parsed.Day())
		parsed = parsed.Add(time.Hour * 24 * 7)
	}

	switch week {
	case Second:
		return days[1]
	case Third:
		return days[2]
	case Fourth:
		return days[3]
	case Last:
		return days[len(days)-1]
	default:
		teens := map[int]bool{13: true, 14: true, 15: true, 16: true, 17: true, 18: true, 19: true}
		for _, v := range days {
			if _, ok := teens[v]; ok {
				return v
			}
		}
	}

	return 0
}
