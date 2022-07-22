package leetcode237

import (
	"testing"
)

type test struct {
	name  string
	head  *ListNode
	given *ListNode
	want  *ListNode
}

func getCases() (tests []test) {

	// Example 1
	head1 := &ListNode{Val: 4,
		Next: &ListNode{Val: 5,
			Next: &ListNode{Val: 1,
				Next: &ListNode{Val: 9},
			},
		},
	}
	want1 := &ListNode{Val: 4,
		Next: &ListNode{Val: 1,
			Next: &ListNode{Val: 9},
		},
	}
	tests = append(tests, test{
		name:  "example1",
		head:  head1,
		given: head1.Next,
		want:  want1,
	})

	// Example 2
	given2 := &ListNode{Val: 4,
		Next: &ListNode{Val: 5,
			Next: &ListNode{Val: 1,
				Next: &ListNode{Val: 9},
			},
		},
	}
	want2 := &ListNode{Val: 4,
		Next: &ListNode{Val: 5,
			Next: &ListNode{Val: 9},
		},
	}
	tests = append(tests, test{
		name:  "example2",
		head:  given2,
		given: given2.Next.Next,
		want:  want2,
	})

	return tests
}

func AssertIdentical(t testing.TB, a, b *ListNode) bool {
	t.Helper()
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

func TestCases(t *testing.T) {
	for _, tt := range getCases() {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.given)
			if !AssertIdentical(t, tt.head, tt.want) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", tt.given, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range getCases() {
			deleteNode(tt.given)
		}
	}
}
