package models

import (
	"database/sql"
)

type ReviewFeedback struct {
	FeedbackID int
	ReviewID   int
	UserID     int
	IsUseful   bool
}

// GetAllReviewFeedbacks retrieves all review feedbacks from the database.
func GetAllReviewFeedbacks(db *sql.DB) ([]ReviewFeedback, error) {
	query := "SELECT feedback_id, review_id, user_id, is_useful FROM review_feedbacks"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feedbacks []ReviewFeedback
	for rows.Next() {
		var feedback ReviewFeedback
		err := rows.Scan(
			&feedback.FeedbackID,
			&feedback.ReviewID,
			&feedback.UserID,
			&feedback.IsUseful,
		)
		if err != nil {
			return nil, err
		}
		feedbacks = append(feedbacks, feedback)
	}
	return feedbacks, nil
}

// GetReviewFeedbackByID retrieves a review feedback by its ID from the database.
func GetReviewFeedbackByID(db *sql.DB, id int) (*ReviewFeedback, error) {
	query := "SELECT feedback_id, review_id, user_id, is_useful FROM review_feedbacks WHERE feedback_id = ?"
	row := db.QueryRow(query, id)
	var feedback ReviewFeedback
	err := row.Scan(&feedback.FeedbackID, &feedback.ReviewID, &feedback.UserID, &feedback.IsUseful)
	if err != nil {
		return nil, err
	}
	return &feedback, nil
}

// CreateReviewFeedback creates a new review feedback in the database.
func CreateReviewFeedback(db *sql.DB, feedback *ReviewFeedback) (int, error) {
	query := "INSERT INTO review_feedbacks (review_id, user_id, is_useful) VALUES (?, ?, ?)"
	result, err := db.Exec(query, feedback.ReviewID, feedback.UserID, feedback.IsUseful)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// UpdateReviewFeedback updates an existing review feedback in the database.
func UpdateReviewFeedback(db *sql.DB, feedback *ReviewFeedback) error {
	query := "UPDATE review_feedbacks SET review_id = ?, user_id = ?, is_useful = ? WHERE feedback_id = ?"
	_, err := db.Exec(
		query,
		feedback.ReviewID,
		feedback.UserID,
		feedback.IsUseful,
		feedback.FeedbackID,
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteReviewFeedback deletes a review feedback from the database.
func DeleteReviewFeedback(db *sql.DB, id int) error {
	query := "DELETE FROM review_feedbacks WHERE feedback_id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
