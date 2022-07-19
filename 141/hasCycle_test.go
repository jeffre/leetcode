package leetcode141

import (
	"testing"
)

type test struct {
	name  string
	given *ListNode
	want  bool
}

func testCases(tb testing.TB) (tests []test) {
	tb.Helper()
	cyc1 := &ListNode{Val: 2,
		Next: &ListNode{Val: 0,
			Next: &ListNode{Val: 4},
		},
	}
	cyc1.Next.Next.Next = cyc1
	tests = append(tests, test{
		name:  "example1-cyclic",
		given: &ListNode{Val: 3, Next: cyc1},
		want:  true,
	})

	cyc2 := &ListNode{Val: 1,
		Next: &ListNode{Val: 2},
	}
	cyc2.Next.Next = cyc2
	tests = append(tests, test{
		name:  "example2-cyclic",
		given: cyc2,
		want:  true,
	})

	cyc3 := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5,
						Next: &ListNode{Val: 6,
							Next: &ListNode{Val: 7,
								Next: &ListNode{Val: 8,
									Next: &ListNode{Val: 9},
								},
							},
						},
					},
				},
			},
		},
	}
	cyc3.Next.Next.Next.Next.Next.Next.Next.Next.Next = cyc3
	tests = append(tests, test{
		name:  "long-cycle",
		given: cyc3,
		want:  true,
	})

	tests = append(tests, test{
		name: "non-cyclic",
		given: &ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4},
				},
			},
		},
		want: false,
	})

	tests = append(tests, test{
		name:  "empty-list",
		given: nil,
		want:  false,
	})

	return
}

func TestCases(t *testing.T) {
	for _, tt := range testCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			got := hasCycle(tt.given)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range testCases(b) {
			hasCycle(tt.given)
		}
	}
}
