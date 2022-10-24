package usecase

import (
	"context"
	"errors"
	"project/src/model"
	"project/src/request"
	"project/src/response"
	"time"
)

type socialMediaUsecase struct {
	socialMediaRepo model.SocialMediaRepository
}

func NewSocialMediaUsecase(socialMedia model.SocialMediaRepository) model.SocialMediaUsecase {
	return &socialMediaUsecase{
		socialMediaRepo: socialMedia,
	}
}

func (u *socialMediaUsecase) CreateSocialMedia(ctx context.Context, id int, request request.SocialMediaRequest) (*response.SocialMediaResponse, error) {
	socialMedia := &model.SocialMedia{
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
		UserID:         id,
	}

	socialMedia, err := u.socialMediaRepo.Create(ctx, socialMedia)
	if err != nil {
		return nil, err
	}

	resp := new(response.SocialMediaResponse)
	resp.ID = socialMedia.ID
	resp.Name = socialMedia.Name
	resp.SocialMediaURL = socialMedia.SocialMediaURL
	resp.UserID = socialMedia.UserID
	resp.CreatedAt = socialMedia.CreatedAt.Format("2006-01-02 15:04:05")

	return resp, nil
}

func (u *socialMediaUsecase) GetSocialMediaList(ctx context.Context) ([]response.SocialMediasResponse, error) {
	socialMedias, err := u.socialMediaRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	var socialMediasResponse []response.SocialMediasResponse
	for _, socialMedia := range socialMedias {
		socialMediasResponse = append(socialMediasResponse, response.SocialMediasResponse{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaURL: socialMedia.SocialMediaURL,
			UserID:         socialMedia.UserID,
			CreatedAt:      socialMedia.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:      socialMedia.UpdatedAt.Format("2006-01-02 15:04:05"),
			User: response.SocialMediaUser{
				ID:       socialMedia.User.ID,
				Username: socialMedia.User.Username,
			},
		})
	}

	return socialMediasResponse, nil
}

func (u *socialMediaUsecase) UpdateSocialMedia(ctx context.Context, id int, request request.SocialMediaRequest) (*response.UpdateSocialMediaResponse, error) {
	socialMedia, err := u.socialMediaRepo.FindByID(ctx, id)

	if err != nil {
		return nil, errors.New("Sorry, we couldn’t find your social media in our records.")
	}

	socialMedia.Name = request.Name
	socialMedia.SocialMediaURL = request.SocialMediaURL
	socialMedia.UpdatedAt = time.Now()

	socialMedia, err = u.socialMediaRepo.Update(ctx, id, socialMedia)

	if err != nil {
		return nil, err
	}

	resp := new(response.UpdateSocialMediaResponse)
	resp.ID = socialMedia.ID
	resp.Name = socialMedia.Name
	resp.SocialMediaURL = socialMedia.SocialMediaURL
	resp.UserID = socialMedia.UserID
	resp.UpdatedAt = socialMedia.UpdatedAt.Format("2006-01-02 15:04:05")

	return resp, nil
}

func (u *socialMediaUsecase) DeleteSocialMedia(ctx context.Context, id int) error {
	if _, err := u.socialMediaRepo.FindByID(ctx, id); err != nil {
		return errors.New("Sorry, we couldn’t find your social media in our records.")
	}

	err := u.socialMediaRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
