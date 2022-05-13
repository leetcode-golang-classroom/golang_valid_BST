package sol

import (
	"strconv"
	"testing"
)

func CreateBinaryTree(input *[]string) *TreeNode {
	var tree *TreeNode
	arr := *input
	result := make([]*TreeNode, len(arr))
	for idx, val := range arr {
		if val != "null" {
			num, _ := strconv.Atoi(val)
			result[idx] = &TreeNode{Val: num}
		}
	}
	tree = result[0]
	for idx, node := range result {
		if 2*idx+1 < len(result) {
			node.Left = result[2*idx+1]
		}
		if 2*idx+2 < len(result) {
			node.Right = result[2*idx+2]
		}
	}
	return tree
}
func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "root = [2,1,3]",
			args: args{CreateBinaryTree(&[]string{"2", "1", "3"})},
			want: true,
		},
		{
			name: "root = [5,1,4,null,null,3,6]",
			args: args{CreateBinaryTree(&[]string{"5", "1", "4", "null", "null", "3", "6"})},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
