package transaction

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"payment-gateway/db/mocks"
	"payment-gateway/internal/services/transaction/contract"
	"testing"
)

func Test_impl_Deposit(t *testing.T) {
	type fields struct {
		db *mocks.DB
	}
	type args struct {
		ctx context.Context
		req *contract.DepositRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *contract.DepositResponse
		wantErr error
	}{
		{
			name: "happy path",
			fields: fields{
				db: mockDBCreateTxn(true, 1, nil),
			},
			args: args{
				ctx: context.Background(),
				req: &contract.DepositRequest{
					UserID:    1,
					Amount:    100,
					Currency:  "AED",
					GatewayID: 1,
					CountryID: 1,
				},
			},
			want: &contract.DepositResponse{
				TransactionID: 1,
				Success:       true,
			},
			wantErr: nil,
		},
		{
			name: "repo error",
			fields: fields{
				db: mockDBCreateTxn(true, 0, errors.New("failed")),
			},
			args: args{
				ctx: context.Background(),
				req: &contract.DepositRequest{
					UserID:    1,
					Amount:    100,
					Currency:  "AED",
					GatewayID: 1,
					CountryID: 1,
				},
			},
			want:    nil,
			wantErr: errors.New("failed to save to txn table: failed userID: 1 gatewayID: 1 countryID: 1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := impl{
				db: tt.fields.db,
			}
			got, err := i.Deposit(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func mockDBCreateTxn(enabled bool, respID int, err error) *mocks.DB {
	c := &mocks.DB{}
	if enabled {
		c.On("CreateTransaction", mock.Anything, mock.Anything).Return(respID, err)
	}
	return c
}
