package integration_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/devShahriar/xm/internal/adapters/db"               // Updated to your module path
	httpAdpater "github.com/devShahriar/xm/internal/adapters/http" // Updated to your module path
	"github.com/devShahriar/xm/internal/config"
	"github.com/devShahriar/xm/internal/entity" // Updated to your module path
	"github.com/devShahriar/xm/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	// Config import for DB
)

var jwtToken string // Stores the JWT token from the login step

// Setup test server using NewDBInstance
func setupTestServer() *echo.Echo {
	// Load the config
	conf := config.GetCmdConfig()
	conf.ReadConfig()
	appConfig := config.GetAppConfig()
	appConfig.DbConfig = config.DbConfig{
		Host:               "127.0.0.1",
		Password:           "postgres",
		User:               "postgres",
		DbName:             "b1",
		Port:               "5432",
		SlowQueryThreshold: 10,
	}

	// Initialize Echo
	e := echo.New()

	// Initialize DB instance
	companyDB := db.NewDBInstance()
	if companyDB == nil {
		log.Println("Failed to connect db")
	}

	// Initialize repository, use case, and HTTP handlers
	companyRepo := companyDB
	companyUsecase := usecase.NewCompanyUsecase(companyRepo)
	companyHandler := httpAdpater.NewServer(companyUsecase)

	// Define routes
	e.POST("/v1/login", companyHandler.Login)
	e.POST("/v1/companies", companyHandler.CreateCompany)
	e.GET("/v1/companies/:id", companyHandler.GetCompanyByID)
	e.PUT("/v1/companies/:id", companyHandler.UpdateCompany)
	e.DELETE("/v1/companies/:id", companyHandler.DeleteCompany)

	return e
}

func TestLogin(t *testing.T) {
	// Setup the server and DB
	e := setupTestServer()

	// Test login endpoint to get JWT token
	req := httptest.NewRequest(http.MethodPost, "/v1/login", bytes.NewReader([]byte(`{"email": "shudip@gmail.com"}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Serve the request
	e.ServeHTTP(rec, req)

	// Assert that the status code is OK (200)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Parse the response body to check the token
	var response map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err) // Assert that there was no error parsing the response

	// Extract the token and assert it's not empty
	jwtToken, exists := response["token"].(string)
	assert.True(t, exists)       // Assert the "token" field exists
	assert.NotEmpty(t, jwtToken) // Assert that the token is not empty
}

func TestCreateCompany(t *testing.T) {
	// Setup the server and DB
	e := setupTestServer()

	// Company payload
	company := map[string]interface{}{
		"id":            "599c4b1f-8901-4d3c-bf87-795bc82b6f66",
		"name":          "Techy",
		"description":   "A leading tech company.",
		"num_employees": 100,
		"registered":    true,
		"type":          "Corporations",
	}
	body, _ := json.Marshal(company)

	// Create company request
	req := httptest.NewRequest(http.MethodPost, "/v1/companies", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", jwtToken)

	rec := httptest.NewRecorder()

	// Serve the request
	e.ServeHTTP(rec, req)

	// Assert the status code is 201 Created
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Optionally, assert the response body if necessary
}

func TestGetCompany(t *testing.T) {
	// Setup the server and DB
	e := setupTestServer()

	// Insert a company into the database
	companyDB := db.NewDBInstance() // Using the DB instance for the test
	companyDB.Db.Exec("INSERT INTO companies (id, name, description, num_employees, registered, type) VALUES ('599c4b1f-8901-4d3c-bf87-795bc82b6f64', 'TechCorp', 'A leading tech company', 100, true, 'Corporations')")
	defer companyDB.Db.Exec(`DELETE FROM companies where id = '599c4b1f-8901-4d3c-bf87-795bc82b6f64'`)
	// Fetch company by ID
	req := httptest.NewRequest(http.MethodGet, "/v1/companies/599c4b1f-8901-4d3c-bf87-795bc82b6f64", nil)
	req.Header.Set("Authorization", jwtToken)

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Serve the request
	e.ServeHTTP(rec, req)

	// Assert that the status code is OK (200)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Parse the response body
	var fetchedCompany entity.Company
	err := json.Unmarshal(rec.Body.Bytes(), &fetchedCompany)
	assert.NoError(t, err) // Ensure there's no error in parsing the response

	// Assert that the company name matches the inserted company
	assert.Equal(t, "TechCorp", fetchedCompany.Name)
}

func TestUpdateCompany(t *testing.T) {
	// Setup the server and DB
	e := setupTestServer()

	// Insert a company to be updated
	companyDB := db.NewDBInstance() // Using the DB instance for the test
	companyDB.Db.Exec("INSERT INTO companies (id, name, description, num_employees, registered, type) VALUES ('eff3cf3d-9959-40b8-9659-f56bc84d60d5', 'newTechCorp', 'A leading tech company', 100, true, 'Corporations')")
	defer companyDB.Db.Exec(`DELETE FROM companies where id = 'eff3cf3d-9959-40b8-9659-f56bc84d60d5'`)
	// Update payload
	updatePayload := map[string]interface{}{
		"name":          "TechCorpV2",
		"num_employees": 200,
	}
	body, _ := json.Marshal(updatePayload)

	// Create the update company request
	req := httptest.NewRequest(http.MethodPut, "/v1/companies/eff3cf3d-9959-40b8-9659-f56bc84d60d5", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", jwtToken)

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Serve the request
	e.ServeHTTP(rec, req)

	// Assert the status code is OK (200)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Optionally, parse the response and assert the updated data
	var updatedCompany entity.Company
	err := json.Unmarshal(rec.Body.Bytes(), &updatedCompany)
	assert.NoError(t, err)

	// Assert that the company was updated
	assert.Equal(t, "TechCorpV2", updatedCompany.Name)
	assert.Equal(t, 200, updatedCompany.NumEmployees)
}

func TestDeleteCompany(t *testing.T) {
	// Setup the server and DB
	e := setupTestServer()

	// Insert a company to be deleted
	companyDB := db.NewDBInstance() // Using the DB instance for the test
	companyDB.Db.Exec("INSERT INTO companies (id, name, description, num_employees, registered, type) VALUES ('e7b42c20-cb00-43f4-8192-2ce7413e4d65', 'deleteTechCorp', 'A leading tech company', 100, true, 'Corporations')")
	defer companyDB.Db.Exec(`DELETE FROM companies where id = 'e7b42c20-cb00-43f4-8192-2ce7413e4d65'`)
	// Create the delete request
	req := httptest.NewRequest(http.MethodDelete, "/v1/companies/e7b42c20-cb00-43f4-8192-2ce7413e4d65", nil)
	req.Header.Set("Authorization", jwtToken)

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Serve the request
	e.ServeHTTP(rec, req)

	// Assert the status code is 204 No Content
	assert.Equal(t, http.StatusNoContent, rec.Code)
}
