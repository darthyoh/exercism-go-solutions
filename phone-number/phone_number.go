package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

//Number normalize phone number
func Number(s string) (string, error) {
	re := regexp.MustCompile(`^(\+1|1)?[ -\.]*\(?([2-9]{1}[0-9]{2})\)?[ -\.]*([2-9]{1}[0-9]{2})[ -\.]*([0-9]{4})[ -\.]*$`)
	groups := re.FindStringSubmatch(s)
	if len(groups) == 0 {
		return "", errors.New("Invalid Number")
	}
	return fmt.Sprintf("%v%v%v", groups[2], groups[3], groups[4]), nil
}

//Format a phone number
func Format(s string) (string, error) {
	number, err := Number(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%v) %v-%v", number[0:3], number[3:6], number[6:]), nil
}

//AreaCode of a phone number
func AreaCode(s string) (string, error) {
	number, err := Number(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", number[0:3]), nil
}
