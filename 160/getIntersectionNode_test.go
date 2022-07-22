package leetcode237

import (
	"testing"
)

type test struct {
	name  string
	given given
	want  *ListNode
}
type given struct {
	a *ListNode
	b *ListNode
}

func getCases() (tests []test) {

	// Example 1
	common1 := &ListNode{8, &ListNode{4, &ListNode{5, nil}}}
	a1 := &ListNode{4, &ListNode{1, common1}}
	b1 := &ListNode{5, &ListNode{6, &ListNode{1, common1}}}
	tests = append(tests, test{
		name:  "example1",
		given: given{a1, b1},
		want:  common1,
	})

	// Example 2
	common2 := &ListNode{2, &ListNode{4, nil}}
	a2 := &ListNode{1, &ListNode{9, &ListNode{1, common2}}}
	b2 := &ListNode{3, common2}
	tests = append(tests, test{
		name:  "example2",
		given: given{a2, b2},
		want:  common2,
	})

	// Example 3
	a3 := &ListNode{2, &ListNode{6, &ListNode{4, nil}}}
	b3 := &ListNode{1, &ListNode{5, nil}}
	tests = append(tests, test{
		name:  "example3",
		given: given{a3, b3},
		want:  nil,
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
			got := getIntersectionNode(tt.given.a, tt.given.b)
			if !AssertIdentical(t, got, tt.want) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", tt.given, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range getCases() {
			getIntersectionNode(tt.given.a, tt.given.b)
		}
	}
}
