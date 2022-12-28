package test

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_test")
	helper.PanifIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}

func setupRouter() http.Handler {
	validate := validator.New()
	db := setupTestDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategory(t *testing.T) {
	router := setupRouter()
	requestBody := strings.NewReader(`{"name":"gadget"}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Set("X-API-Key", "12345")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)
	var responseBody map[string]interface{}
	err := json.Unmarshal(body, &responseBody)
	if err != nil {
		panic(err)
	}
	fmt.Println(responseBody)
	assert.Equal(t, float64(200), responseBody["code"], "must be 200")
	assert.Equal(t, "OK", responseBody["status"], "status be OK")
	assert.Equal(t, "gadget", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, 200, result.StatusCode)
}
