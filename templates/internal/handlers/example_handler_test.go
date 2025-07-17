package handlers

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
}