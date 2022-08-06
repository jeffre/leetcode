package leetcode100

import "testing"

type given struct {
	p *TreeNode
	q *TreeNode
}

var cases = []struct {
	name  string
	given given
	want  bool
}{
	{
		name: "example1",
		given: given{
			p: &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
			q: &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
		},
		want: true,
	}, {
		name: "example2",
		given: given{
			p: &TreeNode{1,
				&TreeNode{2, nil, nil},
				nil,
			},
			q: &TreeNode{1,
				nil,
				&TreeNode{2, nil, nil},
			},
		},
		want: false,
	}, {
		name: "example3",
		given: given{
			p: &TreeNode{1,
				&TreeNode{2, nil, nil},
				&TreeNode{1, nil, nil},
			},
			q: &TreeNode{1,
				&TreeNode{1, nil, nil},
				&TreeNode{2, nil, nil},
			},
		},
		want: false,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := isSameTree(tt.given.p, tt.given.q)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			isSameTree(tt.given.p, tt.given.q)
		}
	}
}
