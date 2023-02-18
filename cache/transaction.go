package cache

import (
	"backend-api/models"
	"backend-api/repositories"

	"github.com/go-redis/redis/v8"
)

type TransactionCache interface {
	FindAllTransactions() ([]models.Transaction, error)
	GetTransactionID(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction) (models.Transaction, error)
}

type transactionCache struct {
	transactionRepository repositories.TransactionRepository
	rdb                   *redis.Client
}

func NewTransactionCache(rdb *redis.Client, transactionRepository repositories.TransactionRepository) *transactionCache {
	return &transactionCache{transactionRepository, rdb}
}

func (c *transactionCache) FindAllTransactions() ([]models.Transaction, error) {

	data, err := c.transactionRepository.FindAllTransactions()
	if err != nil {
		return data, err
	}

	return data, nil
}

func (c *transactionCache) GetTransactionID(ID int) (models.Transaction, error) {

	data, err := c.transactionRepository.GetTransactionID(ID)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (c *transactionCache) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {

	data, err := c.transactionRepository.CreateTransaction(transaction)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (c *transactionCache) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {

	data, err := c.transactionRepository.CreateTransaction(transaction)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (c *transactionCache) DeleteTransaction(transaction models.Transaction) (models.Transaction, error) {

	data, err := c.transactionRepository.DeleteTransaction(transaction)
	if err != nil {
		return data, err
	}

	return data, nil
}
