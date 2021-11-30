package clock

import (
	"fmt"
	"time"
)

//Clock struct simply store hour and minute of a Clock
type Clock struct {
	Hour   int
	Minute int
}

//useful function to get a Time with 0 hour and minute
func getTimeAt00() time.Time {
	start := time.Now()
	start = start.Add(time.Minute * time.Duration((start.Minute()+start.Hour()*60)*-1))
	return start
}

//New function create a new Clock at given hour and minute
func New(hour, minute int) Clock {
	start := getTimeAt00().Add(time.Minute * time.Duration(60*hour+minute))
	return Clock{start.Hour(), start.Minute()}
}

//String method to display the hh:mm Clock hour
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

//Add method to add minutes to a Clock
func (c Clock) Add(minutes int) Clock {
	start := getTimeAt00().Add(time.Minute * time.Duration(c.Minute+minutes+(60*c.Hour)))
	return Clock{start.Hour(), start.Minute()}
}

//Subtract method to subtract method to a Clock
func (c Clock) Subtract(minutes int) Clock {
	start := getTimeAt00().Add(time.Minute * time.Duration(c.Minute-minutes+(60*c.Hour)))
	return Clock{start.Hour(), start.Minute()}
}
