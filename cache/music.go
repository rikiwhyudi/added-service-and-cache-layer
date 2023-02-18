package cache

import (
	"backend-api/models"
	"backend-api/repositories"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type MusicCache interface {
	FindAllMusics() ([]models.Music, error)
	GetMusicID(ID int) (models.Music, error)
	CreateMusic(music models.Music) (models.Music, error)
	UpdateMusic(music models.Music) (models.Music, error)
	DeleteMusic(music models.Music) (models.Music, error)
}

type musicCache struct {
	musicRepository repositories.MusicRepository
	rdb             *redis.Client
}

func NewMusicCache(musicRepository repositories.MusicRepository, rdb *redis.Client) *musicCache {
	return &musicCache{musicRepository, rdb}
}

func (c *musicCache) FindAllMusics() ([]models.Music, error) {
	var musics []models.Music
	cacheKey := fmt.Sprintf("data:%s", "musics")

	if cacheData, err := c.rdb.Get(c.rdb.Context(), cacheKey).Result(); err == nil {
		if err := json.Unmarshal([]byte(cacheData), &musics); err != nil {
			return musics, err
		}

		return musics, nil
	}

	data, err := c.musicRepository.FindAllMusics()
	if err != nil {
		return data, err
	}

	cacheData, err := json.Marshal(data)
	if err != nil {
		return data, err
	}

	c.rdb.Set(c.rdb.Context(), cacheKey, cacheData, time.Hour)
	return data, nil

}

func (c *musicCache) GetMusicID(ID int) (models.Music, error) {
	var music models.Music
	cacheKey := fmt.Sprintf("music:%v", ID)

	if cacheData, err := c.rdb.Get(c.rdb.Context(), cacheKey).Result(); err == nil {
		if err := json.Unmarshal([]byte(cacheData), &music); err != nil {
			return music, err
		}

		return music, nil
	}

	data, err := c.musicRepository.GetMusicID(ID)
	if err != nil {
		return data, err
	}

	cacheData, err := json.Marshal(data)
	if err != nil {
		return data, err
	}

	c.rdb.Set(c.rdb.Context(), cacheKey, cacheData, time.Hour)
	return data, nil
}

func (c *musicCache) CreateMusic(music models.Music) (models.Music, error) {
	music, err := c.musicRepository.CreateMusic(music)
	if err != nil {
		return music, err
	}

	return music, nil
}

func (c *musicCache) UpdateMusic(music models.Music) (models.Music, error) {
	music, err := c.musicRepository.UpdateMusic(music)
	if err != nil {
		return music, err
	}

	return music, nil
}

func (c *musicCache) DeleteMusic(music models.Music) (models.Music, error) {
	music, err := c.musicRepository.DeleteMusic(music)
	if err != nil {
		return music, err
	}

	return music, nil
}
