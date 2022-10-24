package response

type (
	SocialMediaResponse struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		SocialMediaURL string `json:"social_media_url"`
		UserID         int    `json:"user_id"`
		CreatedAt      string `json:"created_at"`
	}

	SocialMediasResponse struct {
		ID             int             `json:"id"`
		Name           string          `json:"name"`
		SocialMediaURL string          `json:"social_media_url"`
		UserID         int             `json:"user_id"`
		CreatedAt      string          `json:"created_at"`
		UpdatedAt      string          `json:"updated_at"`
		User           SocialMediaUser `json:"User"`
	}

	SocialMediaUser struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}

	UpdateSocialMediaResponse struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		SocialMediaURL string `json:"social_media_url"`
		UserID         int    `json:"user_id"`
		UpdatedAt      string `json:"updated_at"`
	}
)
