package main

import (
	"github.com/danielMensah/go-lambda-example/internal/utility/aws"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCustomers(t *testing.T) {
	type args struct {
		req aws.Request
	}
	tests := []struct {
		name string
		args args
		want aws.Response
	}{
		{
			name: "returns 200",
			args: args{
				req: aws.Request{
					Body: `{"store":"abc"}`,
				},
			},
			want: aws.Response{
				StatusCode: 200,
			},
		},
		{
			name: "returns 400",
			args: args{
				req: aws.Request{
					Body: `{}`,
				},
			},
			want: aws.Response{
				StatusCode: 400,
			},
		},
	}
	for _, tt := range tests {
		got := GetCustomers(tt.args.req)

		assert.Equal(t, tt.want.StatusCode, got.StatusCode)
	}
}
