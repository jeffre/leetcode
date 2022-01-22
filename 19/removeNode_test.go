package leetcode19

import (
	"fmt"
	"reflect"
	"testing"
)

type given struct {
	head *ListNode
	n    int
}

var cases = []struct {
	given given
	want  *ListNode
}{
	{
		given: given{
			head: &ListNode{1,
				&ListNode{2,
					&ListNode{3,
						&ListNode{4,
							&ListNode{5, nil},
						},
					},
				},
			},
			n: 2,
		},
		want: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{5, nil},
				},
			},
		},
	},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := removeNthFromEnd(c.given.head, c.given.n)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			removeNthFromEnd(c.given.head, c.given.n)
		}
	}
}
