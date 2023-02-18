package transactiondto

type TransactionResponse struct {
	ID             int    `json:"id"`
	AccountNumber  string `json:"account_number"`
	ProofOfTranser string `json:"proof_of_transer"`
	Status         string `json:"status"`
	Amount         int    `json:"amount"`
	Subscription   string `json:"subscription"`
	User           string `json:"user"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
