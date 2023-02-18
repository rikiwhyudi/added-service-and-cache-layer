package cache

import (
	"backend-api/models"
	"backend-api/repositories"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type SingerCache interface {
	FindAllSingers() ([]models.Singer, error)
	GetSingerID(ID int) (models.Singer, error)
	CreateSinger(singer models.Singer) (models.Singer, error)
	UpdateSinger(singer models.Singer) (models.Singer, error)
	DeleteSinger(singer models.Singer) (models.Singer, error)
}

type singerCache struct {
	singerRepository repositories.SingerRepository
	rdb              *redis.Client
}

func NewSingerCache(singerRepository repositories.SingerRepository, rdb *redis.Client) *singerCache {
	return &singerCache{singerRepository, rdb}
}

func (c *singerCache) FindAllSingers() ([]models.Singer, error) {
	var singers []models.Singer
	cacheKey := fmt.Sprintf("data:%s", "singers")

	if cacheData, err := c.rdb.Get(c.rdb.Context(), cacheKey).Result(); err == nil {
		if err := json.Unmarshal([]byte(cacheData), &singers); err != nil {
			return singers, err
		}
		return singers, nil
	}

	data, err := c.singerRepository.FindAllSingers()
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

func (c *singerCache) GetSingerID(ID int) (models.Singer, error) {
	var singer models.Singer
	cacheKey := fmt.Sprintf("singer:%v", ID)

	if cacheData, err := c.rdb.Get(c.rdb.Context(), cacheKey).Result(); err == nil {
		if err := json.Unmarshal([]byte(cacheData), &singer); err != nil {
			return singer, err
		}
		return singer, nil
	}

	data, err := c.singerRepository.GetSingerID(ID)
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

func (c *singerCache) CreateSinger(singer models.Singer) (models.Singer, error) {
	singer, err := c.singerRepository.CreateSinger(singer)
	if err != nil {
		return singer, err
	}

	return singer, nil
}

func (c *singerCache) UpdateSinger(singer models.Singer) (models.Singer, error) {
	singer, err := c.singerRepository.UpdateSinger(singer)
	if err != nil {
		return singer, err
	}

	return singer, nil
}

func (c *singerCache) DeleteSinger(singer models.Singer) (models.Singer, error) {
	singer, err := c.singerRepository.DeleteSinger(singer)
	if err != nil {
		return singer, err
	}

	return singer, nil
}
