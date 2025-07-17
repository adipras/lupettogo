package generator

// Test template files

var testTemplates = map[string]string{
	"internal/handlers/example_handler_test.go": `package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"{{.ProjectName}}/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockExampleService struct {
	mock.Mock
}

func (m *MockExampleService) GetExample() map[string]interface{} {
	args := m.Called()
	return args.Get(0).(map[string]interface{})
}

func TestExampleHandler_GetExample(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create mock service
	mockService := new(MockExampleService)
	expectedResponse := map[string]interface{}{
		"message": "Hello from LupettoGo! 🐺",
		"status":  "success",
	}
	mockService.On("GetExample").Return(expectedResponse)

	// Create handler
	handler := &ExampleHandler{
		exampleService: mockService,
	}

	// Create router and register route
	router := gin.New()
	router.GET("/example", handler.GetExample)

	// Create request
	req, _ := http.NewRequest("GET", "/example", nil)
	w := httptest.NewRecorder()

	// Perform request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse["message"], response["message"])
	assert.Equal(t, expectedResponse["status"], response["status"])

	// Verify mock was called
	mockService.AssertExpectations(t)
}`,

	"internal/services/example_service_test.go": `package services

import (
	"testing"

	"{{.ProjectName}}/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockExampleRepository struct {
	mock.Mock
}

func (m *MockExampleRepository) FindAll() ([]*models.Example, error) {
	args := m.Called()
	return args.Get(0).([]*models.Example), args.Error(1)
}

func (m *MockExampleRepository) FindByID(id uint) (*models.Example, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Example), args.Error(1)
}

func (m *MockExampleRepository) Create(example *models.Example) (*models.Example, error) {
	args := m.Called(example)
	return args.Get(0).(*models.Example), args.Error(1)
}

func (m *MockExampleRepository) Update(example *models.Example) (*models.Example, error) {
	args := m.Called(example)
	return args.Get(0).(*models.Example), args.Error(1)
}

func (m *MockExampleRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestExampleService_GetExample(t *testing.T) {
	// Create mock repository
	mockRepo := new(MockExampleRepository)

	// Create service
	service := NewExampleService(mockRepo)

	// Test GetExample
	result := service.GetExample()

	// Assertions
	assert.NotNil(t, result)
	assert.Equal(t, "Hello from LupettoGo! 🐺", result["message"])
	assert.Equal(t, "success", result["status"])
}

func TestExampleService_GetAllExamples(t *testing.T) {
	// Create mock repository
	mockRepo := new(MockExampleRepository)

	// Set up mock expectations
	expectedExamples := []*models.Example{
		{ID: 1, Name: "Test 1", Email: "test1@example.com"},
		{ID: 2, Name: "Test 2", Email: "test2@example.com"},
	}
	mockRepo.On("FindAll").Return(expectedExamples, nil)

	// Create service
	service := NewExampleService(mockRepo)

	// Test GetAllExamples
	result, err := service.GetAllExamples()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedExamples, result)
	mockRepo.AssertExpectations(t)
}`,
}