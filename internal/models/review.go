package models

import (
	"database/sql"
)

type Review struct {
	ReviewID          int
	JobID             int
	UserID            int
	Rating            int
	TextReview        string
	RaisesOpportunity bool
	RaiseExplanation  string
}

// GetAllReviews retrieves all reviews from the database.
func GetAllReviews(db *sql.DB) ([]Review, error) {
	query := "SELECT review_id, job_id, user_id, rating, text_review, raises_opportunity, raise_explanation FROM reviews"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var review Review
		err := rows.Scan(
			&review.ReviewID,
			&review.JobID,
			&review.UserID,
			&review.Rating,
			&review.TextReview,
			&review.RaisesOpportunity,
			&review.RaiseExplanation,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

// GetReviewByID retrieves a review by its ID from the database.
func GetReviewByID(db *sql.DB, id int) (*Review, error) {
	query := "SELECT review_id, job_id, user_id, rating, text_review, raises_opportunity, raise_explanation FROM reviews WHERE review_id = $1"
	row := db.QueryRow(query, id)
	var review Review
	err := row.Scan(
		&review.ReviewID,
		&review.JobID,
		&review.UserID,
		&review.Rating,
		&review.TextReview,
		&review.RaisesOpportunity,
		&review.RaiseExplanation,
	)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

// CreateReview creates a new review in the database.
func CreateReview(db *sql.DB, review *Review) (int, error) {
	query := `INSERT INTO reviews (job_id, user_id, rating, text_review, raises_opportunity, raise_explanation) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING review_id`
	var id int
	err := db.QueryRow(
		query,
		review.JobID,
		review.UserID,
		review.Rating,
		review.TextReview,
		review.RaisesOpportunity,
		review.RaiseExplanation,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateReview updates an existing review in the database.
func UpdateReview(db *sql.DB, review *Review) error {
	query := "UPDATE reviews SET job_id = $1, user_id = $2, rating = $3, text_review = $4, raises_opportunity = $5, raise_explanation = $6 WHERE review_id = $7"
	_, err := db.Exec(
		query,
		review.JobID,
		review.UserID,
		review.Rating,
		review.TextReview,
		review.RaisesOpportunity,
		review.RaiseExplanation,
		review.ReviewID,
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteReview deletes a review from the database.
func DeleteReview(db *sql.DB, id int) error {
	query := "DELETE FROM reviews WHERE review_id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func GetReviewsByJobPositionID(db *sql.DB, jobPositionID int) ([]Review, error) {
	query := "SELECT review_id, job_id, user_id, rating, text_review, raises_opportunity, raise_explanation FROM reviews WHERE job_id = $1"
	rows, err := db.Query(query, jobPositionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var review Review
		err := rows.Scan(
			&review.ReviewID,
			&review.JobID,
			&review.UserID,
			&review.Rating,
			&review.TextReview,
			&review.RaisesOpportunity,
			&review.RaiseExplanation,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}
