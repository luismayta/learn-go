package beginner

import "testing"

func Test_newRelease(t *testing.T) {
	type args struct {
		a *Artist
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test basic",
			args: args{
				&Artist{
					Name:  "Luis",
					Gnre:  "Luis",
					Songs: 10,
				},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRelease(tt.args.a); got != tt.want {
				t.Errorf("newRelease() = %v, want %v", got, tt.want)
			}
		})
	}
}
