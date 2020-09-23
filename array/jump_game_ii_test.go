package array

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{[]int{2,3,1,1,4}},
			want: 2,
		},
		{
			name: "2",
			args: args{[]int{5,1,2,3,1,1,1,4}},
			want: 3,
		},
		{
			name: "3",
			args: args{[]int{5}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
