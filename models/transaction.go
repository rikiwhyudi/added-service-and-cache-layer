package models

import "time"

type Transaction struct {
	ID             int       `json:"id" gorm:"primary_key:auto_increment"`
	AccountNumber  string    `json:"account_number gorm:type varchar(255)"`
	ProofOfTranser string    `json:"proof_of_transer" gorm:"type varchar(255)"`
	Status         string    `json:"status" gorm:"type varchar(255)"`
	UserID         int       `json:"user_id" gorm:"foreignkey:UserID"` //has many fields from user
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type TransactionResponse struct {
	ID             int          `json:"id"`
	AccountNumber  string       `json:"account_number"`
	ProofOfTranser string       `json:"proof_of_transer"`
	Status         string       `json:"status"`
	User           UserResponse `json:"user"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
