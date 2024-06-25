package main

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		wantC int
	}{
		// TODO: Add test cases.
		{
			name: "Case 1: Correct case",
			args: args{
				a: 1,
				b: 1,
			},
			wantC: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := Sum(tt.args.a, tt.args.b); gotC != tt.wantC {
				t.Errorf("Sum() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
