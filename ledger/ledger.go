package ledger

/*
- Step 1 : rewrite header formatting with fmt.Sprintf
- Step 2 : rewrite empty entries case
- Step 3 : rewrite invalid Currency case
- Step 4 : rewrite sorting of entries with sort package
- Step 5 : rewrite date formatting
- Step 6 : rewrite change formatting
*/

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

//Entry struct represents an accounting entry
type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

//FormatLedger format entries
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	var entriesCopy []Entry
	for _, e := range entries {
		entriesCopy = append(entriesCopy, e)
	}

	/*
		Step 3 : rewrite invalid currency case
	*/
	if currency == "" || (currency != "USD" && currency != "EUR") {
		return "", errors.New("Invalid currency")
	}

	/*
		Step 4 : rewrite sorting of entries with sort package
	*/
	sort.Slice(entriesCopy, func(i, j int) bool {
		if entriesCopy[i].Date != entriesCopy[j].Date {
			return entriesCopy[i].Date < entriesCopy[j].Date
		}
		return entriesCopy[i].Change < entriesCopy[j].Change

	})

	var s string

	/*
		Step 1 : use fmt.Printf for formatting header
	*/

	if locale != "nl-NL" && locale != "en-US" {
		return "", errors.New("No valid locale")
	}
	dateStr, descriptionStr, changeStr := "Date", "Description", "Change"
	if locale == "nl-NL" {
		dateStr, descriptionStr, changeStr = "Datum", "Omschrijving", "Verandering"
	}
	s = fmt.Sprintf("%-11v| %-26v| %v\n", dateStr, descriptionStr, changeStr)

	/*
		Step 2 : rewrite empty entries case
	*/

	if len(entries) == 0 {
		return s, nil
	}

	// Parallelism, always a great idea
	co := make(chan struct {
		i int
		s string
		e error
	})
	//regexp used for step 5
	re := regexp.MustCompile(`^([0-9]{4})-([0-9]{2})-([0-9]{2})$`)
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {

			/*
				Step 5 : rewrite date formatting
			*/
			groups := re.FindStringSubmatch(entry.Date)
			if len(groups) == 0 {
				co <- struct {
					i int
					s string
					e error
				}{e: errors.New("")}
				return
			}
			year := groups[1]
			month := groups[2]
			day := groups[3]
			var d string
			if locale == "nl-NL" {
				d = fmt.Sprintf("%v-%v-%v", day, month, year)
			} else {
				d = fmt.Sprintf("%v/%v/%v", month, day, year)
			}

			de := entry.Description
			if len(de) > 25 {
				de = de[:22] + "..."
			} else {
				de = de + strings.Repeat(" ", 25-len(de))
			}

			/*
				Step 6 : rewrite change formatting
			*/

			negative := false
			cents := entry.Change
			if cents < 0 {
				cents = cents * -1
				negative = true
			}

			changeStrTemp := fmt.Sprintf("%.2f", float64(cents)/float64(100))

			//add separators
			changeStr := changeStrTemp[len(changeStrTemp)-3:]
			for i := len(changeStrTemp) - 4; i >= 0; i-- {
				if len(changeStrTemp)-4-i%3 == 3 && len(changeStrTemp)-4-i != 0 {
					changeStr = "_" + changeStr
				}
				changeStr = string(changeStrTemp[i]) + changeStr
			}

			//format separators with correct locale
			if locale == "nl-NL" {
				changeStr = " " + changeStr
				changeStr = strings.ReplaceAll(changeStr, ".", ",")
				changeStr = strings.ReplaceAll(changeStr, "_", ".")
			} else {
				changeStr = strings.ReplaceAll(changeStr, "_", ",")
			}

			//format currency symbol
			if currency == "EUR" {
				changeStr = "€" + changeStr
			} else {
				changeStr = "$" + changeStr
			}

			//format negative symbol
			if negative {
				if locale == "nl-NL" {
					changeStr = changeStr + "-"
				} else {
					changeStr = fmt.Sprintf("(%v)", changeStr)
				}
			} else {
				changeStr = changeStr + " "
			}

			//format final line string
			finalStr := fmt.Sprintf("%-11v| %-26v| %13v\n", d, de, changeStr)

			co <- struct {
				i int
				s string
				e error
			}{i: i, s: finalStr}
		}(i, et)
	}
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}
	for i := 0; i < len(entriesCopy); i++ {
		s += ss[i]
	}
	return s, nil
}
