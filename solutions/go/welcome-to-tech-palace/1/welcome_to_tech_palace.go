package techpalace

import (
    "strings"
)
// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, "+strings.ToUpper(customer)
	panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    border := strings.Repeat("*",numStarsPerLine)
    return border+"\n"+welcomeMsg+"\n"+border
	panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	lines := strings.Split(oldMsg, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.Trim(line, "*") == "" {
			continue
		}

		return strings.Trim(line, "* ")
	}

	return ""
}
