package taskmanagement

import "errors"

type Tasks struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description"  db:"description"`
	Status      string `json:"status" db:"status"`
	Created     string `json:"created_at" db:"created_at"`
	Updated     string `json:"updated_at" db:"updated_at"`
}

type UpdateTaskInput struct {
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description"  db:"description"`
	Status      *string `json:"status" db:"status"`
}

func (u *UpdateTaskInput) Validate() error {
	if u.Title == nil && u.Description == nil && u.Status == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
