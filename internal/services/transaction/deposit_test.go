package transaction

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"payment-gateway/db/mocks"
	"payment-gateway/internal/services/transaction/contract"
	ppMocks "payment-gateway/paymentprovider/mocks"
	"testing"
)

func Test_impl_Deposit(t *testing.T) {
	type fields struct {
		db              *mocks.DB
		paymentProvider *ppMocks.PaymentProvider
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
				db:              mockDBCreateTxn(true, 1, nil),
				paymentProvider: mockPaymentProvider(true, nil, false, nil),
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
			name: "payment provider error",
			fields: fields{
				db:              mockDBCreateTxn(false, 1, nil),
				paymentProvider: mockPaymentProvider(true, errors.New("failed"), false, nil),
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
				TransactionID: 0,
				Success:       false,
			},
			wantErr: nil,
		},
		{
			name: "repo error",
			fields: fields{
				db:              mockDBCreateTxn(true, 0, errors.New("failed")),
				paymentProvider: mockPaymentProvider(true, nil, false, nil),
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
				db:              tt.fields.db,
				paymentProvider: tt.fields.paymentProvider,
			}
			got, err := i.Deposit(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
			tt.fields.db.AssertExpectations(t)
			tt.fields.paymentProvider.AssertExpectations(t)
		})
	}
}

func mockPaymentProvider(enabledDeposit bool, depErr error, enabledWithdraw bool, withdrawErr error) *ppMocks.PaymentProvider {
	c := &ppMocks.PaymentProvider{}

	if enabledDeposit {
		c.On("Deposit", mock.Anything, mock.Anything).Return(depErr)
	}

	if enabledWithdraw {
		c.On("Withdraw", mock.Anything, mock.Anything).Return(withdrawErr)
	}

	return c
}

func mockDBCreateTxn(enabled bool, respID int, err error) *mocks.DB {
	c := &mocks.DB{}
	if enabled {
		c.On("CreateTransaction", mock.Anything, mock.Anything).Return(respID, err)
	}
	return c
}
