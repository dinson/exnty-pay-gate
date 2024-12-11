package transaction

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"payment-gateway/db"
	"payment-gateway/internal/services/transaction/contract"
	"testing"
)

func Test_impl_Withdraw(t *testing.T) {
	type fields struct {
		db db.DB
	}
	type args struct {
		ctx context.Context
		req *contract.WithdrawRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *contract.WithdrawResponse
		wantErr error
	}{
		{
			name: "happy path",
			fields: fields{
				db: mockDBCreateTxn(true, 1, nil),
			},
			args: args{
				ctx: context.Background(),
				req: &contract.WithdrawRequest{
					UserID:    1,
					Amount:    100,
					Currency:  "AED",
					GatewayID: 1,
					CountryID: 1,
				},
			},
			want: &contract.WithdrawResponse{
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
				req: &contract.WithdrawRequest{
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
			got, err := i.Withdraw(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
