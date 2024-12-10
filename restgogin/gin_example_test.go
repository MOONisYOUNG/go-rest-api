package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	app := gin.Default()
	SetupHandlers(app)

	req, _ := http.NewRequest("GET", "/students", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(200, w.Code)

	var list []Student
	err := json.NewDecoder(w.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(2, len(list))
	assert.Equal("Teplotaxl", list[0].Name)
	assert.Equal("Robert", list[1].Name)
}

func TestJsonHandler2(t *testing.T) {
	assert := assert.New(t)

	var student Student
	app := gin.Default()
	SetupHandlers(app)

	req, _ := http.NewRequest("GET", "/students/1", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	assert.Equal(200, w.Code)

	err := json.NewDecoder(w.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("Teplotaxl", student.Name)

	req, _ = http.NewRequest("GET", "/students/2", nil)
	w = httptest.NewRecorder()
	app.ServeHTTP(w, req)
	assert.Equal(200, w.Code)

	err = json.NewDecoder(w.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("Robert", student.Name)
}

func TestJsonHandler3(t *testing.T) {
	assert := assert.New(t)

	var student Student
	app := gin.Default()
	SetupHandlers(app)

	body := `{"ID":0, "Name":"Fibonacci", "Age":855, "Score":100}`
	req, _ := http.NewRequest("POST", "/students", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(201, w.Code)

	req, _ = http.NewRequest("GET", "/students/3", nil)
	w = httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(200, w.Code)

	err := json.NewDecoder(w.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("Fibonacci", student.Name)
}

func TestJsonHandler4(t *testing.T) {
	assert := assert.New(t)

	app := gin.Default()
	SetupHandlers(app)

	req, _ := http.NewRequest("DELETE", "/students/1", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	assert.Equal(200, w.Code)

	req, _ = http.NewRequest("GET", "/students", nil)
	w = httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(200, w.Code)

	var list []Student
	err := json.NewDecoder(w.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(1, len(list))
	assert.Equal("Robert", list[0].Name)
}
