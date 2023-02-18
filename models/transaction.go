package models

import "time"

type Transaction struct {
	ID             int    `json:"id" gorm:"primary_key:auto_increment"`
	AccountNumber  string `json:"account_number gorm:type varchar(255)"`
	ProofOfTranser string `json:"proof_of_transer" gorm:"type varchar(255)"`
	Status         string `json:"status" gorm:"type varchar(255)"`
	//catatan untuk saya, amount dan subscription ini bisa di pisah jadi tabel paket subscription kedepanya
	// Amount       int          `json:"amount"`
	// SubscriptionID Subscription `json:"subscription" gorm:"foreignKey:transactionID"` // has many
	UserID    int          `json:"user_id" gorm:"foreignkey:UserID"` //has many fields from user
	User      UserResponse `json:"user"`
	UpdatedAt time.Time    `json:"updated_at"`
}

// type TransactionResponse struct {
// 	ID             int          `json:"id"`
// 	AccountNumber  string       `json:"account_number"`
// 	ProofOfTranser string       `json:"proof_of_transer"`
// 	Status         string       `json:"status"`
// 	User           UserResponse `json:"user"`
// }

func (Transaction) TableName() string {
	return "transactions"
}
