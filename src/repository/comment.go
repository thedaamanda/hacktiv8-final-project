package repository

import (
	"context"
	"project/config"
	"project/src/model"
)

type commentRepository struct {
	Cfg config.Config
}

func NewCommentRepository(cfg config.Config) model.CommentRepository {
	return &commentRepository{
		Cfg: cfg,
	}
}

func (r *commentRepository) Create(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	if err := r.Cfg.Database().
		WithContext(ctx).
		Create(comment).
		Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *commentRepository) Fetch(ctx context.Context) ([]model.Comment, error) {
	var comments []model.Comment

	if err := r.Cfg.Database().
		WithContext(ctx).
		Preload("User").
		Preload("Photo").
		Find(&comments).
		Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *commentRepository) FindByID(ctx context.Context, id int) (*model.Comment, error) {
	var comment model.Comment

	if err := r.Cfg.Database().
		WithContext(ctx).
		Preload("User").
		Preload("Photo").
		Where("id = ?", id).
		First(&comment).
		Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *commentRepository) Update(ctx context.Context, id int, comment *model.Comment) (*model.Comment, error) {
	if err := r.Cfg.Database().
		WithContext(ctx).
		Where("id = ?", id).
		Updates(comment).
		Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *commentRepository) Delete(ctx context.Context, id int) error {
	if err := r.Cfg.Database().
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Comment{}).
		Error; err != nil {
		return err
	}

	return nil
}
