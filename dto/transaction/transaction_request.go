package transactiondto

type TransactionRequest struct {
	AccountNumber  string `json:"account_number" form:"account_number" validate:"required"`
	ProofOfTranser string `json:"proof_of_transer" form:"proof_of_transfer" validate:"required"`
}

type UpdateTransactionRequest struct {
	TransactionID string `json:"transaction_id" form:"transaction_id" validate:"required"`
}
