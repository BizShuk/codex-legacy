package i

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAmazingTalker_OA1(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{input: 20}, want: "26"},
		{args: args{input: -7}, want: "-10"},
		{args: args{input: 0}, want: "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AmazingTalker_OA1(tt.args.input); got != tt.want {
				t.Errorf("AmazingTalker_OA1() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(tt.want)
			}
		})
	}
}

func TestAmazingTalker_OA2(t *testing.T) {
	type args struct {
		list []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{args: args{list: []int{1, 2, 2, 3, 4}, k: 5}, want: [][]int{{1, 3}, {2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AmazingTalker_OA2(tt.args.list, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AmazingTalker_OA2() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got)
			}
		})
	}
}
