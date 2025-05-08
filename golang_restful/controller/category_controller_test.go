package controller

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"golang_restful/exception"
	"golang_restful/model/web"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// MockCategoryService is a mock implementation of service.CategoryService
type MockCategoryService struct {
	// Mock functions to be set in tests
	CreateFunc   func(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	UpdateFunc   func(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	DeleteFunc   func(ctx context.Context, categoryId int)
	FindByIdFunc func(ctx context.Context, categoryId int) web.CategoryResponse
	FindAllFunc  func(ctx context.Context) []web.CategoryResponse
}

// Ensure MockCategoryService implements service.CategoryService
func (m *MockCategoryService) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	return m.CreateFunc(ctx, request)
}

func (m *MockCategoryService) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	return m.UpdateFunc(ctx, request)
}

func (m *MockCategoryService) Delete(ctx context.Context, categoryId int) {
	m.DeleteFunc(ctx, categoryId)
}

func (m *MockCategoryService) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	return m.FindByIdFunc(ctx, categoryId)
}

func (m *MockCategoryService) FindAll(ctx context.Context) []web.CategoryResponse {
	return m.FindAllFunc(ctx)
}

func TestCategoryControllerImpl_Create(t *testing.T) {
	// Create mock service
	mockService := &MockCategoryService{
		CreateFunc: func(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
			return web.CategoryResponse{
				Id:   1,
				Name: request.Name,
			}
		},
	}

	// Create controller with mock service
	controller := NewCategoryController(mockService)

	// Create test request
	requestBody := strings.NewReader(`{"name":"Test Category"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Call the controller method
	controller.Create(recorder, request, httprouter.Params{})

	// Check response
	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Decode response body
	var webResponse web.WebResponse
	err := json.NewDecoder(response.Body).Decode(&webResponse)
	assert.NoError(t, err)

	// Check response data
	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, "Ok", webResponse.Status)

	// Convert data to map to access fields
	dataMap := webResponse.Data.(map[string]interface{})
	assert.Equal(t, float64(1), dataMap["Id"]) // JSON numbers are float64
	assert.Equal(t, "Test Category", dataMap["name"])
}

func TestCategoryControllerImpl_Update(t *testing.T) {
	// Create mock service
	mockService := &MockCategoryService{
		UpdateFunc: func(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
			return web.CategoryResponse{
				Id:   request.Id,
				Name: request.Name,
			}
		},
	}

	// Create controller with mock service
	controller := NewCategoryController(mockService)

	// Create test request
	requestBody := strings.NewReader(`{"name":"Updated Category"}`)
	request := httptest.NewRequest(http.MethodPut, "/api/categories/1", requestBody)
	request.Header.Add("Content-Type", "application/json")

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Call the controller method with params
	params := httprouter.Params{
		httprouter.Param{
			Key:   "categoryId",
			Value: "1",
		},
	}
	controller.Update(recorder, request, params)

	// Check response
	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Decode response body
	var webResponse web.WebResponse
	err := json.NewDecoder(response.Body).Decode(&webResponse)
	assert.NoError(t, err)

	// Check response data
	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, "Ok", webResponse.Status)

	// Convert data to map to access fields
	dataMap := webResponse.Data.(map[string]interface{})
	assert.Equal(t, float64(1), dataMap["Id"]) // JSON numbers are float64
	assert.Equal(t, "Updated Category", dataMap["name"])
}

func TestCategoryControllerImpl_Delete(t *testing.T) {
	// Track if Delete was called
	deleteWasCalled := false

	// Create mock service
	mockService := &MockCategoryService{
		DeleteFunc: func(ctx context.Context, categoryId int) {
			deleteWasCalled = true
			assert.Equal(t, 1, categoryId)
		},
	}

	// Create controller with mock service
	controller := NewCategoryController(mockService)

	// Create test request
	request := httptest.NewRequest(http.MethodDelete, "/api/categories/1", nil)

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Call the controller method with params
	params := httprouter.Params{
		httprouter.Param{
			Key:   "categoryId",
			Value: "1",
		},
	}
	controller.Delete(recorder, request, params)

	// Check if service method was called
	assert.True(t, deleteWasCalled)

	// Check response
	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Decode response body
	var webResponse web.WebResponse
	err := json.NewDecoder(response.Body).Decode(&webResponse)
	assert.NoError(t, err)

	// Check response data
	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, "Ok", webResponse.Status)
	assert.Nil(t, webResponse.Data)
}

func TestCategoryControllerImpl_FindById(t *testing.T) {
	// Create mock service
	mockService := &MockCategoryService{
		FindByIdFunc: func(ctx context.Context, categoryId int) web.CategoryResponse {
			return web.CategoryResponse{
				Id:   categoryId,
				Name: "Test Category",
			}
		},
	}

	// Create controller with mock service
	controller := NewCategoryController(mockService)

	// Create test request
	request := httptest.NewRequest(http.MethodGet, "/api/categories/1", nil)

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Call the controller method with params
	params := httprouter.Params{
		httprouter.Param{
			Key:   "categoryId",
			Value: "1",
		},
	}
	controller.FindById(recorder, request, params)

	// Check response
	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Decode response body
	var webResponse web.WebResponse
	err := json.NewDecoder(response.Body).Decode(&webResponse)
	assert.NoError(t, err)

	// Check response data
	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, "Ok", webResponse.Status)

	// Convert data to map to access fields
	dataMap := webResponse.Data.(map[string]interface{})
	assert.Equal(t, float64(1), dataMap["Id"]) // JSON numbers are float64
	assert.Equal(t, "Test Category", dataMap["name"])
}

func TestCategoryControllerImpl_FindAll(t *testing.T) {
	// Create mock service
	mockService := &MockCategoryService{
		FindAllFunc: func(ctx context.Context) []web.CategoryResponse {
			return []web.CategoryResponse{
				{
					Id:   1,
					Name: "Category 1",
				},
				{
					Id:   2,
					Name: "Category 2",
				},
			}
		},
	}

	// Create controller with mock service
	controller := NewCategoryController(mockService)

	// Create test request
	request := httptest.NewRequest(http.MethodGet, "/api/categories", nil)

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Call the controller method
	controller.FindAll(recorder, request, httprouter.Params{})

	// Check response
	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Decode response body
	var webResponse web.WebResponse
	err := json.NewDecoder(response.Body).Decode(&webResponse)
	assert.NoError(t, err)

	// Check response data
	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, "Ok", webResponse.Status)

	// Convert data to slice to access elements
	dataSlice := webResponse.Data.([]interface{})
	assert.Equal(t, 2, len(dataSlice))

	// Check first category
	category1 := dataSlice[0].(map[string]interface{})
	assert.Equal(t, float64(1), category1["Id"]) // JSON numbers are float64
	assert.Equal(t, "Category 1", category1["name"])

	// Check second category
	category2 := dataSlice[1].(map[string]interface{})
	assert.Equal(t, float64(2), category2["Id"]) // JSON numbers are float64
	assert.Equal(t, "Category 2", category2["name"])
}

// Test error cases

func TestCategoryControllerImpl_Create_InvalidJSON(t *testing.T) {
	// Create mock service
	mockService := &MockCategoryService{
		CreateFunc: func(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
			return web.CategoryResponse{
				Id:   1,
				Name: request.Name,
			}
		},
	}

	// Create controller with mock service
	controller := NewCategoryController(mockService)

	// Create test request with invalid JSON
	requestBody := strings.NewReader(`{"name":Test Category"}`) // Missing quote
	request := httptest.NewRequest(http.MethodPost, "/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")

	// Create response recorder
	recorder := httptest.NewRecorder()

	// We expect this to panic, so we need to recover
	defer func() {
		err := recover()
		assert.NotNil(t, err, "Expected panic but none occurred")
	}()

	// Call the controller method
	controller.Create(recorder, request, httprouter.Params{})

	// If we get here, the test failed
	t.Fail()
}

func TestCategoryControllerImpl_Update_InvalidCategoryId(t *testing.T) {
	// Create mock service
	mockService := &MockCategoryService{
		UpdateFunc: func(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
			return web.CategoryResponse{
				Id:   request.Id,
				Name: request.Name,
			}
		},
	}

	// Create controller with mock service
	controller := NewCategoryController(mockService)

	// Create test request
	requestBody := strings.NewReader(`{"name":"Updated Category"}`)
	request := httptest.NewRequest(http.MethodPut, "/api/categories/invalid", requestBody)
	request.Header.Add("Content-Type", "application/json")

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Call the controller method with invalid categoryId
	params := httprouter.Params{
		httprouter.Param{
			Key:   "categoryId",
			Value: "invalid", // Not a number
		},
	}

	// We expect this to panic, so we need to recover
	defer func() {
		err := recover()
		assert.NotNil(t, err, "Expected panic but none occurred")
	}()

	// Call the controller method
	controller.Update(recorder, request, params)

	// If we get here, the test failed
	t.Fail()
}

func TestCategoryControllerImpl_FindById_NotFound(t *testing.T) {
	// Create mock service that simulates a not found error
	mockService := &MockCategoryService{
		FindByIdFunc: func(ctx context.Context, categoryId int) web.CategoryResponse {
			// Simulate not found by panicking with NotFoundError
			panic(exception.NewNotFoundError("Category not found"))
		},
	}

	// Create controller with mock service
	controller := NewCategoryController(mockService)

	// Create test request
	request := httptest.NewRequest(http.MethodGet, "/api/categories/999", nil)

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Call the controller method with params
	params := httprouter.Params{
		httprouter.Param{
			Key:   "categoryId",
			Value: "999",
		},
	}

	// We expect this to panic, so we need to recover
	defer func() {
		err := recover()
		assert.NotNil(t, err, "Expected panic but none occurred")

		// Check that it's the right type of error
		_, ok := err.(exception.NotFoundError)
		assert.True(t, ok, "Expected NotFoundError but got different error type")
	}()

	// Call the controller method
	controller.FindById(recorder, request, params)

	// If we get here, the test failed
	t.Fail()
}
