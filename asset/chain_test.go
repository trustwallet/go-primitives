package asset

import (
	"reflect"
	"testing"
)

func TestGetChainFromAssetType(t *testing.T) {
	type args struct {
		type_ string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test ERC20",
			args: args{
				type_: "ERC20",
			},
			want:    "ethereum",
			wantErr: false,
		},
		{
			name: "Test custom chain type",
			args: args{
				type_: "CSTM20",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChainFromAssetType(tt.args.type_)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoinForId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoinForId() got = %v, want %v", got, tt.want)
			}
		})
	}
}
