package models

import (
	"database/sql"
)

type JobPosition struct {
	JobID        int
	Title        string
	DepartmentID int
	BaseSalary   float64
}

// GetAllJobPositions retrieves all job positions from the database.
func GetAllJobPositions(db *sql.DB) ([]JobPosition, error) {
	query := "SELECT job_id, title, dept_id, base_salary FROM jobpositions"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var positions []JobPosition
	for rows.Next() {
		var position JobPosition
		err := rows.Scan(
			&position.JobID,
			&position.Title,
			&position.DepartmentID,
			&position.BaseSalary,
		)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	return positions, nil
}

// GetJobPositionByID retrieves a job position by its ID from the database.
func GetJobPositionByID(db *sql.DB, id int) (*JobPosition, error) {
	query := "SELECT job_id, title, dept_id, base_salary FROM jobpositions WHERE job_id = $1"
	row := db.QueryRow(query, id)
	var position JobPosition
	err := row.Scan(
		&position.JobID,
		&position.Title,
		&position.DepartmentID,
		&position.BaseSalary,
	)
	if err != nil {
		return nil, err
	}
	return &position, nil
}

// CreateJobPosition creates a new job position in the database.
func CreateJobPosition(db *sql.DB, position *JobPosition) (int, error) {
	query := `INSERT INTO jobpositions (title, dept_id, base_salary) 
	          VALUES ($1, $2, $3) RETURNING job_id`
	var id int
	err := db.QueryRow(
		query,
		position.Title,
		position.DepartmentID,
		position.BaseSalary,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateJobPosition updates an existing job position in the database.
func UpdateJobPosition(db *sql.DB, position *JobPosition) error {
	query := "UPDATE jobpositions SET title = $1, dept_id = $2, base_salary = $3 WHERE job_id = $4"
	_, err := db.Exec(
		query,
		position.Title,
		position.DepartmentID,
		position.BaseSalary,
		position.JobID,
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteJobPosition deletes a job position from the database.
func DeleteJobPosition(db *sql.DB, id int) error {
	query := "DELETE FROM jobpositions WHERE job_id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

// GetJobPositionsByDepartmentID retrieves all job positions for a specific department from the database.
func GetJobPositionsByDepartmentID(db *sql.DB, departmentID int) ([]JobPosition, error) {
	query := "SELECT job_id, title, dept_id, base_salary FROM jobpositions WHERE dept_id = $1"
	rows, err := db.Query(query, departmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var positions []JobPosition
	for rows.Next() {
		var position JobPosition
		err := rows.Scan(
			&position.JobID,
			&position.Title,
			&position.DepartmentID,
			&position.BaseSalary,
		)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	return positions, nil
}
