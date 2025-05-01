package requests

type FarmRequest struct {
	Name   string `json:"name" validate:"required"`
	UserID uint   `json:"user_id"`
}
