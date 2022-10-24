package model

import (
	"context"
	"project/src/request"
	"project/src/response"
	"time"
)

type Comment struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id" gorm:"foreignkey:UserID"`
	PhotoID   int       `json:"photo_id" gorm:"foreignkey:PhotoID"`
	Message   string    `json:"message" gorm:"type:varchar(200);not null"`
	User      User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Photo     Photo     `json:"photo" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentRepository interface {
	Create(ctx context.Context, comment *Comment) (*Comment, error)
	Fetch(ctx context.Context) ([]Comment, error)
	FindByID(ctx context.Context, id int) (*Comment, error)
	Update(ctx context.Context, id int, comment *Comment) (*Comment, error)
	Delete(ctx context.Context, id int) error
}

type CommentUsecase interface {
	CreateComment(ctx context.Context, id int, request request.CreateCommentRequest) (*response.CommentResponse, error)
	GetCommentList(ctx context.Context) ([]response.CommentsResponse, error)
	UpdateComment(ctx context.Context, id int, request request.UpdateCommentRequest) (*response.UpdateCommentResponse, error)
	DeleteComment(ctx context.Context, id int) error
}
