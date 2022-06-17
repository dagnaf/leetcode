package count

import (
	"encoding/json"
	"strings"
	"testing"
)

func Test_countRangeSum(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
		s     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: `[-2,5,-1]
-2
2`,
			},
			want: 3,
		},
		{
			args: args{
				s: `[0]
0
0`,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := strings.Split(tt.args.s, "\n")
			json.Unmarshal([]byte(s[0]), &tt.args.nums)
			json.Unmarshal([]byte(s[1]), &tt.args.lower)
			json.Unmarshal([]byte(s[2]), &tt.args.upper)
			if got := countRangeSum(tt.args.nums, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("countRangeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
给定数组a，求子数组数量，满足求和在给定范围lower到upper内
求前缀和si=a1+...+ai
i-j区间和为sj-si在l-r之间，那si就在sj-r到sj-l之间
求这样的si的个数
线段树维护前缀和si的频数直方图
*/
