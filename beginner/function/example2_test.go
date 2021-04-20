package function

import "testing"

func TestMulti(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO 30: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multi(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Multi() = %v, want %v", got, tt.want)
			}
		})
	}
}
