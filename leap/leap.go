// Package leap for Leap Year
package leap

// IsLeapYear is a simple function to know if a year is leap
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
