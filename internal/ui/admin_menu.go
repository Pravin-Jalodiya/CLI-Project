package ui

import (
	"cli-project/pkg/utils/formatting"
	"fmt"
	"strings"
)

func (ui *UI) ShowAdminMenu() {
	for {
		// Clear the screen
		fmt.Print("\033[H\033[2J")

		fmt.Println(formatting.Colorize("====================================", "cyan", "bold"))
		fmt.Println(formatting.Colorize("             ADMIN MENU             ", "cyan", "bold"))
		fmt.Println(formatting.Colorize("====================================", "cyan", "bold"))
		fmt.Println(formatting.Colorize("1. View platform stats", "", ""))
		fmt.Println(formatting.Colorize("2. Add or remove questions", "", ""))
		fmt.Println(formatting.Colorize("3. Manage users", "", ""))
		//fmt.Println(formatting.Colorize("4. Post Announcement", "", ""))
		fmt.Println(formatting.Colorize("4. Logout", "", ""))

		fmt.Print(formatting.Colorize("Enter your choice: ", "yellow", "bold"))
		choice, err := ui.reader.ReadString('\n')
		choice = strings.TrimSuffix(choice, "\n")
		choice = strings.TrimSpace(choice)

		if err != nil {
			fmt.Println(formatting.Colorize("Error reading input:", "red", "bold"), err)
			return
		}

		switch choice {
		case "1":
			ui.DisplayPlatformStats()
		case "2":
			ui.ManageQuestions()
		case "3":
			ui.ManageUsers()
		case "4":
			fmt.Println("Logging out...")
			return
		//case "4":
		//	ui.PostAnnouncement()
		default:
			fmt.Println(formatting.Colorize("Invalid choice. Please select a valid option.", "red", "bold"))
		}
	}
}
