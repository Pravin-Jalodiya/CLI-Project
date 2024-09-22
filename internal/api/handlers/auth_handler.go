package handlers

import (
	"cli-project/internal/domain/interfaces"
	"cli-project/internal/domain/models"
	"cli-project/pkg/errors"
	"cli-project/pkg/globals"
	"cli-project/pkg/logger"
	"cli-project/pkg/utils"
	"cli-project/pkg/validation"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

type AuthHandler struct {
	authService interfaces.AuthService
}

func NewAuthHandler(authService interfaces.AuthService) *AuthHandler {
	validate = validator.New()
	return &AuthHandler{
		authService: authService,
	}
}

func (a *AuthHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		StandardUser struct {
			Username     string `json:"username"`
			Password     string `json:"password"`
			Name         string `json:"name"`
			Email        string `json:"email"`
			Organization string `json:"organisation"`
			Country      string `json:"country"`
		} `json:"standard_user"`
		LeetcodeID string `json:"leetcode_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errs.NewInvalidRequestBodyError("Invalid request body").ToJSON(w)
		logger.Logger.Errorw("Error decoding request body", "method", r.Method, "error", err, "time", time.Now())
		return
	}

	user := models.StandardUser{
		StandardUser: models.User{
			Username:     input.StandardUser.Username,
			Password:     input.StandardUser.Password,
			Name:         input.StandardUser.Name,
			Email:        input.StandardUser.Email,
			Organisation: input.StandardUser.Organization,
			Country:      input.StandardUser.Country,
		},
		LeetcodeID:      input.LeetcodeID,
		QuestionsSolved: []string{},
		LastSeen:        time.Time{},
	}

	// Perform custom validations
	if !validation.ValidateUsername(user.StandardUser.Username) {
		errs.NewBadRequestError("Invalid username").ToJSON(w)
		return
	}
	if !validation.ValidatePassword(user.StandardUser.Password) {
		errs.NewBadRequestError("Invalid password").ToJSON(w)
		return
	}
	if !validation.ValidateName(user.StandardUser.Name) {
		errs.NewBadRequestError("Invalid name").ToJSON(w)
		return
	}
	isEmailValid, isReputable := validation.ValidateEmail(user.StandardUser.Email)
	if !isEmailValid {
		errs.NewBadRequestError("Invalid email format").ToJSON(w)
		return
	} else if !isReputable {
		errs.NewBadRequestError("Email domain is not reputable").ToJSON(w)
		return
	}
	isOrgValid, orgErr := validation.ValidateOrganizationName(user.StandardUser.Organisation)
	if !isOrgValid {
		errs.NewBadRequestError(orgErr.Error()).ToJSON(w)
		return
	}
	isCountryValid, countryErr := validation.ValidateCountryName(user.StandardUser.Country)
	if !isCountryValid {
		errs.NewBadRequestError(countryErr.Error()).ToJSON(w)
		return
	}

	err = a.authService.Signup(r.Context(), &user)
	if err != nil {
		if errors.Is(err, errs.ErrUserNameAlreadyExists) {
			errs.NewConflictError("User already exists").ToJSON(w)
		} else if errors.Is(err, errs.ErrEmailAlreadyExists) {
			errs.NewConflictError("Email already registered").ToJSON(w)
		} else if errors.Is(err, errs.ErrLeetcodeIDAlreadyExists) {
			errs.NewConflictError("LeetcodeID already registered").ToJSON(w)
		} else {
			errs.NewInternalServerError("Signup failed").ToJSON(w)
		}
		logger.Logger.Errorw("Signup failed", "method", r.Method, "error", err, "time", time.Now())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse := map[string]interface{}{
		"message": "User successfully registered",
		"code":    http.StatusOK,
		"user_info": map[string]any{
			"username":     user.StandardUser.Username,
			"organisation": user.StandardUser.Organisation,
			"country":      user.StandardUser.Country,
		},
	}
	json.NewEncoder(w).Encode(jsonResponse)
	logger.Logger.Infow("Signup Successful", "method", r.Method, "username", user.StandardUser.Username, "time", time.Now())
}

func (a *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		errs.NewInvalidRequestBodyError("Invalid request body").ToJSON(w)
		logger.Logger.Errorw("Error decoding request body", "method", r.Method, "error", err, "time", time.Now())
		return
	}

	// Validate the request body
	if !validation.ValidateUsername(requestBody.Username) {
		errs.NewBadRequestError("Invalid username").ToJSON(w)
		return
	}
	if !validation.ValidatePassword(requestBody.Password) {
		errs.NewBadRequestError("Invalid username or password").ToJSON(w)
		return
	}

	user, err := a.authService.Login(r.Context(), requestBody.Username, requestBody.Password)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, errs.ErrInvalidPassword) {
			errs.NewInvalidRequestBodyError("Invalid username or password").ToJSON(w)
			w.WriteHeader(http.StatusBadRequest)
		} else if errors.Is(err, errs.ErrUserNotFound) {
			errs.NewNotFoundError("User not found").ToJSON(w)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			errs.NewInternalServerError("Login failed").ToJSON(w)
		}
		logger.Logger.Errorw("Authentication failed", "method", r.Method, "error", err, "username", requestBody.Username, "time", time.Now())
		return
	}

	token, err := utils.CreateJwtToken(user.StandardUser.Username, user.StandardUser.ID, user.StandardUser.Role)
	if err != nil {
		errs.NewInternalServerError("Failed to generate token").ToJSON(w)
		return
	}

	globals.ActiveUserID = user.StandardUser.ID

	w.Header().Set("Content-Type", "application/json")
	jsonResponse := map[string]any{"code": http.StatusOK, "message": "Login successful", "token": token, "role": user.StandardUser.Role}
	json.NewEncoder(w).Encode(jsonResponse)
	logger.Logger.Infow("Login Successful", "method", r.Method, "request", requestBody, "time", time.Now())
}

func (a *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	err := a.authService.Logout(r.Context())
	if err != nil {
		errs.NewInternalServerError(err.Error()).ToJSON(w)
		logger.Logger.Errorw("Logout failed", "method", r.Method, "error", err, "time", time.Now())
		return
	}

	globals.ActiveUserID = ""

	w.Header().Set("Content-Type", "application/json")
	jsonResponse := map[string]any{"code": http.StatusOK, "message": "Logout successful"}
	json.NewEncoder(w).Encode(jsonResponse)
	logger.Logger.Infow("Logout Successful", "method", r.Method, "time", time.Now())
}
