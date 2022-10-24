package response

type CommentResponse struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	PhotoID   int    `json:"photo_id"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type CommentsResponse struct {
	ID        int          `json:"id"`
	Message   string       `json:"message"`
	PhotoID   int          `json:"photo_id"`
	UserID    int          `json:"user_id"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
	User      CommentUser  `json:"User"`
	Photo     CommentPhoto `json:"Photo"`
}

type CommentUser struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CommentPhoto struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}

type UpdateCommentResponse struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	PhotoID   int    `json:"photo_id"`
	UserID    int    `json:"user_id"`
	UpdatedAt string `json:"updated_at"`
}
