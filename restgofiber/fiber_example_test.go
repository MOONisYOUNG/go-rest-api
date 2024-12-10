package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	app := MakeWebHandler()

	req, _ := http.NewRequest("GET", "/students", nil)
	res, err := app.Test(req, -1) // Timeout을 -1로 설정해 무제한 대기
	assert.Nil(err)
	assert.Equal(200, res.StatusCode)

	var list []Student
	err = json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(2, len(list))
	assert.Equal("Teplotaxl", list[0].Name)
	assert.Equal("Robert", list[1].Name)
}

func TestJsonHandler2(t *testing.T) {
	assert := assert.New(t)

	var student Student
	app := MakeWebHandler()

	req, _ := http.NewRequest("GET", "/students/1", nil)
	res, err := app.Test(req, -1)
	assert.Nil(err)
	assert.Equal(200, res.StatusCode)

	err = json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("Teplotaxl", student.Name)

	req, _ = http.NewRequest("GET", "/students/2", nil)
	res, err = app.Test(req, -1)
	assert.Nil(err)
	assert.Equal(200, res.StatusCode)

	err = json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("Robert", student.Name)
}

func TestJsonHandler3(t *testing.T) {
	assert := assert.New(t)

	var student Student
	app := MakeWebHandler()

	body := `{"ID":0, "Name":"Fibonacci", "Age":855, "Score":100}`
	req, _ := http.NewRequest("POST", "/students", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	res, err := app.Test(req, -1)
	assert.Nil(err)
	assert.Equal(201, res.StatusCode)

	req, _ = http.NewRequest("GET", "/students/3", nil)
	res, err = app.Test(req, -1)
	assert.Nil(err)
	assert.Equal(200, res.StatusCode)

	err = json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("Fibonacci", student.Name)
}

func TestJsonHandler4(t *testing.T) {
	assert := assert.New(t)

	app := MakeWebHandler()

	req, _ := http.NewRequest("DELETE", "/students/1", nil)
	res, err := app.Test(req, -1)
	assert.Nil(err)
	assert.Equal(200, res.StatusCode)

	req, _ = http.NewRequest("GET", "/students", nil)
	res, err = app.Test(req, -1)
	assert.Nil(err)
	assert.Equal(200, res.StatusCode)

	var list []Student
	err = json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(1, len(list))
	assert.Equal("Robert", list[0].Name)
}
