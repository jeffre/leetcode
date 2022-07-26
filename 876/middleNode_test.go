package leetcode876

import (
	"testing"
)

type given struct {
	head *ListNode
}

type test struct {
	name  string
	given given
	want  *ListNode
}

var tests = []test{
	{
		name: "example1",
		given: given{
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4,
				&ListNode{5, nil}}}}},
		},
		want: &ListNode{3, &ListNode{4, &ListNode{5, nil}}},
	}, {
		name: "example2",
		given: given{
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4,
				&ListNode{5, &ListNode{6, nil}}}}}},
		},
		want: &ListNode{4, &ListNode{5, &ListNode{6, nil}}},
	},
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := middleNode(tt.given.head)
			if !AssertIdentical(t, got, tt.want) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range tests {
			middleNode(tt.given.head)
		}
	}
}
