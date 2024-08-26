package repositories

import (
	"cli-project/internal/config"
	"cli-project/internal/domain/interfaces"
	"cli-project/internal/domain/models"
	"cli-project/pkg/globals"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type userRepo struct {
	collection *mongo.Collection
}

func NewUserRepo() interfaces.UserRepository {
	return &userRepo{
		collection: client.Database(config.DB_NAME).Collection(config.USER_COLLECTION),
	}
}

func (r *userRepo) RegisterUser(user *models.StandardUser) error {
	// Convert the user model to BSON format
	userBson := bson.M{
		"id":               user.StandardUser.ID,
		"username":         user.StandardUser.Username,
		"password":         user.StandardUser.Password,
		"name":             user.StandardUser.Name,
		"email":            user.StandardUser.Email,
		"role":             user.StandardUser.Role,
		"organisation":     user.StandardUser.Organisation,
		"country":          user.StandardUser.Country,
		"leetcode_id":      user.LeetcodeID,
		"questions_solved": user.QuestionsSolved,
		"last_seen":        user.LastSeen,
	}

	// Insert the user document into the collection
	_, err := r.collection.InsertOne(context.TODO(), userBson)
	if err != nil {
		return fmt.Errorf("could not insert user: %v", err)
	}

	return nil
}

func (r *userRepo) UpdateUserProgress(solvedQuestionID string) error {
	// Set a context with a timeout for the database operation
	ctx, cancel := CreateContext()
	defer cancel()

	// Find the current user
	filter := bson.M{"id": globals.ActiveUserID}

	// Add the solved question ID to the QuestionsSolved slice
	update := bson.M{
		"$addToSet": bson.M{
			"questions_solved": solvedQuestionID,
		},
	}

	// Update the user document
	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update progress: %v", err)
	}

	return nil
}

func (r *userRepo) FetchAllUsers() (*[]models.StandardUser, error) {
	// Set a context with a timeout for the database operation
	ctx, cancel := CreateContext()
	defer cancel()

	// Define an empty filter to match all documents
	filter := bson.M{}

	// Find all users
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	var users []models.StandardUser

	// Iterate through the cursor and decode each document into a StandardUser
	for cursor.Next(ctx) {
		var user models.StandardUser
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Check if there were any errors during the iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *userRepo) FetchUserByID(userID string) (*models.StandardUser, error) {
	// Set a context with a timeout for the database operation
	ctx, cancel := CreateContext()
	defer cancel()

	filter := bson.M{"id": userID}

	var user models.StandardUser

	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &user, mongo.ErrNoDocuments // User not found
		}
		return &user, err
	}

	// Return the found user
	return &user, nil
}

func (r *userRepo) FetchUserByUsername(username string) (*models.StandardUser, error) {
	// Set a context with a timeout for the database operation
	ctx, cancel := CreateContext()
	defer cancel()

	filter := bson.M{"username": username}

	var user models.StandardUser

	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &user, mongo.ErrNoDocuments // User not found
		}
		return &user, err
	}

	// Return the found user
	return &user, nil
}

func (r *userRepo) UpdateUserDetails(user *models.StandardUser) error {
	// Check if user UUID is provided
	if user.StandardUser.ID == "" {
		return errors.New("user ID is required")
	}

	// Create a filter to find the user by ID
	filter := bson.M{"id": user.StandardUser.ID}

	// Define the update fields
	update := bson.M{
		"$set": bson.M{
			"username":         user.StandardUser.Username,
			"email":            user.StandardUser.Email,
			"password":         user.StandardUser.Password, // if user wants to change password
			"name":             user.StandardUser.Name,
			"organisation":     user.StandardUser.Organisation,
			"country":          user.StandardUser.Country,
			"leetcode_id":      user.LeetcodeID,
			"last_seen":        user.LastSeen,
			"questions_solved": user.QuestionsSolved,
			// Add other fields you want to update
		},
	}

	// Set options to return the updated document
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Update the document
	ctx, cancel := CreateContext()
	defer cancel()

	result := r.collection.FindOneAndUpdate(ctx, filter, update, opts)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

func (r *userRepo) CountActiveUsersInLast24Hours() (int64, error) {

	now := time.Now().UTC()
	twentyFourHoursAgo := now.Add(-24 * time.Hour)

	filter := bson.M{
		"last_seen": bson.M{
			"$gte": twentyFourHoursAgo,
		},
	}

	count, err := r.collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, fmt.Errorf("could not count active users: %v", err)
	}

	return count, nil
}

func (r *userRepo) IsEmailUnique(email string) (bool, error) {
	var result models.StandardUser
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func (r *userRepo) IsUsernameUnique(username string) (bool, error) {
	var result models.StandardUser
	err := r.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func (r *userRepo) IsLeetcodeIDUnique(leetcodeID string) (bool, error) {
	var result models.StandardUser
	err := r.collection.FindOne(context.Background(), bson.M{"leetcode_id": leetcodeID}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}
