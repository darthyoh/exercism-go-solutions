package logs

import "strings"

// Application identifies the application emitting the given log.
func Application(log string) string {   
    for _,r := range log {
		if r == '❗' {
            return "recommendation"
        } 
    	if r == '🔍' {
            return "search"
        }
    	if r == '☀' {
            return "weather"
        }
    }
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    var str strings.Builder
    for _,r := range log {
        if r==oldRune {
            str.WriteRune(newRune)
        } else {
        	str.WriteRune(r)
        }
    }
	return str.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	long := 0
    for range log {
        long++
        if long > limit {
            return false
        }
    }
	return true
}
