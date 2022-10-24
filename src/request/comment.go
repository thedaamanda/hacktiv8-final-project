package request

type (
	CreateCommentRequest struct {
		Message string `json:"message" validate:"required"`
		PhotoID int    `json:"photo_id" validate:"required"`
	}

	UpdateCommentRequest struct {
		Message string `json:"message" validate:"required"`
	}
)
