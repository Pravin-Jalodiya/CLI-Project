package ui

import (
	"cli-project/pkg/utils/data_cleaning"
	"cli-project/pkg/utils/formatting"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
)

func (ui *UI) ViewQuestionsPage() {

	for {
		// Clear the screen
		fmt.Print("\033[H\033[2J")

		fmt.Println(formatting.Colorize("====================================", "cyan", "bold"))
		fmt.Println(formatting.Colorize("              QUESTIONS             ", "cyan", "bold"))
		fmt.Println(formatting.Colorize("====================================", "cyan", "bold"))
		fmt.Println("1. View questions")
		fmt.Println("2. Go back")
		fmt.Print("Enter your choice : ")

		// Read user input
		choice, _ := ui.reader.ReadString('\n')
		choice = strings.TrimSuffix(choice, "\n")
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			ui.ViewQuestions()
		case "2":
			return
		default:
			fmt.Println(formatting.Colorize("Invalid choice. Please select a valid option.", "red", "bold"))
		}

	}

}

func (ui *UI) ViewQuestions() {

	// Load all questions in the db
	questionsList, err := ui.questionService.GetAllQuestions()
	if err != nil {
		fmt.Println("Failed to load questions")
		return
	}

	// If no questions found, notify the user
	if len(*questionsList) == 0 {
		fmt.Println("Trouble loading questions. Try again later.")
		return
	}

	// Create a new table writer to format the output as a table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Difficulty", "Link", "Topic-Tags", "Company-Tags"})

	// Print table rows
	for _, question := range *questionsList {
		// Convert slices to comma-separated strings for display
		topicTags := strings.Join(question.TopicTags, ", ")
		companyTags := strings.Join(question.CompanyTags, ", ")

		// Add the row to the table
		table.Append([]string{
			question.QuestionID,
			question.QuestionTitle,
			question.Difficulty,
			question.QuestionLink,
			topicTags,
			companyTags,
		})
	}

	// Render the table to the console
	table.Render()

	// Prompt user for input
	fmt.Println("Press 'f' to view filtered questions or any other key to return")
	var input string
	_, err = fmt.Scanln(&input)
	if err != nil {
		return
	}

	// Check the input and call ViewFilteredQuestions if input is 'f'
	if strings.ToLower(input) == "f" {
		ui.ViewFilteredQuestions()
	}
}

func (ui *UI) ViewFilteredQuestions() {

	// Prompt for difficulty
	fmt.Print("Enter difficulty (press enter to skip): ")
	difficulty, _ := ui.reader.ReadString('\n')
	difficulty = strings.TrimSuffix(difficulty, "\n")
	difficulty = data_cleaning.CleanString(difficulty)

	// Prompt for topic
	fmt.Print("Enter topic (press enter to skip): ")
	topic, _ := ui.reader.ReadString('\n')
	topic = strings.TrimSuffix(topic, "\n")
	topic = data_cleaning.CleanString(topic)

	// Prompt for company
	fmt.Print("Enter company (press enter to skip")
	company, _ := ui.reader.ReadString('\n')
	company = strings.TrimSuffix(company, "\n")
	company = data_cleaning.CleanString(company)

	// Fetch filtered questions
	filteredQuestions, err := ui.questionService.GetQuestionsByFilters(difficulty, topic, company)
	if err != nil {
		fmt.Printf("Error fetching filtered questions: %v\n", err)
		return
	}

	// If no questions found, notify the user
	if len(*filteredQuestions) == 0 {
		fmt.Println(formatting.Colorize("No questions match the filter", "yellow", "bold"))
		return
	}

	// Create a new table writer to format the output as a table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Difficulty", "Link", "Topic-Tags", "Company-Tags"})

	// Print table rows
	for _, question := range *filteredQuestions {
		// Convert slices to comma-separated strings for display
		topicTags := strings.Join(question.TopicTags, ", ")
		companyTags := strings.Join(question.CompanyTags, ", ")

		// Add the row to the table
		table.Append([]string{
			question.QuestionID,
			question.QuestionTitle,
			question.Difficulty,
			question.QuestionLink,
			topicTags,
			companyTags,
		})
	}

	// Render the table to the console
	table.Render()

	// Prompt the user to go back
	fmt.Println("\nPress any key to go back...")

	_, _ = ui.reader.ReadString('\n')
}
