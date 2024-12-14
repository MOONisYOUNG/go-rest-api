package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student
var lastId int

func MakeWebHandler() *fiber.App {
	app := fiber.New()

	app.Get("/students", GetStudentListHandler)
	app.Get("/students/:id", GetStudentHandler)
	app.Post("/students", PostStudentHandler)
	app.Delete("/students/:id", DeleteStudentHandler)

	students = make(map[int]Student)
	students[1] = Student{1, "Teplotaxl", 999, 100}
	students[2] = Student{2, "Robert", 10, 70}
	lastId = 2

	return app
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func GetStudentListHandler(c *fiber.Ctx) error {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}
	sort.Sort(list)
	return c.Status(fiber.StatusOK).JSON(list)
}

func GetStudentHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	student, ok := students[id]
	if !ok {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Status(fiber.StatusOK).JSON(student)
}

func PostStudentHandler(c *fiber.Ctx) error {
	var student Student
	if err := c.BodyParser(&student); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	lastId++
	student.Id = lastId
	students[lastId] = student
	return c.SendStatus(fiber.StatusCreated)
}

func DeleteStudentHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	_, ok := students[id]
	if !ok {
		return c.SendStatus(fiber.StatusNotFound)
	}
	delete(students, id)
	return c.SendStatus(fiber.StatusOK)
}

func main() {
	fmt.Println("Starting server on :3000...")
	app := MakeWebHandler()
	if err := app.Listen(":3000"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
