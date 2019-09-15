package anacondaaaa

import (
	"testing"
)

func TestCreateCRCToken(t *testing.T) {
	type args struct {
		crcToken       string
		consumerSecret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Create valid token",
			args: args{
				crcToken:       "some_token",
				consumerSecret: "some_secret_key",
			},
			want: "sha256=vSQ9tjVLlUSi3O5rXRcmUAT8BYqR3Shn+nJos1bLGqU=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateCRCToken(tt.args.crcToken, tt.args.consumerSecret); got != tt.want {
				t.Errorf("CreateCRCToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
