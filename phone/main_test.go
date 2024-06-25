package main

import "testing"

func TestSanitizePhoneNumber(t *testing.T) {
	type args struct {
		phoneNumber string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1: +62",
			args: args{
				phoneNumber: "+6281220058xxx",
			},
			want: "6281220058xxx",
		},
		{
			name: "case 2: 62",
			args: args{
				phoneNumber: "6281220058xxx",
			},
			want: "6281220058xxx",
		},
		{
			name: "case 3: 08",
			args: args{
				phoneNumber: "081220058xxx",
			},
			want: "6281220058xxx",
		},
		{
			name: "case 4: 8",
			args: args{
				phoneNumber: "81220058xxx",
			},
			want: "6281220058xxx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizePhoneNumber(tt.args.phoneNumber); got != tt.want {
				t.Errorf("SanitizePhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
