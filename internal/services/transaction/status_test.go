package transaction

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"payment-gateway/db"
	"payment-gateway/db/mocks"
	"payment-gateway/enum"
	"payment-gateway/internal/services/transaction/contract"
	"testing"
	"time"
)

var (
	getTxnResp = &db.Transaction{
		ID:        1,
		Amount:    100,
		Type:      enum.TxnDeposit.String(),
		Status:    "initialized",
		UserID:    1,
		GatewayID: 1,
		CountryID: 1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
)

func Test_impl_UpdateStatus(t *testing.T) {
	type fields struct {
		db *mocks.DB
	}
	type args struct {
		ctx context.Context
		req *contract.UpdateStatusRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "happy path",
			fields: fields{
				db: mockDB(true, getTxnResp, nil, true, nil),
			},
			args: args{
				ctx: context.Background(),
				req: &contract.UpdateStatusRequest{
					TransactionID: 1,
					Status:        enum.TxnStatusSuccess,
				},
			},
			wantErr: nil,
		},
		{
			name: "get txn error",
			fields: fields{
				db: mockDB(true, nil, errors.New("failed"), false, nil),
			},
			args: args{
				ctx: context.Background(),
				req: &contract.UpdateStatusRequest{
					TransactionID: 1,
					Status:        enum.TxnStatusSuccess,
				},
			},
			wantErr: errors.New("failed to retrieve txn by id: 1 err: failed"),
		},
		{
			name: "missing txn",
			fields: fields{
				db: mockDB(true, nil, nil, false, nil),
			},
			args: args{
				ctx: context.Background(),
				req: &contract.UpdateStatusRequest{
					TransactionID: 1,
					Status:        enum.TxnStatusSuccess,
				},
			},
			wantErr: errors.New("failed to retrieve txn by id: 1 err: <nil>"),
		},
		{
			name: "error updating status",
			fields: fields{
				db: mockDB(true, getTxnResp, nil, true, errors.New("failed")),
			},
			args: args{
				ctx: context.Background(),
				req: &contract.UpdateStatusRequest{
					TransactionID: 1,
					Status:        enum.TxnStatusSuccess,
				},
			},
			wantErr: errors.New("failed to update txn by id: 1 err: failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := impl{
				db: tt.fields.db,
			}
			assert.Equal(t, tt.wantErr, i.UpdateStatus(tt.args.ctx, tt.args.req))
		})
	}
}

func mockDB(enabledGet bool, getResp *db.Transaction, getErr error, enabledUpdate bool, updateErr error) *mocks.DB {
	c := &mocks.DB{}

	if enabledGet {
		c.On("GetTransactionByID", mock.Anything, mock.Anything).Return(getResp, getErr)
	}

	if enabledUpdate {
		c.On("UpdateTransactionByID", mock.Anything, mock.Anything, mock.Anything).Return(updateErr)
	}

	return c
}
