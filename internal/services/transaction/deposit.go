package transaction

import (
	"context"
	"fmt"
	"log"
	"payment-gateway/db"
	"payment-gateway/enum"
	"payment-gateway/internal/services/transaction/contract"
	ppContract "payment-gateway/paymentprovider/contract"
	"time"
)

func (i impl) Deposit(ctx context.Context, req *contract.DepositRequest) (*contract.DepositResponse, error) {
	// call payment gateway provider to make deposit
	if err := i.paymentProvider.Deposit(ctx, &ppContract.DepositRequest{}); err != nil {
		log.Println(fmt.Sprintf("error from payment provider: %v userID: %d provider: %s", err, req.UserID, req.GatewayProvider))
		return &contract.DepositResponse{
			TransactionID: 0,
			Success:       false,
		}, nil
	}

	// save to transaction table
	txnModel := &db.Transaction{
		Amount:    req.Amount,
		Type:      enum.TxnDeposit.String(),
		Status:    enum.TxnStatusInitialized.String(), // default
		UserID:    req.UserID,
		GatewayID: req.GatewayID,
		CountryID: req.CountryID,
		CreatedAt: time.Now().UTC(),
	}
	id, err := i.db.CreateTransaction(ctx, txnModel)
	if err != nil {
		errMsg := fmt.Errorf("failed to save to txn table: %v userID: %d gatewayID: %d countryID: %d", err, req.UserID, req.GatewayID, req.CountryID)
		log.Println(errMsg.Error())
		return nil, errMsg
	}

	return &contract.DepositResponse{
		TransactionID: id,
		Success:       true,
	}, nil
}
