package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kozhamseitova/test-task/api"
	"github.com/kozhamseitova/test-task/internal/config"
	"github.com/kozhamseitova/test-task/internal/entity"
	fakerepo "github.com/kozhamseitova/test-task/internal/repository/fake_repo"
	"github.com/kozhamseitova/test-task/internal/service"
	"github.com/kozhamseitova/test-task/pkg/jwttoken"
	"github.com/kozhamseitova/test-task/pkg/logger"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	app := fiber.New()

	cfg, err := config.InitConfig("../../config.yaml")
	require.NoError(t, err)

	jwtToken := jwttoken.New("ekfmkejmrj")

	logger, err := logger.NewLogger(cfg.App.Production)
	require.NoError(t, err)

	repo := fakerepo.NewRepository()
	srvc := service.NewService(repo, jwtToken, cfg, logger)
	hndlr := NewHandler(srvc, cfg, logger)

	app.Post("/api/v1/users", hndlr.createUser)

	request := &entity.CreateUserRequest{
		Username:   "testuser",
		Password:   "testpassword",
		FirstName:  "Test",
		LastName:   "Test",
		City:       "Test",
		BirthDate:  time.Now(),
	}
	
	jsonRequest, err := json.Marshal(request)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(jsonRequest))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	// Check if the id is present
	var responseBody struct {
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	require.NoError(t, err)

	// Check if the id is not empty
	require.NotZero(t, responseBody.Data.ID)

	// Check if the user created and exists
	user, err := repo.GetUsersById(context.Background(), responseBody.Data.ID)
	require.NoError(t, err)

	require.Equal(t, request.Username, user.Username)
	require.Equal(t, request.FirstName, user.FirstName)
	require.Equal(t, request.LastName, user.LastName)
	require.Equal(t, request.City, user.City)


}

func TestLogin(t *testing.T) {
	app := fiber.New()

	cfg, err := config.InitConfig("../../config.yaml")
	require.NoError(t, err)

	jwtToken := jwttoken.New("ekfmkejmrj")

	logger, err := logger.NewLogger(cfg.App.Production)
	require.NoError(t, err)

	repo := fakerepo.NewRepository()
	srvc := service.NewService(repo, jwtToken, cfg, logger)
	hndlr := NewHandler(srvc, nil, nil)

	app.Post("/api/v1/login", hndlr.login)

	// Create a user to add to the fake repository
	fakeUser := &entity.CreateUserRequest{
		Username:   "testuser",
		Password:   "testpassword",
		FirstName:  "Test",
		LastName:   "Test",
		City:       "Test",
		BirthDate:  time.Now(),
	}

	_, err = srvc.CreateUser(context.Background(), fakeUser)
	require.NoError(t, err)

	request := api.LoginRequest{
		Username: "testuser",
		Password: "testpassword",
	}

	jsonRequest, err := json.Marshal(request)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(jsonRequest))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	require.NoError(t, err)

	// Check status code 
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// Check if the token is present
	var responseBody struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	require.NoError(t, err)

	// Check if the token is not empty
	require.NotEmpty(t, responseBody.Data.Token)
}

func TestGetUsers(t *testing.T) {
	app := fiber.New()

	cfg, err := config.InitConfig("../../config.yaml")
	require.NoError(t, err)

	jwtToken := jwttoken.New("ekfmkejmrj")

	logger, err := logger.NewLogger(cfg.App.Production)
	require.NoError(t, err)

	repo := fakerepo.NewRepository()
	srvc := service.NewService(repo, jwtToken, cfg, logger)
	hndlr := NewHandler(srvc, cfg, logger)

	app.Get("/api/v1/users", hndlr.getUsers)
	app.Use(hndlr.userIdentity)
	
	// Create a user to add to the fake repository
	users := []*entity.CreateUserRequest{
		{
			Username:  "user1",
			FirstName: "John",
			LastName:  "Doe",
			City:      "New York",
		},
		{
			Username:  "user2",
			FirstName: "Jane",
			LastName:  "Smith",
			City:      "Los Angeles",
		},
	}
	for _, u := range users {
		_, err := repo.CreateUser(context.Background(), u)
		require.NoError(t, err)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users?page=1&amount=10&city=Astana", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)

	// Check if the status code is correct
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var responseBody struct {
		Data struct {
			Users []*entity.User `json:"users"`
		} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	require.NoError(t, err)

	// Check if the returned users equals the expected number of users
	require.Equal(t, len(users), len(responseBody.Data.Users))
}