package service

import (
	cacheMusics "backend-api/cache"
	musicdto "backend-api/dto/music"
	singerdto "backend-api/dto/singer"
	"backend-api/models"
	// "context"
	// "os"
	// "github.com/cloudinary/cloudinary-go/v2"
	// "github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type MusicService interface {
	FindAllMusics() (*[]musicdto.MusicResponse, error)
	GetMusicID(ID int) (*musicdto.MusicResponse, error)
	CreateMusic(request musicdto.MusicRequest) (*musicdto.MusicResponse, error)
	UpdateMusic(request musicdto.UpdatedMusicRequest, ID int) (*musicdto.MusicResponse, error)
	DeleteMusic(ID int) (*musicdto.MusicResponse, error)
}

type musicService struct {
	musicCache cacheMusics.MusicCache
}

func NewMusicService(musicCache cacheMusics.MusicCache) *musicService {
	return &musicService{musicCache}
}

func (s *musicService) FindAllMusics() (*[]musicdto.MusicResponse, error) {
	music, err := s.musicCache.FindAllMusics()
	if err != nil {
		return nil, err
	}

	// Construct, loop and return response
	response := make([]musicdto.MusicResponse, 0)
	for _, data := range music {
		singerResponse := singerdto.SingerResponse{
			ID:          data.Singer.ID,
			Name:        data.Singer.Name,
			Thumbnail:   data.Singer.Thumbnail,
			Old:         data.Singer.Old,
			Category:    data.Singer.Category,
			StartCareer: data.Singer.StartCareer.Format("02-01-2006"),
		}

		musicResponse := musicdto.MusicResponse{
			ID:         data.ID,
			Title:      data.Title,
			Year:       data.Year.Format("02-01-2006"),
			SingerName: singerResponse,
		}

		response = append(response, musicResponse)
	}

	return &response, nil
}

func (s *musicService) GetMusicID(ID int) (*musicdto.MusicResponse, error) {

	data, err := s.musicCache.GetMusicID(ID)
	if err != nil {
		return nil, err
	}

	// Construct and return response
	singerResponse := singerdto.SingerResponse{
		ID:          data.Singer.ID,
		Name:        data.Singer.Name,
		Thumbnail:   data.Singer.Thumbnail,
		Old:         data.Singer.Old,
		Category:    data.Singer.Category,
		StartCareer: data.Singer.StartCareer.Format("02-01-2006"),
	}

	response := musicdto.MusicResponse{
		ID:         data.ID,
		Title:      data.Title,
		Year:       data.Year.Format("02-01-2006"),
		SingerName: singerResponse,
		// Thumbnail:  data.Thumbnail,
		// MusicFile:  data.MusicFile,
	}

	return &response, nil
}

func (s *musicService) CreateMusic(request musicdto.MusicRequest) (*musicdto.MusicResponse, error) {
	// Upload music file & thumbnail to cloudinary
	// ctx := context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")

	// cloud, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// thumbnail, err := cloud.Upload.Upload(ctx, request.Thumbnail, uploader.UploadParams{Folder: "waysbuck"})
	// if err != nil {
	// 	return nil, err
	// }

	// musicFile, err := cloud.Upload.Upload(ctx, request.MusicFile, uploader.UploadParams{Folder: "waysbuck"})
	// if err != nil {
	// 	return nil, err
	// }

	// Create new music models instance
	music := models.Music{
		Title: request.Title,
		// Thumbnail: thumbnail.SecureURL,
		Year:     request.Year,
		SingerID: request.SingerID,
		// MusicFile: musicFile.SecureURL,
	}

	// Store music data into cache
	data, err := s.musicCache.CreateMusic(music)
	if err != nil {
		return nil, err
	}

	response := musicdto.MusicResponse{
		ID:         data.ID,
		Title:      data.Title,
		Year:       data.Year.Format("02-01-2006"),
		SingerName: data.SingerID,
		// Thumbnail:  data.Thumbnail,
		// MusicFile:  data.MusicFile,
	}

	return &response, nil
}

func (s *musicService) UpdateMusic(request musicdto.UpdatedMusicRequest, ID int) (*musicdto.MusicResponse, error) {
	// Upload music file & thumbnail to cloudinary
	// ctx := context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")

	// cloud, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// thumbnail, err := cloud.Upload.Upload(ctx, request.Thumbnail, uploader.UploadParams{Folder: "waysbuck"})
	// if err != nil {
	// 	return nil, err
	// }

	// musicFile, err := cloud.Upload.Upload(ctx, request.MusicFile, uploader.UploadParams{Folder: "waysbuck"})
	// if err != nil {
	// 	return nil, err
	// }

	music, err := s.musicCache.GetMusicID(ID)
	if err != nil {
		return nil, err
	}

	if request.Title != "" {
		music.Title = request.Title
	}

	// if request.Thumbnail != "" {
	// 	music.Thumbnail = thumbnail.SecureURL
	// }

	if !request.Year.IsZero() {
		music.Year = request.Year
	}

	// if request.SingerID != 0 {
	// 	music.SingerID = request.SingerID
	// }

	// if request.Thumbnail != "false" {
	// 	music.Thumbnail = thumbnail.SecureURL

	// }

	// Store music data into cache
	data, err := s.musicCache.UpdateMusic(music)
	if err != nil {
		return nil, err
	}

	// Construct and return response
	response := musicdto.MusicResponse{
		ID:    data.ID,
		Title: data.Title,
		// SingerName: data.Singer.Name,
		// Thumbnail:  data.Thumbnail,
		Year: data.Year.Format("02-01-2006"),
		// MusicFile:  data.MusicFile,
	}

	return &response, nil
}

func (s *musicService) DeleteMusic(ID int) (*musicdto.MusicResponse, error) {

	music, err := s.musicCache.GetMusicID(ID)
	if err != nil {
		return nil, err
	}

	data, err := s.musicCache.DeleteMusic(music)
	if err != nil {
		return nil, err
	}

	// Construct and return response
	response := musicdto.MusicResponse{
		ID:    data.ID,
		Title: data.Title,
		// SingerName: data.Singer.Name,
		// Thumbnail:  data.Thumbnail,
		Year: data.Year.Format("02-01-2006"),
		// MusicFile:  data.MusicFile,
	}

	return &response, nil
}
