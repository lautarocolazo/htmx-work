package handlers

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"htmx-fiber/cmd/web"
	"htmx-fiber/internal/database"
	"htmx-fiber/internal/models"
)

/********** Handlers for Deparment Views **********/

// Render List Page with success/error messages
func HandleViewDepartmentList(c *fiber.Ctx, db *sql.DB) error {

	department := new(models.Department)

	deparmentsSlice, err := department.GetAllDepartments(db)
	if err != nil {
		println("Error in the HandleViewList")
	}

	deparmentsIndex := web.DepartmentIndex(deparmentsSlice)

	handler := adaptor.HTTPHandler(templ.Handler(deparmentsIndex))

	return handler(c)
}

// Render Create Todo Page with success/error messages
func HandleViewCreatePage(c *fiber.Ctx, db *sql.DB) error {

	if c.Method() == "POST" {

		newDepartment := new(models.Department)
		println(newDepartment.Name)
		newDepartment.Name = strings.Trim(c.FormValue("name"), " ")
		println(newDepartment.Name)
		err := newDepartment.CreateDepartment(db)
		if err != nil {
			fmt.Println("Error creating the department:", err)
			return err
		}

		return c.Redirect("/department/list")
	}

	createDepartmentIndex := web.CreateDepartment()

	handler := adaptor.HTTPHandler(templ.Handler(createDepartmentIndex))

	return handler(c)
}

func HandleViewDeleteDepartment(c *fiber.Ctx, db *sql.DB) error {

	if c.Method() == "DELETE" {

		id := c.Params("id")
		departmentId, err := strconv.Atoi(id)
		fmt.Println("The department ID was received:", departmentId)
		if err != nil {
			return err
		}

		department := &models.Department{ID: departmentId}
		err = department.DeleteDepartment(database.DB)

		departmentNew := new(models.Department)

		deparmentsSlice, err := departmentNew.GetAllDepartments(db)
		if err != nil {
			println("Error in the HandleViewList")
		}

		deparmentsIndex := web.DepartmentIndex(deparmentsSlice)

		handler := adaptor.HTTPHandler(templ.Handler(deparmentsIndex))

		return handler(c)

	}

	createDepartmentIndex := web.CreateDepartment()

	handler := adaptor.HTTPHandler(templ.Handler(createDepartmentIndex))

	return handler(c)

}
