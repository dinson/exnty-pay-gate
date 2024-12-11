package gateway

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"payment-gateway/db"
	"payment-gateway/db/mocks"
	"payment-gateway/internal/services/gateway/contract"
	"testing"
)

var (
	countryGatewayModels = []*db.GatewaysForCountry{
		{
			ID:   1,
			Name: "stripe",
		},
		{
			ID:   2,
			Name: "link",
		},
		{
			ID:   3,
			Name: "razorpay",
		},
	}

	listCountryGatewayResp = []*contract.Gateway{
		{
			ID:   1,
			Name: "stripe",
		},
		{
			ID:   2,
			Name: "link",
		},
		{
			ID:   3,
			Name: "razorpay",
		},
	}
)

func Test_impl_GetByCountry(t *testing.T) {
	type fields struct {
		db *mocks.DB
	}
	type args struct {
		ctx context.Context
		req *contract.GetGatewayByCountryRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*contract.Gateway
		wantErr error
	}{
		{
			name: "happy path",
			fields: fields{
				db: mockListCountryGateways(countryGatewayModels, nil),
			},
			args: args{
				ctx: context.Background(),
				req: &contract.GetGatewayByCountryRequest{
					CountryID: 1,
				},
			},
			want:    listCountryGatewayResp,
			wantErr: nil,
		},
		{
			name: "empty result",
			fields: fields{
				db: mockListCountryGateways(nil, nil),
			},
			args: args{
				ctx: context.Background(),
				req: &contract.GetGatewayByCountryRequest{
					CountryID: 1,
				},
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "repo error",
			fields: fields{
				db: mockListCountryGateways(nil, errors.New("failed")),
			},
			args: args{
				ctx: context.Background(),
				req: &contract.GetGatewayByCountryRequest{
					CountryID: 1,
				},
			},
			want:    nil,
			wantErr: errors.New("failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := impl{
				db: tt.fields.db,
			}
			got, err := i.GetByCountry(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
			tt.fields.db.AssertExpectations(t)
		})
	}
}

func mockListCountryGateways(resp []*db.GatewaysForCountry, err error) *mocks.DB {
	c := &mocks.DB{}
	c.On("ListCountryGatewaysByPriority", mock.Anything, mock.Anything).Return(resp, err)
	return c
}
