package leetcode101

import "testing"

type given struct {
	root *TreeNode
}

var cases = []struct {
	name  string
	given given
	want  bool
}{
	{
		name: "example1",
		given: given{
			root: &TreeNode{1,
				&TreeNode{2,
					&TreeNode{3, nil, nil},
					&TreeNode{4, nil, nil},
				},
				&TreeNode{2,
					&TreeNode{4, nil, nil},
					&TreeNode{3, nil, nil},
				},
			},
		},
		want: true,
	}, {
		name: "example2",
		given: given{
			root: &TreeNode{1,
				&TreeNode{2,
					nil,
					&TreeNode{3, nil, nil},
				},
				&TreeNode{2,
					nil,
					&TreeNode{3, nil, nil},
				},
			},
		},
		want: false,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := isSymmetric(tt.given.root)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			isSymmetric(tt.given.root)
		}
	}
}
