package cache

import (
	"backend-api/models"
	"backend-api/repositories"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type AuthCache interface {
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
	GetUserID(ID int) (models.User, error)
}

type authCache struct {
	authRepository repositories.AuthRepository
	rdb            *redis.Client
}

func NewAuthCache(authRepository repositories.AuthRepository, rdb *redis.Client) *authCache {
	return &authCache{authRepository, rdb}
}

func (c *authCache) Register(user models.User) (models.User, error) {
	user, err := c.authRepository.Register(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (c *authCache) Login(email string) (models.User, error) {
	var data models.User

	cacheKey := fmt.Sprintf("user:%s", email)

	if cacheData, err := c.rdb.Get(c.rdb.Context(), cacheKey).Result(); err == nil {
		if err := json.Unmarshal([]byte(cacheData), &data); err != nil {
			return data, err
		}
		return data, nil
	}

	user, err := c.authRepository.Login(email)
	if err != nil {
		return user, err
	}

	cacheData, err := json.Marshal(user)
	if err != nil {
		return user, err
	}

	c.rdb.Set(c.rdb.Context(), cacheKey, cacheData, time.Hour)
	return user, nil
}

func (c *authCache) GetUserID(ID int) (models.User, error) {
	var data models.User

	cacheKey := fmt.Sprintf("user:%v", ID)

	if cacheData, err := c.rdb.Get(c.rdb.Context(), cacheKey).Result(); err == nil {
		if err := json.Unmarshal([]byte(cacheData), &data); err != nil {
			return data, err
		}
		return data, nil
	}

	user, err := c.authRepository.GetUserID(ID)
	if err != nil {
		return user, err
	}

	cacheData, err := json.Marshal(user)
	if err != nil {
		return user, err
	}

	c.rdb.Set(c.rdb.Context(), cacheKey, cacheData, time.Hour)
	return user, nil
}
