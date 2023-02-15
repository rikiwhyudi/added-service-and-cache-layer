package service

import (
	singerCache "backend-api/cache"
	singerdto "backend-api/dto/singer"
	"backend-api/models"
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
)

type SingerService interface {
	FindAllSingers() ([]models.Singer, error)
	GetSinger(ID int) (*singerdto.SingerResponse, error)
	CreateSinger(request singerdto.SingerRequest) (*singerdto.SingerResponse, error)
	// UpdateSinger(singer singerdto.SingerRequest) (*singerdto.SingerResponse, error)
	// DeleteSinger(singer singerdto.SingerRequest) (*singerdto.SingerResponse, error)
}

type singerService struct {
	singerCache singerCache.SingerCache
	validator   *validator.Validate
}

func NewSingerService(singerCache singerCache.SingerCache) *singerService {
	return &singerService{singerCache, validator.New()}
}

func (s *singerService) FindAllSingers() ([]models.Singer, error) {

	singers, err := s.singerCache.FindAllSingers()
	if err != nil {
		return nil, err
	}

	return singers, err
}

func (s *singerService) GetSinger(ID int) (*singerdto.SingerResponse, error) {

	singer, err := s.singerCache.GetSinger(ID)
	if err != nil {
		return nil, err
	}

	response := singerdto.SingerResponse{
		ID:          singer.ID,
		Name:        singer.Name,
		Thumbnail:   singer.Thumbnail,
		Old:         singer.Old,
		Category:    singer.Category,
		StartCareer: singer.StartCareer,
	}

	return &response, nil
}

func (s *singerService) CreateSinger(request singerdto.SingerRequest) (*singerdto.SingerResponse, error) {
	// Validate request input using go-playground/validator
	err := s.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	// Upload singer thumbnail to cloudinary
	ctx := context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, request.Thumbnail, uploader.UploadParams{Folder: "waysbuck"})
	if err != nil {
		return nil, err
	}

	// Create new singer model instance
	singer := models.Singer{
		Name:        request.Name,
		Thumbnail:   resp.SecureURL,
		Old:         request.Old,
		Category:    request.Category,
		StartCareer: request.StartCareer,
	}

	// Store singer data into cache
	data, err := s.singerCache.CreateSinger(singer)
	if err != nil {
		return nil, err
	}

	// Construct and return response
	response := singerdto.SingerResponse{
		Name:        data.Name,
		Thumbnail:   data.Thumbnail,
		Old:         data.Old,
		Category:    data.Category,
		StartCareer: data.StartCareer,
	}

	return &response, nil
}
