package techpalace

import (
    "strings"
    "fmt"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return fmt.Sprintf("%[1]v\n%[2]v\n%[1]v", strings.Repeat("*",numStarsPerLine), welcomeMsg)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    return strings.TrimSpace(strings.ReplaceAll(strings.Split(oldMsg, "\n")[1], "*",""))    
}
