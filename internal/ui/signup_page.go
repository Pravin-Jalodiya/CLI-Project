package ui

import (
	"cli-project/internal/domain/models"
	"cli-project/pkg/utils/data_cleaning"
	"cli-project/pkg/utils/emojis"
	"cli-project/pkg/utils/formatting"
	"cli-project/pkg/validation"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

func (ui *UI) ShowSignupPage() {
	// Clear the screen
	fmt.Print("\033[H\033[2J")

	fmt.Println(formatting.Colorize("====================================", "cyan", "bold"))
	fmt.Println(formatting.Colorize("               Signup               ", "cyan", "bold"))
	fmt.Println(formatting.Colorize("====================================", "cyan", "bold"))

	// Read Username
	var username string
	for {
		fmt.Print(formatting.Colorize("Username: ", "yellow", ""))
		username, _ = ui.reader.ReadString('\n')
		username = strings.TrimSuffix(username, "\n")
		username = data_cleaning.CleanString(username)

		if validation.ValidateUsername(username) {
			unique, err := ui.authService.IsUsernameUnique(username)
			if err != nil {
				fmt.Println(emojis.Error, "Error checking username uniqueness. Try again.")
				continue
			}
			if !unique {
				fmt.Println(emojis.Info, "Username already taken. Choose another username.")
				continue
			}

		} else {
			fmt.Println(emojis.Error, "Invalid username. It should be between 4 and 20 characters long, should not be only numbers and contain no spaces.")
			continue
		}
		break
	}

	// Read Password
	var password, confirmPassword string
	for {
		fmt.Print(formatting.Colorize("Password: ", "yellow", ""))
		passwordBytes, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
		password = string(passwordBytes)
		password = strings.TrimSpace(password)
		fmt.Println()

		// Read Confirm Password
		fmt.Print(formatting.Colorize("Confirm Password: ", "yellow", ""))
		confirmPasswordBytes, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
		confirmPassword = string(confirmPasswordBytes)
		confirmPassword = strings.TrimSpace(confirmPassword)
		fmt.Println()

		if password != confirmPassword {
			fmt.Println(emojis.Error, "Passwords do not match. Please try again.")
			continue
		}

		if !validation.ValidatePassword(password) {
			fmt.Println(emojis.Error, "Invalid password. It must be at least 8 characters long and include at least 1 uppercase & lowercase letters, 1 digit, and 1 special character.")
			continue
		}

		break
	}

	// Read Name
	var name string
	for {
		fmt.Print(formatting.Colorize("Name: ", "yellow", ""))
		name, _ = ui.reader.ReadString('\n')
		name = strings.TrimSuffix(name, "\n")
		name = strings.TrimSpace(name)
		if !validation.ValidateName(name) {
			fmt.Println(emojis.Error, "Invalid name. It should be 3 to 30 characters long and contain only letters and spaces.")
			continue
		}
		break
	}

	// Read Email
	var email string
	for {
		fmt.Print(formatting.Colorize("Email: ", "yellow", ""))
		email, _ = ui.reader.ReadString('\n')
		email = strings.TrimSuffix(email, "\n")
		email = data_cleaning.CleanString(email)

		if check1, check2 := validation.ValidateEmail(email); check1 == true && check2 == true {
			unique, err := ui.authService.IsEmailUnique(email)
			if err != nil {
				fmt.Println(emojis.Error, "Error checking email uniqueness. Try again.")
				continue
			}
			if !unique {
				fmt.Println(emojis.Info, "Email already registered. Use a different email.")
				continue
			}
		} else if check1 == false {
			fmt.Println(emojis.Error, "Invalid email format.")
			continue
		} else if check2 == false {
			fmt.Println(emojis.Error, "Invalid email domain. We only support gmail, outlook, yahoo, hotmail, icloud, watchguard emails.")
			continue
		}
		break
	}

	// Read Leetcode Username
	var LeetcodeID string
	for {
		fmt.Print(formatting.Colorize("Leetcode Username: ", "yellow", ""))
		LeetcodeID, _ = ui.reader.ReadString('\n')
		LeetcodeID = strings.TrimSuffix(LeetcodeID, "\n")
		LeetcodeID = strings.TrimSpace(LeetcodeID)

		// Check if Leetcode ID is unique in the database
		isUnique, err := ui.authService.IsLeetcodeIDUnique(LeetcodeID)

		if err != nil {
			fmt.Println(emojis.Error, "Error checking Leetcode ID uniqueness. Try again.")
			continue
		}
		if !isUnique {
			fmt.Println(emojis.Error, "Leetcode ID is already taken. Please choose a different ID.")
			continue
		}

		// Validate Leetcode Username with Leetcode API
		exists, err := ui.authService.ValidateLeetcodeUsername(LeetcodeID)
		if err != nil {
			fmt.Println(emojis.Error, "Error validating Leetcode username:", err)
			continue
		}
		if !exists {
			fmt.Println(emojis.Error, "Leetcode username does not exist.")
			continue
		}
		break
	}

	var organisation string
	for {
		fmt.Print(formatting.Colorize("Organisation: ", "yellow", ""))
		organisation, _ = ui.reader.ReadString('\n')
		organisation = strings.TrimSuffix(organisation, "\n")
		organisation = data_cleaning.CleanString(organisation)
		valid, err := validation.ValidateOrganizationName(organisation)
		if !valid {
			fmt.Println(emojis.Error, err)
			continue
		}
		break
	}

	var country string
	for {
		fmt.Print(formatting.Colorize("Country: ", "yellow", ""))
		country, _ = ui.reader.ReadString('\n')
		country = strings.TrimSuffix(country, "\n")
		country = data_cleaning.CleanString(country)
		valid, err := validation.ValidateCountryName(country)
		if !valid {
			fmt.Println(emojis.Error, err)
			continue
		}
		break
	}

	// Create User Object
	user := models.StandardUser{
		StandardUser: models.User{
			Username:     username,
			Password:     password,
			Name:         name,
			Email:        email,
			Organisation: organisation,
			Country:      country,
		},
		LeetcodeID: LeetcodeID,
	}

	// Call Signup Service
	err := ui.userService.Signup(&user)
	if err != nil {
		fmt.Println(emojis.Error, "Signup failed:", err)
		return
	}

	fmt.Println(emojis.Success, "Signup successful!")

	return
}
