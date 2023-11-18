package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Jehanv60/app"
	"github.com/Jehanv60/controller"
	"github.com/Jehanv60/middleware"
	"github.com/Jehanv60/model/domain"
	"github.com/Jehanv60/repository"
	"github.com/Jehanv60/service"
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func TestDb(t *testing.T) {
	app.NewDb()
}

func truncateDB(db *sql.DB) {
	db.Exec("truncate barang RESTART IDENTITY")
}

func NewDBTest() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost user=han port=5432 password=solo dbname=pos1_test sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
func setupRouter(db *sql.DB) http.Handler {
	DB := app.NewDb()
	validate := validator.New()
	barangRepository := repository.NewRepositoryBarang()
	barangService := service.NewBarangService(barangRepository, DB, validate)
	barangController := controller.NewBarangController(barangService)
	penggunaRepository := repository.NewRepositoryPengguna()
	penggunaService := service.NewPenggunaService(penggunaRepository, DB, validate)
	penggunaController := controller.NewPenggunaController(penggunaService)
	router := app.NewRouter(barangController, penggunaController)
	return middleware.NewAuthMiddleware(router)
}

func TestCreateSucces(t *testing.T) {
	db := NewDBTest()
	truncateDB(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"nameprod": "Surya", "hargaprod": 1000,"keterangan": "test"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/barang", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "Rahasia")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64))) //alternatif convert 200 menjadi float64(200)
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, "Surya", responseBody["data"].(map[string]interface{})["nameprod"])
	assert.Equal(t, 1000, int(responseBody["data"].(map[string]interface{})["hargaprod"].(float64)))
	assert.Equal(t, "test", responseBody["data"].(map[string]interface{})["keterangan"])
}

func TestCreateFailed(t *testing.T) {
	db := NewDBTest()
	truncateDB(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"nameprod": "", "hargaprod": 0,"keterangan": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/barang", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "Rahasia")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64))) //alternatif convert 200 menjadi float64(200)
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestUpdateSucces(t *testing.T) {
	db := NewDBTest()
	truncateDB(db)

	tx, _ := db.Begin()
	barangRepository := repository.NewRepositoryBarang()
	barang := barangRepository.Save(context.Background(), tx, domain.Barang{
		NameProd:   "lala",
		Hargaprod:  10,
		Keterangan: "gg",
	})
	tx.Commit()
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"nameprod": "hehe", "hargaprod": 2000,"keterangan": "rokok"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/barang/"+strconv.Itoa(barang.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "Rahasia")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64))) //alternatif convert 200 menjadi float64(200)
	assert.Equal(t, "Ok", responseBody["status"])
	//klu udemy code menjadi assert.Equal(t, barang.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, float64(barang.Id), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, "hehe", responseBody["data"].(map[string]interface{})["nameprod"])
	assert.Equal(t, 2000, int(responseBody["data"].(map[string]interface{})["hargaprod"].(float64)))
	assert.Equal(t, "rokok", responseBody["data"].(map[string]interface{})["keterangan"])
}

func TestUpdateFailed(t *testing.T) {
	db := NewDBTest()
	truncateDB(db)

	tx, _ := db.Begin()
	barangRepository := repository.NewRepositoryBarang()
	barang := barangRepository.Save(context.Background(), tx, domain.Barang{
		NameProd:   "lala",
		Hargaprod:  10,
		Keterangan: "gg",
	})
	tx.Commit()
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"nameprod": "", "hargaprod": 0,"keterangan": ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/barang/"+strconv.Itoa(barang.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "Rahasia")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64))) //alternatif convert 200 menjadi float64(200)
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestGetBarangSucces(t *testing.T) {
	db := NewDBTest()
	truncateDB(db)

	tx, _ := db.Begin()
	barangRepository := repository.NewRepositoryBarang()
	barang := barangRepository.Save(context.Background(), tx, domain.Barang{
		NameProd:   "lala",
		Hargaprod:  10,
		Keterangan: "gg",
	})
	tx.Commit()
	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/barang/"+strconv.Itoa(barang.Id), nil)
	request.Header.Add("X-API-Key", "Rahasia")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64))) //alternatif convert 200 menjadi float64(200)
	assert.Equal(t, "Ok", responseBody["status"])
	//klu udemy code menjadi assert.Equal(t, barang.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, float64(barang.Id), responseBody["data"].(map[string]interface{})["id"])
	assert.Equal(t, barang.NameProd, responseBody["data"].(map[string]interface{})["nameprod"])
	assert.Equal(t, barang.Hargaprod, int(responseBody["data"].(map[string]interface{})["hargaprod"].(float64)))
	assert.Equal(t, barang.Keterangan, responseBody["data"].(map[string]interface{})["keterangan"])
}

func TestGetBarangFailed(t *testing.T) {
	db := NewDBTest()
	truncateDB(db)

	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/barang/404", nil)
	request.Header.Add("X-API-Key", "Rahasia")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64))) //alternatif convert 200 menjadi float64(200)
	assert.Equal(t, "Data Tidak Ditemukan", responseBody["status"])
}

func TestDeleteSucces(t *testing.T) {
	db := NewDBTest()
	truncateDB(db)

	tx, _ := db.Begin()
	barangRepository := repository.NewRepositoryBarang()
	barang := barangRepository.Save(context.Background(), tx, domain.Barang{
		NameProd:   "lala",
		Hargaprod:  10,
		Keterangan: "gg",
	})
	tx.Commit()
	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/barang/"+strconv.Itoa(barang.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "Rahasia")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64))) //alternatif convert 200 menjadi float64(200)
	assert.Equal(t, "Ok", responseBody["status"])
}
func TestDeleteFailed(t *testing.T) {
	db := NewDBTest()
	truncateDB(db)

	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/barang/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "Rahasia")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64))) //alternatif convert 200 menjadi float64(200)
	assert.Equal(t, "Data Tidak Ditemukan", responseBody["status"])
}
func TestListBarangSucces(t *testing.T) {
	db := NewDBTest()
	truncateDB(db)

	tx, _ := db.Begin()
	barangRepository := repository.NewRepositoryBarang()
	barang := barangRepository.Save(context.Background(), tx, domain.Barang{
		NameProd:   "lala",
		Hargaprod:  10,
		Keterangan: "gg",
	})
	barang1 := barangRepository.Save(context.Background(), tx, domain.Barang{
		NameProd:   "lala1",
		Hargaprod:  101,
		Keterangan: "gg1",
	})
	tx.Commit()
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/barang", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "Rahasia")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	var barangs = responseBody["data"].([]interface{})
	barangResponse := barangs[0].(map[string]interface{})
	fmt.Println(barangResponse)
	barangResponse1 := barangs[1].(map[string]interface{})
	fmt.Println(barangResponse1)
	assert.Equal(t, float64(barang.Id), barangResponse["id"])
	assert.Equal(t, barang.NameProd, barangResponse["nameprod"])
	assert.Equal(t, float64(barang1.Id), barangResponse1["id"])
	assert.Equal(t, barang1.NameProd, barangResponse1["nameprod"])
}
func TestUnauthorized(t *testing.T) {
	db := NewDBTest()
	truncateDB(db)
	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/barang", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "salah")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	assert.Equal(t, 401, int(responseBody["code"].(float64))) //alternatif convert 200 menjadi float64(200)
	assert.Equal(t, "Unauthorized", responseBody["status"])
}
