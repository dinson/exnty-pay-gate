package transaction

import (
	"context"
	"fmt"
	"log"
	"payment-gateway/internal/services/transaction/contract"
)

func (i impl) UpdateStatus(ctx context.Context, req *contract.UpdateStatusRequest) error {
	txn, err := i.db.GetTransactionByID(ctx, req.TransactionID)
	if err != nil || txn == nil {
		errMsg := fmt.Errorf("failed to retrieve txn by id: %d err: %v", req.TransactionID, err)
		log.Println(errMsg)
		return errMsg
	}

	txn.Status = req.Status.String()

	if err = i.db.UpdateTransactionByID(ctx, req.TransactionID, txn); err != nil {
		errMsg := fmt.Errorf("failed to update txn by id: %d err: %v", req.TransactionID, err)
		log.Println(errMsg)
		return errMsg
	}

	return nil
}
