package repository

import (
	"context"
	"project/config"
	"project/src/model"
)

type socialMediaRepository struct {
	Cfg config.Config
}

func NewSocialMediaRepository(cfg config.Config) model.SocialMediaRepository {
	return &socialMediaRepository{
		Cfg: cfg,
	}
}

func (u *socialMediaRepository) Create(ctx context.Context, socialMedia *model.SocialMedia) (*model.SocialMedia, error) {
	if err := u.Cfg.Database().
		WithContext(ctx).
		Create(socialMedia).
		Error; err != nil {
		return nil, err
	}

	return socialMedia, nil
}

func (u *socialMediaRepository) Fetch(ctx context.Context) ([]model.SocialMedia, error) {
	var socialMedias []model.SocialMedia

	if err := u.Cfg.Database().
		WithContext(ctx).
		Preload("User").
		Find(&socialMedias).
		Error; err != nil {
		return nil, err
	}

	return socialMedias, nil
}

func (u *socialMediaRepository) FindByID(ctx context.Context, id int) (*model.SocialMedia, error) {
	var socialMedia model.SocialMedia

	if err := u.Cfg.Database().
		WithContext(ctx).
		Preload("User").
		Where("id = ?", id).
		First(&socialMedia).
		Error; err != nil {
		return nil, err
	}

	return &socialMedia, nil
}

func (u *socialMediaRepository) Update(ctx context.Context, id int, socialMedia *model.SocialMedia) (*model.SocialMedia, error) {
	if err := u.Cfg.Database().
		WithContext(ctx).
		Where("id = ?", id).
		Updates(socialMedia).
		Error; err != nil {
		return nil, err
	}

	return socialMedia, nil
}

func (u *socialMediaRepository) Delete(ctx context.Context, id int) error {
	if err := u.Cfg.Database().
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.SocialMedia{}).
		Error; err != nil {
		return err
	}

	return nil
}
