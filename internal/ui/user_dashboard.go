package ui

import (
	"bufio"
	"cli-project/pkg/globals"
	"cli-project/pkg/utils/formatting"
	"fmt"
	"os"
)

// ShowUserDashboard displays the user's LeetCode stats on the dashboard.
func (ui *UI) ShowUserDashboard() {
	// Clear the screen
	fmt.Print("\033[H\033[2J")

	fmt.Println(formatting.Colorize("====================================", "magenta", "bold"))
	fmt.Println(formatting.Colorize("            USER DASHBOARD          ", "magenta", "bold"))
	fmt.Println(formatting.Colorize("====================================", "magenta", "bold"))

	// Fetch LeetCode stats (assuming you have a method to get these stats)
	stats, err := ui.userService.GetLeetCodeStats(globals.ActiveUserID)
	if err != nil {
		fmt.Println("Error fetching stats:", err)
		return
	}

	// Display stats with color coding
	fmt.Println(formatting.Colorize("Questions solved", "cyan", "bold"))
	fmt.Println(formatting.Colorize(fmt.Sprintf("Easy : %d/%d", stats.EasyDoneCount, stats.TotalEasyCount), "green", ""))
	fmt.Println(formatting.Colorize(fmt.Sprintf("Medium : %d/%d", stats.MediumDoneCount, stats.TotalMediumCount), "yellow", ""))
	fmt.Println(formatting.Colorize(fmt.Sprintf("Hard : %d/%d", stats.HardDoneCount, stats.TotalHardCount), "red", ""))

	// Display recent AC submissions
	fmt.Println(formatting.Colorize("\nRecent Accepted Submissions", "cyan", "bold"))
	for _, submission := range stats.RecentACSubmissions {
		fmt.Println("- " + submission)
	}

	fmt.Println("\nPress any key to go back...")

	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
