package models

import "time"

type Subscription struct {
	ID            int       `json:"id" gorm:"primary_key:auto_increment"`
	Amount        int       `json:"amount"`
	Subscription  string    `json:"subscription" gorm:"type varchar(255)"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	TransactionID int       `json:"transaction_id" gorm:"foreignKey:TransactionID"` //has many fields from transaction
	// Music         []Music   `json:"music" gorm:"foreignKey:SubscriptionID"`         // has many
}

func (Subscription) TableName() string {
	return "subscriptions"
}
