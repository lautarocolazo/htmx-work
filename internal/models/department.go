package models

import "database/sql"

type Department struct {
	ID   int
	Name string
}

func (d *Department) GetAllDepartments(db *sql.DB) ([]Department, error) {
	query := "SELECT dept_id, name FROM Departments"

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []Department

	for rows.Next() {
		var department Department
		if err := rows.Scan(&department.ID, &department.Name); err != nil {
			return nil, err
		}
		departments = append(departments, department)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}

func (d *Department) GetDepartmentByID(db *sql.DB, id int) (*Department, error) {
	query := "SELECT dept_id, name FROM Departments WHERE dept_id = $1"

	var department Department
	err := db.QueryRow(query, id).Scan(&department.ID, &department.Name)
	if err != nil {
		return nil, err
	}

	return &department, nil
}

func (d *Department) CreateDepartment(db *sql.DB) error {
	query := "INSERT INTO Departments (name) VALUES ($1) RETURNING dept_id"

	err := db.QueryRow(query, d.Name).Scan(&d.ID)
	if err != nil {
		return err
	}
	return nil
}

func (d *Department) UpdateDepartment(db *sql.DB) error {
	query := "UPDATE Departments SET name = $1 WHERE dept_id = $2"

	_, err := db.Exec(query, d.Name, d.ID)
	if err != nil {
		return err
	}
	return nil
}

func (d *Department) DeleteDepartment(db *sql.DB) error {
	query := "DELETE FROM Departments WHERE dept_id = $1"

	_, err := db.Exec(query, d.ID)
	if err != nil {
		return err
	}
	return nil
}
