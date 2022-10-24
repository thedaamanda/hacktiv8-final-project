package model

import (
	"context"
	"project/src/request"
	"project/src/response"
	"time"
)

type SocialMedia struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"type:varchar(100);not null"`
	SocialMediaURL string    `json:"social_media_url" gorm:"type:varchar(200);not null"`
	UserID         int       `json:"user_id" gorm:"foreignKey:UserID"`
	User           User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaRepository interface {
	Create(ctx context.Context, socialMedia *SocialMedia) (*SocialMedia, error)
	Fetch(ctx context.Context) ([]SocialMedia, error)
	FindByID(ctx context.Context, id int) (*SocialMedia, error)
	Update(ctx context.Context, id int, socialMedia *SocialMedia) (*SocialMedia, error)
	Delete(ctx context.Context, id int) error
}

type SocialMediaUsecase interface {
	CreateSocialMedia(ctx context.Context, id int, request request.SocialMediaRequest) (*response.SocialMediaResponse, error)
	GetSocialMediaList(ctx context.Context) ([]response.SocialMediasResponse, error)
	UpdateSocialMedia(ctx context.Context, id int, request request.SocialMediaRequest) (*response.UpdateSocialMediaResponse, error)
	DeleteSocialMedia(ctx context.Context, id int) error
}
