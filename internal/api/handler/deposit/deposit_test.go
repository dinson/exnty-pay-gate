package deposit

import (
	"bytes"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"payment-gateway/constant"
	"payment-gateway/internal/services/gateway/contract"
	gatewayMocks "payment-gateway/internal/services/gateway/mocks"
	txnContract "payment-gateway/internal/services/transaction/contract"
	txnMocks "payment-gateway/internal/services/transaction/mocks"
	"testing"
)

var (
	getGatewayResp = []*contract.Gateway{
		{
			ID:   1,
			Name: "stripe",
		},
		{
			ID:   2,
			Name: "link",
		},
	}

	txnDepositResp = &txnContract.DepositResponse{
		TransactionID: 1,
		Success:       true,
	}

	txnDepositFailedResp = &txnContract.DepositResponse{
		TransactionID: 1,
		Success:       false,
	}
)

func TestHandler_InitDeposit(t *testing.T) {
	type fields struct {
		Gateway *gatewayMocks.Gateway
		Txn     *txnMocks.Transaction
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
	}{
		{
			name: "happy path",
			fields: fields{
				Gateway: mockGatewayService(true, getGatewayResp, nil),
				Txn:     mockTransactionDepositService(true, txnDepositResp, nil),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/deposit", bytes.NewBufferString(`{"amount":100,"currency":"aed"}`)),
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "invalid payload",
			fields: fields{
				Gateway: mockGatewayService(false, nil, nil),
				Txn:     mockTransactionDepositService(false, nil, nil),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/deposit", bytes.NewBufferString(`invalid json`)),
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "get gateways error",
			fields: fields{
				Gateway: mockGatewayService(true, nil, errors.New("failed")),
				Txn:     mockTransactionDepositService(false, txnDepositResp, nil),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/deposit", bytes.NewBufferString(`{"amount":100,"currency":"aed"}`)),
			},
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name: "no gateways found",
			fields: fields{
				Gateway: mockGatewayService(true, nil, nil),
				Txn:     mockTransactionDepositService(false, txnDepositResp, nil),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/deposit", bytes.NewBufferString(`{"amount":100,"currency":"aed"}`)),
			},
			wantStatusCode: http.StatusForbidden,
		},
		{
			name: "deposit error",
			fields: fields{
				Gateway: mockGatewayService(true, getGatewayResp, nil),
				Txn:     mockTransactionDepositService(true, nil, errors.New("failed")),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/deposit", bytes.NewBufferString(`{"amount":100,"currency":"aed"}`)),
			},
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name: "all deposit attempts failed",
			fields: fields{
				Gateway: mockGatewayService(true, getGatewayResp, nil),
				Txn:     mockTransactionDepositService(true, txnDepositFailedResp, nil),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/deposit", bytes.NewBufferString(`{"amount":100,"currency":"aed"}`)),
			},
			wantStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(tt.args.r.Context(), constant.UserID, 1)
			ctx = context.WithValue(ctx, constant.CountryID, 2)
			tt.args.r = tt.args.r.WithContext(ctx)

			h := Handler{
				Gateway: tt.fields.Gateway,
				Txn:     tt.fields.Txn,
			}

			rr := httptest.NewRecorder()
			h.InitDeposit(rr, tt.args.r)

			assert.Equal(t, tt.wantStatusCode, rr.Code)
			tt.fields.Txn.AssertExpectations(t)
			tt.fields.Gateway.AssertExpectations(t)
		})
	}
}

func mockGatewayService(enabled bool, resp []*contract.Gateway, err error) *gatewayMocks.Gateway {
	c := &gatewayMocks.Gateway{}

	if enabled {
		c.On("GetByCountry", mock.Anything, mock.Anything).Return(resp, err)
	}

	return c
}

func mockTransactionDepositService(enabled bool, resp *txnContract.DepositResponse, err error) *txnMocks.Transaction {
	c := &txnMocks.Transaction{}

	if enabled {
		c.On("Deposit", mock.Anything, mock.Anything).Return(resp, err)
	}

	return c
}
