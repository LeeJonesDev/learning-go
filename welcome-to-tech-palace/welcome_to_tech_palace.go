package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starsBorder := strings.Repeat("*", numStarsPerLine)

	return starsBorder + "\n" + welcomeMsg + "\n" + starsBorder

}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {

	removedStars := strings.ReplaceAll(oldMsg, "*", "")
	removedLinebreaks := strings.ReplaceAll(removedStars, "\n", "")

	return strings.Trim(removedLinebreaks, " ")
}