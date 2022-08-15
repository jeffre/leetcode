package leetcode116

import "testing"

type given struct {
	root *Node
}

type want struct {
	*Node
}

func runTest(given given) want {
	got := connect(given.root)
	return want{got}
}

var tests []test

// Add example1
func init() {
	g := given{
		&Node{1,
			&Node{2,
				&Node{4, nil, nil, nil},
				&Node{5, nil, nil, nil},
				nil,
			},
			&Node{3,
				&Node{6, nil, nil, nil},
				&Node{7, nil, nil, nil},
				nil,
			},
			nil,
		},
	}

	w := want{
		&Node{1,
			&Node{2,
				&Node{4, nil, nil, nil},
				&Node{5, nil, nil, nil},
				nil,
			},
			&Node{3,
				&Node{6, nil, nil, nil},
				&Node{7, nil, nil, nil},
				nil,
			},
			nil,
		},
	}
	w.Left.Next = w.Right
	w.Left.Left.Next = w.Left.Right
	w.Left.Right.Next = w.Right.Left
	w.Right.Left.Next = w.Right.Right

	tests = append(tests, test{"example1", g, w})
}

type test struct {
	name  string
	given given
	want  want
}

func TestCases(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runTest(tt.given)
			if !assertIdentical(t, got.Node, tt.want.Node) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range tests {
			runTest(tt.given)
		}
	}
}

func assertIdentical(t testing.TB, a, b *Node) bool {
	t.Helper()

	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else if a.Val != b.Val {
		return false
	}

	if (a.Next == nil) != (b.Next == nil) {
		return false
	}
	if (a.Left == nil) != (b.Left == nil) {
		return false
	}
	if (a.Right == nil) != (b.Right == nil) {
		return false
	}

	if a.Next != nil && a.Next.Val != b.Next.Val {
		return false
	}

	return assertIdentical(t, a.Left, b.Left) && assertIdentical(t, a.Right, b.Right)
}
