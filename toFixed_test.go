package toFixed

import "testing"

func TestToFixed(t *testing.T) {
	type args struct {
		n         float64
		precision int
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		// Test cases.
		{
			name: "routine",
			args: args{1.2345678, 2},
			want: 1.23,
		},
		{
			name: "round",
			args: args{1.2345678, 5},
			want: 1.23457,
		},
		{
			name: "out of range",
			args: args{1.278, 5},
			want: 1.278,
		},
		{
			name: "integer",
			args: args{92234, 4},
			want: 92234,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFixed(tt.args.n, tt.args.precision); got != tt.want {
				t.Errorf("ToFixed() = %v, want %v", got, tt.want)
			}
		})
	}
}
