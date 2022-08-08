package leetcode102

import (
	"reflect"
	"testing"
)

type given struct {
	root *TreeNode
}

var cases = []struct {
	name  string
	given given
	want  [][]int
}{
	{
		name: "example1",
		given: given{
			root: &TreeNode{3,
				&TreeNode{9,
					nil,
					nil,
				},
				&TreeNode{20,
					&TreeNode{15, nil, nil},
					&TreeNode{7, nil, nil},
				},
			},
		},
		want: [][]int{{3}, {9, 20}, {15, 7}},
	}, {
		name: "example2",
		given: given{
			root: &TreeNode{1, nil, nil},
		},
		want: [][]int{{1}},
	}, {
		name: "example3",
		given: given{
			root: nil,
		},
		want: [][]int(nil),
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(tt.given.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			levelOrder(tt.given.root)
		}
	}
}
