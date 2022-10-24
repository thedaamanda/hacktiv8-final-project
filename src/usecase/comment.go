package usecase

import (
	"context"
	"errors"
	"project/src/model"
	"project/src/request"
	"project/src/response"
	"time"
)

type commentUsecase struct {
	commentRepo model.CommentRepository
	photoRepo   model.PhotoRepository
}

func NewCommentUsecase(comment model.CommentRepository, photo model.PhotoRepository) model.CommentUsecase {
	return &commentUsecase{
		commentRepo: comment,
		photoRepo:   photo,
	}
}

func (u *commentUsecase) CreateComment(ctx context.Context, id int, request request.CreateCommentRequest) (*response.CommentResponse, error) {
	if _, err := u.photoRepo.FindByID(ctx, request.PhotoID); err != nil {
		return nil, errors.New("Couldn’t find the photo you’re trying to comment on.")
	}

	comment := &model.Comment{
		Message: request.Message,
		PhotoID: request.PhotoID,
		UserID:  id,
	}

	comment, err := u.commentRepo.Create(ctx, comment)
	if err != nil {
		return nil, err
	}

	return &response.CommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (u *commentUsecase) GetCommentList(ctx context.Context) ([]response.CommentsResponse, error) {
	comments, err := u.commentRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	var commentsResponse []response.CommentsResponse
	for _, comment := range comments {
		commentsResponse = append(commentsResponse, response.CommentsResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: comment.UpdatedAt.Format("2006-01-02 15:04:05"),
			User: response.CommentUser{
				ID:       comment.User.ID,
				Username: comment.User.Username,
				Email:    comment.User.Email,
			},
			Photo: response.CommentPhoto{
				ID:       comment.Photo.ID,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoURL: comment.Photo.PhotoURL,
				UserID:   comment.Photo.UserID,
			},
		})
	}

	return commentsResponse, nil
}

func (u *commentUsecase) UpdateComment(ctx context.Context, id int, request request.UpdateCommentRequest) (*response.UpdateCommentResponse, error) {
	comment, err := u.commentRepo.FindByID(ctx, id)

	if err != nil {
		return nil, errors.New("Sorry, we couldn’t find your comment in our records.")
	}

	comment.Message = request.Message
	comment.UpdatedAt = time.Now()

	comment, err = u.commentRepo.Update(ctx, id, comment)

	if err != nil {
		return nil, err
	}

	resp := new(response.UpdateCommentResponse)
	resp.ID = comment.ID
	resp.Message = comment.Message
	resp.PhotoID = comment.PhotoID
	resp.UserID = comment.UserID
	resp.UpdatedAt = comment.UpdatedAt.Format("2006-01-02 15:04:05")

	return resp, nil
}

func (u *commentUsecase) DeleteComment(ctx context.Context, id int) error {
	if _, err := u.commentRepo.FindByID(ctx, id); err != nil {
		return errors.New("Sorry, we couldn’t find your comment in our records.")
	}

	err := u.commentRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
