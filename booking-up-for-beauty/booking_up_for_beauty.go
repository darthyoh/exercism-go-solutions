package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
    
    t,_ :=  time.Parse("1/02/2006 15:04:05",date)
    return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
now := time.Now()
parsed,_ := time.Parse("January 2, 2006 15:04:05",date)
    return parsed.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
    parsed,_ := time.Parse("Monday, January 2, 2006 15:04:05",date)
    return parsed.Hour() >= 12 && parsed.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	parsed,_ := time.Parse("1/2/2006 15:04:05",date)
    formated := parsed.Format("Monday, January 2, 2006, at 15:04.")
    return "You have an appointment on "+formated
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
    layout:="2006-01-02"
parsed,_ :=time.Parse(layout, fmt.Sprintf("%v-09-15",time.Now().Year()))
    return parsed
}
