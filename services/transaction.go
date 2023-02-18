package service

import (
	transactionCache "backend-api/cache"
	transactiondto "backend-api/dto/transaction"
	"backend-api/models"
	"context"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type TransactionService interface {
	FindAllTransactions() (*[]models.Transaction, error)
	GetTransactionID(idTrx int) (*transactiondto.TransactionResponse, error)
	CreateTransaction(userID int, request transactiondto.TransactionRequest) (*transactiondto.TransactionResponse, error)
	UpdateTransaction(idTrx int) (*transactiondto.TransactionResponse, error)
	DeleteTransaction(idTrx int) (*transactiondto.TransactionResponse, error)
}

type transactionService struct {
	transactionCache transactionCache.TransactionCache
}

func NewTransactionService(transactionCache transactionCache.TransactionCache) *transactionService {
	return &transactionService{transactionCache}
}

func (s *transactionService) FindAllTransactions() (*[]models.Transaction, error) {

	response, err := s.transactionCache.FindAllTransactions()
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *transactionService) GetTransactionID(idTrx int) (*transactiondto.TransactionResponse, error) {

	transaction, err := s.transactionCache.GetTransactionID(idTrx)
	if err != nil {
		return nil, err
	}

	// Construct and return response
	response := transactiondto.TransactionResponse{
		ID:             transaction.ID,
		AccountNumber:  transaction.AccountNumber,
		ProofOfTranser: transaction.ProofOfTranser,
		Status:         transaction.Status,
		User:           transaction.User.Name,
		UpdatedAt:      transaction.UpdatedAt.Format("02-01-2006 15:04"),
	}

	return &response, nil
}

func (s *transactionService) CreateTransaction(userID int, request transactiondto.TransactionRequest) (*transactiondto.TransactionResponse, error) {
	// Upload proof_of_transfer to cloudinary
	ctx := context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cloud, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	proof, err := cloud.Upload.Upload(ctx, request.ProofOfTranser, uploader.UploadParams{Folder: "waysbuck"})
	if err != nil {
		return nil, err
	}

	// ambil id & harga yang di dapat dari paket yang di pilih dari tabel subscription
	// do something

	// Create new transaction model instance
	idTrx := time.Now().Unix()
	transaction := models.Transaction{
		ID:             int(idTrx),
		AccountNumber:  request.AccountNumber,
		ProofOfTranser: proof.SecureURL,
		// Amount:         45000,
		// Subscription:   "not premium",
		Status: "waiting approve",
		UserID: userID,
	}

	// Store transaction data into cache
	data, err := s.transactionCache.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}

	// Construct and return response
	response := transactiondto.TransactionResponse{
		ID:             data.ID,
		AccountNumber:  data.AccountNumber,
		ProofOfTranser: data.ProofOfTranser,
		Status:         data.Status,
		// Amount:         data.SubscriptionID.Amount,
		// Subscription:   data.Subscription,
		User:      data.User.Name,
		UpdatedAt: data.UpdatedAt.Format("02-01-2006 15:04"),
	}

	return &response, nil
}

func (s *transactionService) UpdateTransaction(idTrx int) (*transactiondto.TransactionResponse, error) {

	transaction, err := s.transactionCache.GetTransactionID(idTrx)
	if err != nil {
		return nil, err
	}

	// transaction.Subscription = "premium"
	transaction.Status = "approved"

	// Store transaction data into cache
	data, err := s.transactionCache.UpdateTransaction(transaction)
	if err != nil {
		return nil, err
	}

	// Construct and return response
	response := transactiondto.TransactionResponse{
		ID:             data.ID,
		AccountNumber:  data.AccountNumber,
		ProofOfTranser: data.ProofOfTranser,
		Status:         data.Status,
		// Amount:         data.Amount,
		// Subscription:   data.Subscription,
		User:      data.User.Name,
		UpdatedAt: data.UpdatedAt.Format("02-01-2006 15:04"),
	}

	return &response, nil
}

func (s *transactionService) DeleteTransaction(idTrx int) (*transactiondto.TransactionResponse, error) {

	transaction, err := s.transactionCache.GetTransactionID(idTrx)
	if err != nil {
		return nil, err
	}

	data, err := s.transactionCache.DeleteTransaction(transaction)
	if err != nil {
		return nil, err
	}

	response := transactiondto.TransactionResponse{
		ID:             data.ID,
		AccountNumber:  data.AccountNumber,
		ProofOfTranser: data.ProofOfTranser,
		Status:         data.Status,
		// Amount:         data.Amount,
		// Subscription:   data.Subscription,
		User:      data.User.Name,
		UpdatedAt: data.UpdatedAt.Format("02-01-2006 15:04"),
	}

	return &response, nil
}
