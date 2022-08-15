package leetcode116

import "testing"

type given struct {
	root *Node
}

type result struct {
	*Node
}

func (tt *test) run() result {
	got := connect(tt.given.root)
	return result{got}
}

var tests []test

// Add example1
func init() {

	common := func() *Node {
		return &Node{1,
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
		}
	}

	g := given{common()}
	w := result{common()}

	w.Left.Next = w.Right
	w.Left.Left.Next = w.Left.Right
	w.Left.Right.Next = w.Right.Left
	w.Right.Left.Next = w.Right.Right

	tests = append(tests, test{"example1", g, w})
}

// Add example2
func init() {
	tests = append(tests, test{
		"example2",
		given{},
		result{},
	})
}

type test struct {
	name  string
	given given
	want  result
}

func TestCases(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.run()
			if !assertIdentical(t, got.Node, tt.want.Node) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range tests {
			tt.run()
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
