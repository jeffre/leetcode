package leetcode104

import "testing"

type given struct {
	root *TreeNode
}

var cases = []struct {
	name  string
	given given
	want  int
}{
	{
		name: "example1",
		given: given{
			root: &TreeNode{3,
				&TreeNode{9, nil, nil},
				&TreeNode{20,
					&TreeNode{15, nil, nil},
					&TreeNode{7, nil, nil},
				},
			},
		},
		want: 3,
	}, {
		name: "example2",
		given: given{
			root: &TreeNode{1,
				nil,
				&TreeNode{2, nil, nil},
			},
		},
		want: 2,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDepth(tt.given.root)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			maxDepth(tt.given.root)
		}
	}
}
