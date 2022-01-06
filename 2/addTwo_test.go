package leetcode2

import (
	"fmt"
	"reflect"
	"testing"
)

type given struct {
	l1 *ListNode
	l2 *ListNode
}

var cases = []struct {
	given given
	want  *ListNode
}{
	{
		given: given{
			l1: &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			l2: &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
		},
		want: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
	}, {
		given: given{
			l1: &ListNode{0, &ListNode{}},
			l2: &ListNode{0, &ListNode{}},
		},
		want: &ListNode{0, &ListNode{}},
	}, {
		given: given{
			l1: &ListNode{9, &ListNode{}},
			l2: &ListNode{1, &ListNode{}},
		},
		want: &ListNode{0, &ListNode{1, nil}},
	}, {
		given: given{
			l1: &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{}}}}}}}},
			l2: &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{}}}}},
		},
		want: &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}},
	},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := addTwoNumbers(c.given.l1, c.given.l2)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}

		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			addTwoNumbers(c.given.l1, c.given.l2)
		}
	}
}
