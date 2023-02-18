package service

import (
	singerCache "backend-api/cache"
	musicdto "backend-api/dto/music"
	singerdto "backend-api/dto/singer"
	"backend-api/models"
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type SingerService interface {
	FindAllSingers() (*[]singerdto.AllSingerResponse, error)
	GetSingerID(ID int) (*singerdto.AllSingerResponse, error)
	CreateSinger(request singerdto.SingerRequest) (*singerdto.SingerResponse, error)
	UpdateSinger(request singerdto.UpdateSingerRequest, ID int) (*singerdto.SingerResponse, error)
	DeleteSinger(ID int) (*singerdto.SingerResponse, error)
}

type singerService struct {
	singerCache singerCache.SingerCache
}

func NewSingerService(singerCache singerCache.SingerCache) *singerService {
	return &singerService{singerCache}
}

func (s *singerService) FindAllSingers() (*[]singerdto.AllSingerResponse, error) {

	singer, err := s.singerCache.FindAllSingers()
	if err != nil {
		return nil, err
	}

	// Construc, loop and return response
	response := make([]singerdto.AllSingerResponse, 0)
	for _, data := range singer {
		musicResponse := make([]musicdto.AllMusicResponse, 0)
		for _, music := range data.Music {
			dataMusic := musicdto.AllMusicResponse{
				ID:    music.ID,
				Title: music.Title,
				Year:  music.Year.Format("02-01-2006"),
			}

			musicResponse = append(musicResponse, dataMusic)
		}

		singerResponse := singerdto.AllSingerResponse{
			ID:          data.ID,
			Name:        data.Name,
			Thumbnail:   data.Thumbnail,
			Old:         data.Old,
			Category:    data.Category,
			StartCareer: data.StartCareer.Format("02-01-2006"),
			Musics:      musicResponse,
		}

		response = append(response, singerResponse)
	}

	return &response, nil

}

func (s *singerService) GetSingerID(ID int) (*singerdto.AllSingerResponse, error) {

	singer, err := s.singerCache.GetSingerID(ID)
	if err != nil {
		return nil, err
	}

	// Construct and return response
	musicResponse := make([]musicdto.AllMusicResponse, 0)
	for _, music := range singer.Music {
		dataMusic := musicdto.AllMusicResponse{
			ID:    music.ID,
			Title: music.Title,
			Year:  music.Year.Format("02-01-2006"),
		}
		musicResponse = append(musicResponse, dataMusic)
	}

	response := singerdto.AllSingerResponse{
		ID:          singer.ID,
		Name:        singer.Name,
		Thumbnail:   singer.Thumbnail,
		Old:         singer.Old,
		Category:    singer.Category,
		StartCareer: singer.StartCareer.Format("02-01-2006"),
		Musics:      musicResponse,
	}

	return &response, nil
}

func (s *singerService) CreateSinger(request singerdto.SingerRequest) (*singerdto.SingerResponse, error) {
	// Upload singer thumbnail to cloudinary
	ctx := context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cloud, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cloud.Upload.Upload(ctx, request.Thumbnail, uploader.UploadParams{Folder: "waysbuck"})
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
		ID:          data.ID,
		Name:        data.Name,
		Thumbnail:   data.Thumbnail,
		Old:         data.Old,
		Category:    data.Category,
		StartCareer: data.StartCareer.Format("02-01-2006"),
	}

	return &response, nil
}

func (s *singerService) UpdateSinger(request singerdto.UpdateSingerRequest, ID int) (*singerdto.SingerResponse, error) {

	// Upload singer thumbnail to cloudinary
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cloud, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	thumbnail, _ := cloud.Upload.Upload(ctx, request.Thumbnail, uploader.UploadParams{Folder: "waysbuck"})

	singer, err := s.singerCache.GetSingerID(ID)
	if err != nil {
		return nil, err
	}

	if request.Name != "" {
		singer.Name = request.Name
	}

	if request.Thumbnail != "false" {
		singer.Thumbnail = thumbnail.SecureURL
	}

	if request.Old != 0 {
		singer.Old = request.Old
	}

	if request.Category != "" {
		singer.Category = request.Category
	}

	if !request.StartCareer.IsZero() {
		singer.StartCareer = request.StartCareer
	}

	data, err := s.singerCache.UpdateSinger(singer)
	if err != nil {
		return nil, err
	}

	// Construct and return response
	response := singerdto.SingerResponse{
		ID:          data.ID,
		Name:        data.Name,
		Thumbnail:   data.Thumbnail,
		Old:         data.Old,
		Category:    data.Category,
		StartCareer: data.StartCareer.Format("02-01-2006"),
	}

	return &response, nil
}

func (s *singerService) DeleteSinger(ID int) (*singerdto.SingerResponse, error) {

	singer, err := s.singerCache.GetSingerID(ID)
	if err != nil {
		return nil, err
	}

	data, err := s.singerCache.DeleteSinger(singer)
	if err != nil {
		return nil, err
	}

	// Construct and return response
	response := singerdto.SingerResponse{
		ID:          data.ID,
		Name:        data.Name,
		Thumbnail:   data.Thumbnail,
		Old:         data.Old,
		Category:    data.Category,
		StartCareer: data.StartCareer.Format("02-01-2006"),
	}

	return &response, nil
}
