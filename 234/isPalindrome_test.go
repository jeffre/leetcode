package leetcode141

import (
	"testing"
)

var tests = []struct {
	name  string
	given *ListNode
	want  bool
}{
	{
		name: "example1",
		given: &ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 1},
				},
			},
		},
		want: true,
	}, {
		name: "example2",
		given: &ListNode{Val: 1,
			Next: &ListNode{Val: 2},
		},
		want: false,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.given)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range tests {
			isPalindrome(tt.given)
		}
	}
}
