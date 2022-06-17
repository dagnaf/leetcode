package online

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		cmd  []string
		arg  [][]int
		scmd string
		sarg string
	}
	tests := []struct {
		name  string
		args  args
		want  []interface{}
		swant string
	}{
		{
			args: args{
				scmd: `["MajorityChecker", "query", "query", "query"]`,
				sarg: `[[[1, 1, 2, 2, 1, 1]], [0, 5, 4], [0, 3, 3], [2, 3, 2]]`,
			},
			swant: `[null, 1, -1, 2]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			json.Unmarshal([]byte(tt.args.scmd), &tt.args.cmd)
			i := strings.Index(tt.args.sarg, ", [")
			json.Unmarshal([]byte("["+tt.args.sarg[i+2:]), &tt.args.arg)
			tt.args.arg = append([][]int{{}}, tt.args.arg...)
			json.Unmarshal([]byte(tt.args.sarg[2:i-1]), &tt.args.arg[0])
			json.Unmarshal([]byte(tt.swant), &tt.want)
			if got := run(tt.args.cmd, tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
