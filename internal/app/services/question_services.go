package services

import (
	"cli-project/internal/domain/dto"
	"cli-project/internal/domain/interfaces"
	"cli-project/internal/domain/models"
	"cli-project/pkg/errors"
	"cli-project/pkg/utils"
	"cli-project/pkg/validation"
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type QuestionService struct {
	questionRepo interfaces.QuestionRepository
}

func NewQuestionService(questionRepo interfaces.QuestionRepository) interfaces.QuestionService {
	return &QuestionService{
		questionRepo: questionRepo,
	}
}

var (
	CSVReader                  = utils.ReadCSV
	ValidateQuestionID         = validation.ValidateQuestionID
	ValidateQuestionDifficulty = validation.ValidateQuestionDifficulty
	ValidateQuestionLink       = validation.ValidateQuestionLink
	ValidateTitleSlug          = validation.ValidateTitleSlug
)

// AddQuestionsFromFile processes the CSV file and adds new questions to the database.
func (s *QuestionService) AddQuestionsFromFile(ctx context.Context, questionFilePath string) (bool, error) {
	// Open the CSV file
	file, err := os.Open(questionFilePath)
	if err != nil {
		return false, fmt.Errorf("%w: %v", errs.ErrCSVFileOpening, err)
	}
	defer file.Close()

	// Read the CSV content
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return false, fmt.Errorf("%w: %v", errs.ErrReadingCSVFile, err)
	}

	var questions []models.Question
	newQuestionsAdded := false

	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}

		if len(record) != 7 {
			return false, errs.ErrInvalidCSVFormat
		}

		titleSlug := utils.CleanString(record[0])
		questionID := utils.CleanString(record[1])
		questionTitle := utils.CleanString(record[2])
		difficulty := record[3]
		questionLink := record[4]
		topicTags := utils.CleanTags(record[5])
		companyTags := utils.CleanTags(record[6])

		valid, err := ValidateQuestionID(questionID)
		if !valid {
			return false, fmt.Errorf("%w: %v", errs.ErrInvalidQuestionID, err)
		}

		difficulty, err = ValidateQuestionDifficulty(difficulty)
		if err != nil {
			return false, fmt.Errorf("%w: %v", errs.ErrInvalidQuestionDifficulty, err)
		}

		questionLink, err = ValidateQuestionLink(questionLink)
		if err != nil {
			return false, fmt.Errorf("%w: %v", errs.ErrInvalidQuestionLink, err)
		}

		question := models.Question{
			QuestionTitleSlug: titleSlug,
			QuestionID:        questionID,
			QuestionTitle:     questionTitle,
			Difficulty:        difficulty,
			QuestionLink:      questionLink,
			TopicTags:         topicTags,
			CompanyTags:       companyTags,
		}

		exists, err := s.QuestionExistsByID(ctx, questionID)
		if err != nil {
			return false, fmt.Errorf("%w: %v", errs.ErrDbOperation, err)
		}

		if exists {
			continue
		}

		questions = append(questions, question)
		newQuestionsAdded = true
	}

	if newQuestionsAdded {
		err = s.questionRepo.AddQuestions(ctx, &questions)
		if err != nil {
			return false, fmt.Errorf("%w: %v", errs.ErrDbOperation, err)
		}
	}

	return newQuestionsAdded, nil
}

func (s *QuestionService) RemoveQuestionByID(ctx context.Context, questionID string) error {
	exists, err := s.QuestionExistsByID(ctx, questionID)
	if err != nil {
		return fmt.Errorf("%w: %v", errs.ErrDbError, err)
	}

	if !exists {
		return fmt.Errorf("%w: question with ID %s not found", errs.ErrNoRows, questionID)
	}

	err = s.questionRepo.RemoveQuestionByID(ctx, questionID)
	if err != nil {
		return fmt.Errorf("%w: %v", errs.ErrDbError, err)
	}

	return nil
}

func (s *QuestionService) GetQuestionByID(ctx context.Context, questionID string) (*models.Question, error) {
	exists, err := s.QuestionExistsByTitleSlug(ctx, questionID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.ErrDbError, err)
	}
	if !exists {
		return nil, fmt.Errorf("%w: question with title slug %s not found", errs.ErrNoRows, questionID)
	}

	question, err := s.questionRepo.FetchQuestionByID(ctx, questionID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.ErrDbError, err)
	}

	return question, nil
}

func (s *QuestionService) GetAllQuestions(ctx context.Context) ([]dto.Question, error) {
	questions, err := s.questionRepo.FetchAllQuestions(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.ErrDbError, err)
	}
	return questions, nil
}

func (s *QuestionService) GetQuestionsByFilters(ctx context.Context, difficulty, topic, company string) ([]dto.Question, error) {
	var validDifficulty string
	var err error

	if difficulty != "" && strings.ToLower(difficulty) != "any" {
		validDifficulty, err = ValidateQuestionDifficulty(difficulty)
		if err != nil {
			return []dto.Question{}, fmt.Errorf("%w: invalid difficulty", errs.ErrInvalidParameterError)
		}
	}

	cleanCompany := utils.CleanString(company)
	cleanTopic := utils.CleanString(topic)

	questions, err := s.questionRepo.FetchQuestionsByFilters(ctx, validDifficulty, cleanTopic, cleanCompany)
	if err != nil {
		return []dto.Question{}, fmt.Errorf("%w: %v", errs.ErrDbError, err)
	}

	return questions, nil
}

func (s *QuestionService) QuestionExistsByID(ctx context.Context, questionID string) (bool, error) {
	valid, err := validation.ValidateQuestionID(questionID)
	if !valid {
		return false, fmt.Errorf("%w: %v", errs.ErrInvalidParameterError, err)
	}

	exists, err := s.questionRepo.QuestionExistsByID(ctx, questionID)
	if err != nil {
		return false, fmt.Errorf("%w: %v", errs.ErrDbError, err)
	}

	return exists, nil
}

func (s *QuestionService) QuestionExistsByTitleSlug(ctx context.Context, titleSlug string) (bool, error) {

	valid, err := ValidateTitleSlug(titleSlug)
	if !valid {
		return false, fmt.Errorf("%w: %v", errs.ErrInvalidParameterError, err)
	}

	exists, err := s.questionRepo.QuestionExistsByTitleSlug(ctx, titleSlug)
	if err != nil {
		return false, fmt.Errorf("%w: %v", errs.ErrDbError, err)
	}

	return exists, nil
}

func (s *QuestionService) GetTotalQuestionsCount(ctx context.Context) (int, error) {
	count, err := s.questionRepo.CountQuestions(ctx)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", errs.ErrDbError, err)
	}
	return count, nil
}
