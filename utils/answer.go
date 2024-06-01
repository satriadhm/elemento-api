package utils

import "github.com/google/uuid"

type AnswerRequest struct {
	AnswerID       uuid.UUID
	AnswerTitle    string
	AnswerSubtitle string
	QuestionID     uuid.UUID
}

type UserAnswerRequest struct {
	QuestionID uuid.UUID
	AnswerID   uuid.UUID
}
