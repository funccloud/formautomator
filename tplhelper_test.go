package formautomator

import "testing"

func Test_in(t *testing.T) {
	type args struct {
		a []string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success true",
			args: args{
				a: []string{"a", "b", "c"},
				b: "b",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := in(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("in() = %v, want %v", got, tt.want)
			}
		})
	}
}
