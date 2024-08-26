package ui

import (
	"cli-project/pkg/utils"
	"cli-project/pkg/utils/data_cleaning"
	"cli-project/pkg/utils/formatting"
	"cli-project/pkg/validation"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func (ui *UI) ManageUsers() {
	for {
		// Clear the screen
		fmt.Print("\033[H\033[2J")

		fmt.Println(formatting.Colorize("====================================", "cyan", "bold"))
		fmt.Println(formatting.Colorize("           MANAGE USERS             ", "cyan", "bold"))
		fmt.Println(formatting.Colorize("====================================", "cyan", "bold"))
		fmt.Println(formatting.Colorize("1. View all users", "", ""))
		fmt.Println(formatting.Colorize("2. Ban a user", "", ""))
		fmt.Println(formatting.Colorize("3. Unban a user", "", ""))
		fmt.Println(formatting.Colorize("4. Go back", "", ""))

		fmt.Print(formatting.Colorize("Enter your choice: ", "yellow", "bold"))
		choice, err := ui.reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if err != nil {
			fmt.Println(formatting.Colorize("error reading input:", "red", "bold"), err)
			return
		}

		switch choice {
		case "1":
			ui.viewAllUsers()
		case "2":
			ui.banUser()
		case "3":
			ui.unbanUser()
		case "4":
			return // Go back to the previous menu
		default:
			fmt.Println(formatting.Colorize("invalid choice", "red", "bold"))
		}
	}
}

func (ui *UI) viewAllUsers() {
	// Get all users from the user service
	users, err := ui.userService.GetAllUsers()
	if err != nil {
		fmt.Println("failed to load users.")
		return
	}

	// If no users found, notify the admin
	if len(*users) == 0 {
		fmt.Println("no users found.")
		return
	}

	// Create a new tab writer to format the output as a table
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	// Print table headers
	_, err = fmt.Fprintln(w, "Username\tName\tEmail\tLeetcode ID\tOrganisation\tCountry\tIsBlocked\tLast Seen (IST)")
	if err != nil {
		fmt.Println("error rendering page.")
		return
	}

	// Print table rows, excluding admin users
	for _, user := range *users {
		if user.StandardUser.Role != "user" {
			continue
		}

		// Convert Last Seen time to IST and format it
		lastSeenIST := utils.ConvertToIST(user.LastSeen)

		// Format the user details into table rows
		_, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%t\t%s\n",
			user.StandardUser.Username,
			user.StandardUser.Name,
			user.StandardUser.Email,
			user.LeetcodeID,
			user.StandardUser.Organisation,
			user.StandardUser.Country,
			user.StandardUser.IsBanned,
			lastSeenIST,
		)

		if err != nil {
			fmt.Println("Error rendering page.")
			return
		}
	}

	// Flush the writer to ensure all output is printed
	err = w.Flush()
	if err != nil {
		fmt.Println("Error rendering page.")
		return
	}
}

func (ui *UI) banUser() {

	// View all users
	ui.viewAllUsers()

	// Logic to ban a user
	var username string
	var err error
	for {
		fmt.Print("Enter the username to ban: ")
		username, err = ui.reader.ReadString('\n')
		username = data_cleaning.CleanString(username)
		if err != nil {
			fmt.Println(formatting.Colorize("error reading input:", "red", "bold"), err)
			return
		}

		valid := validation.ValidateUsername(username)

		if !valid {
			fmt.Println(formatting.Colorize("enter a valid username", "yellow", "bold"))
			continue
		}
		break
	}

	// banning logic
	err = ui.userService.BanUser(username)
	if err != nil {
		fmt.Println(formatting.Colorize("error banning user: ", "red", "bold"), err)
		return
	}

	fmt.Println(formatting.Colorize("user banned successfully", "green", "bold"))
}

func (ui *UI) unbanUser() {

	// View all users
	ui.viewAllUsers()

	// Logic to unban a user
	var username string
	var err error
	for {
		fmt.Print("Enter the username to unban: ")
		username, err = ui.reader.ReadString('\n')
		username = data_cleaning.CleanString(username)
		if err != nil {
			fmt.Println(formatting.Colorize("error reading input:", "red", "bold"), err)
			return
		}

		valid := validation.ValidateUsername(username)

		if !valid {
			fmt.Println(formatting.Colorize("enter a valid username", "yellow", "bold"))
			continue
		}
		break
	}

	// Unbanning logic
	err = ui.userService.UnbanUser(username)
	if err != nil {
		fmt.Println(formatting.Colorize("error unbanning user: ", "red", "bold"), err)
		return
	}

	fmt.Println(formatting.Colorize("user unbanned successfully", "green", "bold"))
}
