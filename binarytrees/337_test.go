package binarytrees

import (
	"reflect"
	"testing"
)

func TestRob1(t *testing.T) {
	root := new(TreeNode)
	root.Val = 3
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = nil
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 3
	root.Right = new(TreeNode)
	root.Right.Val = 3
	root.Right.Left = nil
	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 1

	got := rob(root)
	want := 7
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
