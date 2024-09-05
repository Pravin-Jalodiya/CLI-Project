package repositories

import (
	"cli-project/internal/domain/interfaces"
	"cli-project/internal/domain/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"strings"
	"time"
)

type userRepo struct {
}

func NewUserRepo() interfaces.UserRepository {
	return &userRepo{}
}

// getDBConnection returns a PostgreSQL client connection and handles errors.
func (r *userRepo) getDBConnection() (*sql.DB, error) {
	db, err := GetPostgresClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get PostgreSQL connection: %v", err)
	}
	return db, nil
}

func (r *userRepo) getTableName() string {
	return "users"
}

func (r *userRepo) CreateUser(user *models.StandardUser) error {
	// Get a database connection
	db, err := r.getDBConnection()
	if err != nil {
		return fmt.Errorf("failed to get DB connection: %v", err)
	}

	// Insert the user into the Users table
	query := `
		INSERT INTO Users (
			id, username, password, name, email, role, last_seen, organisation, country, leetcode_id, is_banned
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		)
	`
	_, err = db.ExecContext(context.TODO(), query,
		user.StandardUser.ID,
		strings.ToLower(user.StandardUser.Username),
		user.StandardUser.Password,
		user.StandardUser.Name,
		strings.ToLower(user.StandardUser.Email),
		user.StandardUser.Role,
		user.LastSeen,
		user.StandardUser.Organisation,
		user.StandardUser.Country,
		user.LeetcodeID,
		user.StandardUser.IsBanned,
	)
	if err != nil {
		return fmt.Errorf("could not insert user: %v", err)
	}

	return nil
}

func (r *userRepo) UpdateUserProgress(solvedQuestionID string) error {

	//collection, err := r.getCollection()
	//if err != nil {
	//	return fmt.Errorf("failed to get collection: %v", err)
	//}
	//// Set a context with a timeout for the database operation
	//ctx, cancel := CreateContext()
	//defer cancel()
	//
	//// Find the current user
	//filter := bson.M{"id": globals.ActiveUserID}
	//
	//// Add the solved question ID to the QuestionsSolved slice
	//update := bson.M{
	//	"$addToSet": bson.M{
	//		"questions_solved": solvedQuestionID,
	//	},
	//}
	//
	//// Update the user document
	//_, err = collection.UpdateOne(ctx, filter, update)
	//if err != nil {
	//	return fmt.Errorf("failed to update progress: %v", err)
	//}
	//
	return nil
}

func (r *userRepo) FetchAllUsers() (*[]models.StandardUser, error) {
	// Get a database connection
	db, err := r.getDBConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to get DB connection: %v", err)
	}

	// Define the query to fetch all users
	query := `
		SELECT
			id, username, password, name, email, role, last_seen, organisation, country, leetcode_id, is_banned
		FROM Users
	`

	// Execute the query
	rows, err := db.QueryContext(context.TODO(), query)
	if err != nil {
		return nil, fmt.Errorf("could not fetch users: %v", err)
	}
	defer rows.Close()

	var users []models.StandardUser

	// Iterate through the result set
	for rows.Next() {
		var user models.StandardUser
		err := rows.Scan(
			&user.StandardUser.ID,
			&user.StandardUser.Username,
			&user.StandardUser.Password,
			&user.StandardUser.Name,
			&user.StandardUser.Email,
			&user.StandardUser.Role,
			&user.LastSeen,
			&user.StandardUser.Organisation,
			&user.StandardUser.Country,
			&user.LeetcodeID,
			&user.StandardUser.IsBanned,
		)
		if err != nil {
			return nil, fmt.Errorf("could not scan user: %v", err)
		}
		users = append(users, user)
	}

	// Check if there were any errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over users: %v", err)
	}

	return &users, nil
}

func (r *userRepo) FetchUserByID(userID string) (*models.StandardUser, error) {
	// Get a database connection
	db, err := r.getDBConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to get DB connection: %v", err)
	}

	// Define the query to fetch a user by ID
	query := `
		SELECT
			id, username, password, name, email, role, last_seen, organisation, country, leetcode_id, is_banned
		FROM Users
		WHERE id = $1
	`

	// Execute the query
	row := db.QueryRowContext(context.TODO(), query, userID)

	var user models.StandardUser
	err = row.Scan(
		&user.StandardUser.ID,
		&user.StandardUser.Username,
		&user.StandardUser.Password,
		&user.StandardUser.Name,
		&user.StandardUser.Email,
		&user.StandardUser.Role,
		&user.LastSeen,
		&user.StandardUser.Organisation,
		&user.StandardUser.Country,
		&user.LeetcodeID,
		&user.StandardUser.IsBanned,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found") // User not found
		}
		return nil, fmt.Errorf("could not fetch user: %v", err)
	}

	// Return the found user
	return &user, nil
}

func (r *userRepo) FetchUserByUsername(username string) (*models.StandardUser, error) {
	// Get a database connection
	db, err := r.getDBConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to get DB connection: %v", err)
	}

	// Define the query to fetch a user by username
	query := `
		SELECT
			id, username, password, name, email, role, last_seen, organisation, country, leetcode_id, is_banned
		FROM Users
		WHERE username = $1
	`

	// Execute the query
	row := db.QueryRowContext(context.TODO(), query, username)

	var user models.StandardUser
	err = row.Scan(
		&user.StandardUser.ID,
		&user.StandardUser.Username,
		&user.StandardUser.Password,
		&user.StandardUser.Name,
		&user.StandardUser.Email,
		&user.StandardUser.Role,
		&user.LastSeen,
		&user.StandardUser.Organisation,
		&user.StandardUser.Country,
		&user.LeetcodeID,
		&user.StandardUser.IsBanned,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found") // User not found
		}
		return nil, fmt.Errorf("could not fetch user: %v", err)
	}

	// Return the found user
	return &user, nil
}

func (r *userRepo) UpdateUserDetails(user *models.StandardUser) error {
	// Get a database connection
	db, err := r.getDBConnection()
	if err != nil {
		return fmt.Errorf("failed to get DB connection: %v", err)
	}

	// Check if user UUID is provided
	if user.StandardUser.ID == "" {
		return errors.New("user ID is required")
	}

	// Define the SQL query to update user details
	query := `
		UPDATE Users
		SET 
			username = $1,
			email = $2,
			password = $3,
			name = $4,
			organisation = $5,
			country = $6,
			leetcode_id = $7,
			last_seen = $8
		WHERE id = $9
	`

	// Execute the update query
	_, err = db.ExecContext(
		context.TODO(),
		query,
		user.StandardUser.Username,
		user.StandardUser.Email,
		user.StandardUser.Password, // if user wants to change password
		user.StandardUser.Name,
		user.StandardUser.Organisation,
		user.StandardUser.Country,
		user.LeetcodeID,
		user.LastSeen,
		user.StandardUser.ID,
	)
	if err != nil {
		return fmt.Errorf("could not update user details: %v", err)
	}

	return nil
}

func (r *userRepo) BanUser(userID string) error {
	// Get a database connection
	db, err := r.getDBConnection()
	if err != nil {
		return fmt.Errorf("failed to get DB connection: %v", err)
	}

	// Check if user ID is provided
	if userID == "" {
		return errors.New("user ID is required")
	}

	// Define the SQL query to ban the user
	query := `
		UPDATE Users
		SET is_banned = TRUE
		WHERE id = $1
	`

	// Execute the update query
	result, err := db.ExecContext(context.TODO(), query, userID)
	if err != nil {
		return fmt.Errorf("could not ban user: %v", err)
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %s not found", userID)
	}

	return nil
}

func (r *userRepo) UnbanUser(userID string) error {
	// Get a database connection
	db, err := r.getDBConnection()
	if err != nil {
		return fmt.Errorf("failed to get DB connection: %v", err)
	}

	// Check if user ID is provided
	if userID == "" {
		return errors.New("user ID is required")
	}

	// Define the SQL query to unban the user
	query := `
		UPDATE Users
		SET is_banned = FALSE
		WHERE id = $1
	`

	// Execute the update query
	result, err := db.ExecContext(context.TODO(), query, userID)
	if err != nil {
		return fmt.Errorf("could not unban user: %v", err)
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %s not found", userID)
	}

	return nil
}

func (r *userRepo) CountActiveUsersInLast24Hours() (int64, error) {
	// Get a database connection
	db, err := r.getDBConnection()
	if err != nil {
		return 0, fmt.Errorf("failed to get DB connection: %v", err)
	}

	// Define the time range for the last 24 hours
	now := time.Now().UTC()
	twentyFourHoursAgo := now.Add(-24 * time.Hour)

	// Define the SQL query to count active users in the last 24 hours
	query := `
		SELECT COUNT(*)
		FROM Users
		WHERE last_seen >= $1
	`

	// Execute the query
	row := db.QueryRowContext(context.TODO(), query, twentyFourHoursAgo)

	// Scan the result into a count variable
	var count int64
	err = row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not count active users: %v", err)
	}

	return count, nil
}

func (r *userRepo) IsEmailUnique(email string) (bool, error) {
	// Get a database connection
	db, err := r.getDBConnection()
	if err != nil {
		return false, fmt.Errorf("failed to get DB connection: %v", err)
	}

	// Define the SQL query to check for the existence of the email
	query := `
		SELECT COUNT(*)
		FROM Users
		WHERE email = $1
	`

	// Execute the query
	row := db.QueryRowContext(context.TODO(), query, email)

	// Scan the result into a count variable
	var count int
	err = row.Scan(&count)
	if err != nil {
		return false, fmt.Errorf("could not check email uniqueness: %v", err)
	}

	// Return true if count is 0, meaning the email is unique
	return count == 0, nil
}

func (r *userRepo) IsUsernameUnique(username string) (bool, error) {
	// Get a database connection
	db, err := r.getDBConnection()
	if err != nil {
		return false, fmt.Errorf("failed to get DB connection: %v", err)
	}

	// Define the SQL query to check for the existence of the username
	query := `
		SELECT COUNT(*)
		FROM users
		WHERE username = $1
	`

	// Execute the query
	row := db.QueryRowContext(context.TODO(), query, username)

	// Scan the result into a count variable
	var count int
	err = row.Scan(&count)
	if err != nil {
		return false, fmt.Errorf("could not check username uniqueness: %v", err)
	}

	// Return true if count is 0, meaning the username is unique
	return count == 0, nil
}

func (r *userRepo) IsLeetcodeIDUnique(LeetcodeID string) (bool, error) {
	// Get a database connection
	db, err := r.getDBConnection()
	if err != nil {
		return false, fmt.Errorf("failed to get DB connection: %v", err)
	}

	// Define the SQL query to check for the existence of the LeetcodeID
	query := `
		SELECT COUNT(*)
		FROM Users
		WHERE leetcode_id = $1
	`

	// Execute the query
	row := db.QueryRowContext(context.TODO(), query, LeetcodeID)

	// Scan the result into a count variable
	var count int
	err = row.Scan(&count)
	if err != nil {
		return false, fmt.Errorf("could not check LeetcodeID uniqueness: %v", err)
	}

	// Return true if count is 0, meaning the LeetcodeID is unique
	return count == 0, nil
}
