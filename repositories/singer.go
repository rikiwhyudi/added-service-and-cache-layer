package repositories

import (
	"backend-api/models"

	"gorm.io/gorm"
)

type SingerRepository interface {
	FindAllSingers() ([]models.Singer, error)
	GetSingerID(ID int) (models.Singer, error)
	CreateSinger(singer models.Singer) (models.Singer, error)
	UpdateSinger(singer models.Singer) (models.Singer, error)
	DeleteSinger(singer models.Singer) (models.Singer, error)
}

func RepositorySinger(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllSingers() ([]models.Singer, error) {
	var singers []models.Singer
	err := r.db.Preload("Music").Find(&singers).Error

	return singers, err
}

func (r *repository) GetSingerID(ID int) (models.Singer, error) {
	var singer models.Singer
	err := r.db.Preload("Music").First(&singer, ID).Error

	return singer, err

}

func (r *repository) CreateSinger(singer models.Singer) (models.Singer, error) {
	err := r.db.Create(&singer).Error

	return singer, err
}

func (r *repository) UpdateSinger(singer models.Singer) (models.Singer, error) {
	err := r.db.Save(&singer).Error

	return singer, err
}

func (r *repository) DeleteSinger(singer models.Singer) (models.Singer, error) {
	err := r.db.Delete(&singer).Error

	return singer, err
}
