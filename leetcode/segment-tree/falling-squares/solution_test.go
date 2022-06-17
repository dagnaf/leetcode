package falling

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_fallingSquares(t *testing.T) {
	type args struct {
		positions [][]int
		s         string
	}
	tests := []struct {
		name string
		args args
		want []int
		s    string
	}{
		{
			args: args{s: `[[1,2],[2,3],[6,1]]`},
			s:    `[2,5,5]`,
		},
		{
			args: args{s: `[[100,100],[200,100]]`},
			s:    `[100,100]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			json.Unmarshal([]byte(tt.args.s), &tt.args.positions)
			json.Unmarshal([]byte(tt.s), &tt.want)
			if got := fallingSquares(tt.args.positions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fallingSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
掉落方块，求每次堆积起来的高度
线段树，节点含义，叶子节点[1]表示区间1~2，节点[2]表示2~3，节点[1,2]表示1~2,2~3
*/
