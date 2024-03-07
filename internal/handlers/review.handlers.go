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
	"htmx-fiber/internal/models"
)

// Render Create Todo Page with success/error messages
func HandleViewCreateReview(c *fiber.Ctx, db *sql.DB) error {

	if c.Method() == "POST" {

		hardcodedUserID := 1
		departmentID, err := strconv.Atoi(c.FormValue("department-id"))
		fmt.Println("The department ID was received:", departmentID)
		if err != nil {
			return err
		}
		positionName := strings.Trim(c.FormValue("position-name"), " ")
		fmt.Println("The position name was received:", positionName)
		rating, err := strconv.Atoi(c.FormValue("rating"))
		fmt.Println("The rating was received:", rating)
		if err != nil {
			return err
		}
		baseSalary, err := strconv.ParseFloat(c.FormValue("base-salary"), 64)
		fmt.Println("The base salary was received:", baseSalary)
		if err != nil {
			return err
		}
		raisesOpportunity := c.FormValue("raise-opportunity")
		fmt.Println("The raise opportunity was received:", raisesOpportunity)
		raiseExplanation := c.FormValue("raise-explanation")
		fmt.Println("The raise explanation was received:", raiseExplanation)
		reviewText := strings.Trim(c.FormValue("review"), " ")
		fmt.Println("The review was received:", reviewText)

		newJobPosition := &models.JobPosition{
			Title:        positionName,
			DepartmentID: departmentID,
			BaseSalary:   baseSalary,
		}

		newJobPositionID, err := models.CreateJobPosition(db, newJobPosition)
		if err != nil {
			fmt.Println("The new job position was failed", newJobPositionID)
			return err
		}
		fmt.Println("The new job position was created", newJobPositionID)

		newReview := &models.Review{
			JobID:             newJobPositionID,
			UserID:            hardcodedUserID,
			Rating:            rating,
			TextReview:        reviewText,
			RaisesOpportunity: raisesOpportunity == "on",
			RaiseExplanation:  raiseExplanation,
		}
		_, err = models.CreateReview(db, newReview)
		if err != nil {
			fmt.Println("The new review was failed", newJobPositionID)
			return err
		}

		fmt.Println("If you got to here, everything was submitted!")

		return c.Redirect("/review/success")
	}

	department := new(models.Department)

	deparmentsSlice, err := department.GetAllDepartments(db)
	if err != nil {
		println("Error in the HandleViewCreateReview")
	}

	createDepartmentIndex := web.CreateReviewPage(deparmentsSlice)

	handler := adaptor.HTTPHandler(templ.Handler(createDepartmentIndex))

	return handler(c)
}

// Handler
func HandleViewDepartmentDetail(c *fiber.Ctx, db *sql.DB) error {
	// Get the department ID from the route parameter
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	newDepartment := new(models.Department)
	// Get the department from the database
	department, err := newDepartment.GetDepartmentByID(db, id)
	if err != nil {
		return err
	}

	// Get the job positions for this department from the database
	jobPositions, err := models.GetJobPositionsByDepartmentID(db, id)
	if err != nil {
		return err
	}

	// Render the department detail page
	departmentDetail := web.DepartmentDetail(*department, jobPositions)
	handler := adaptor.HTTPHandler(templ.Handler(departmentDetail))
	return handler(c)
}
