package tests

import (
	"github.com/danielMensah/go-lambda-example/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomer_GetCustomers(t *testing.T) {
	type args struct {
		store string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "missing store",
			args: args{
				store: "",
			},
			wantErr: true,
		},
		{
			name: "provided store",
			args: args{
				store: "abc",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		repo := repository.NewRepository()

		_, err := repo.CustomerRepo.GetCustomers(tt.args.store)

		assert.True(t, (err != nil) == tt.wantErr)
	}
}
