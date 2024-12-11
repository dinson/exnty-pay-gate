package callback

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"payment-gateway/internal/services/transaction/mocks"
	"testing"
)

func TestHandler_HandleDepositSuccess(t *testing.T) {
	type fields struct {
		Txn *mocks.Transaction
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	var tests = []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
	}{
		{
			name: "happy path",
			fields: fields{
				Txn: mockTransactionService(true, nil),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/gateway/callback/deposit/success", bytes.NewBufferString(`{"transaction_id":1}`)),
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "invalid payload",
			fields: fields{
				Txn: mockTransactionService(false, nil),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/gateway/callback/deposit/success", bytes.NewBufferString(`invalid`)),
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "txn service error",
			fields: fields{
				Txn: mockTransactionService(true, errors.New("failed")),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/gateway/callback/deposit/success", bytes.NewBufferString(`{"transaction_id":1}`)),
			},
			wantStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.r = tt.args.r.WithContext(tt.args.r.Context())
			h := Handler{
				Txn: tt.fields.Txn,
			}
			rr := httptest.NewRecorder()
			h.HandleDepositSuccess(rr, tt.args.r)

			assert.Equal(t, tt.wantStatusCode, rr.Code)
			tt.fields.Txn.AssertExpectations(t)
		})
	}
}

func mockTransactionService(enabled bool, err error) *mocks.Transaction {
	c := &mocks.Transaction{}
	if enabled {
		c.On("UpdateStatus", mock.Anything, mock.Anything).Return(err)
	}
	return c
}

func TestHandler_HandleDepositFailure(t *testing.T) {
	type fields struct {
		Txn *mocks.Transaction
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	var tests = []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
	}{
		{
			name: "happy path",
			fields: fields{
				Txn: mockTransactionService(true, nil),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/gateway/callback/deposit/failure", bytes.NewBufferString(`{"transaction_id":1, "reason":""}`)),
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "invalid payload",
			fields: fields{
				Txn: mockTransactionService(false, nil),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/gateway/callback/deposit/failure", bytes.NewBufferString(`invalid`)),
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "txn service error",
			fields: fields{
				Txn: mockTransactionService(true, errors.New("failed")),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/gateway/callback/deposit/failure", bytes.NewBufferString(`{"transaction_id":1}`)),
			},
			wantStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.r = tt.args.r.WithContext(tt.args.r.Context())
			h := Handler{
				Txn: tt.fields.Txn,
			}
			rr := httptest.NewRecorder()
			h.HandleDepositFailure(rr, tt.args.r)

			assert.Equal(t, tt.wantStatusCode, rr.Code)
			tt.fields.Txn.AssertExpectations(t)
		})
	}
}

func TestHandler_HandleWithdrawalSuccess(t *testing.T) {
	type fields struct {
		Txn *mocks.Transaction
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	var tests = []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
	}{
		{
			name: "happy path",
			fields: fields{
				Txn: mockTransactionService(true, nil),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/gateway/callback/withdrawal/success", bytes.NewBufferString(`{"transaction_id":1}`)),
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "invalid payload",
			fields: fields{
				Txn: mockTransactionService(false, nil),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/gateway/callback/withdrawal/success", bytes.NewBufferString(`invalid`)),
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "txn service error",
			fields: fields{
				Txn: mockTransactionService(true, errors.New("failed")),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/gateway/callback/withdrawal/success", bytes.NewBufferString(`{"transaction_id":1}`)),
			},
			wantStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.r = tt.args.r.WithContext(tt.args.r.Context())
			h := Handler{
				Txn: tt.fields.Txn,
			}
			rr := httptest.NewRecorder()
			h.HandleWithdrawalSuccess(rr, tt.args.r)

			assert.Equal(t, tt.wantStatusCode, rr.Code)
			tt.fields.Txn.AssertExpectations(t)
		})
	}
}

func TestHandler_HandleWithdrawalFailure(t *testing.T) {
	type fields struct {
		Txn *mocks.Transaction
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	var tests = []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
	}{
		{
			name: "happy path",
			fields: fields{
				Txn: mockTransactionService(true, nil),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/gateway/callback/withdrawal/failure", bytes.NewBufferString(`{"transaction_id":1,"reason":""}`)),
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "invalid payload",
			fields: fields{
				Txn: mockTransactionService(false, nil),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/gateway/callback/withdrawal/failure", bytes.NewBufferString(`invalid`)),
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "txn service error",
			fields: fields{
				Txn: mockTransactionService(true, errors.New("failed")),
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("POST", "/gateway/callback/withdrawal/failure", bytes.NewBufferString(`{"transaction_id":1}`)),
			},
			wantStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.r = tt.args.r.WithContext(tt.args.r.Context())
			h := Handler{
				Txn: tt.fields.Txn,
			}
			rr := httptest.NewRecorder()
			h.HandleWithdrawalFailure(rr, tt.args.r)

			assert.Equal(t, tt.wantStatusCode, rr.Code)
			tt.fields.Txn.AssertExpectations(t)
		})
	}
}
