package leetcode704

import (
	"math"
	"testing"
)

type given struct {
	nums   []int
	target int
}

var cases = []struct {
	name  string
	given given
	want  int
}{
	{
		name: "example1",
		given: given{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
		},
		want: 4,
	}, {
		name: "example2",
		given: given{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
		},
		want: -1,
	}, {
		name: "test1",
		given: given{
			nums:   []int{5},
			target: 5,
		},
		want: 0,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.given.nums, tt.given.target)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			search(tt.given.nums, tt.given.target)
		}
	}
}

func TestOverflow(t *testing.T) {
	high := math.MaxInt
	low := high - 2
	want := high - 1

	//got := (high + low) / 2 //overflowed
	got := low + (high-low)/2 //0.2766 ns/op

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkOverflow(b *testing.B) {
	high := math.MaxInt
	low := high - 2
	want := high - 1

	for n := 0; n < b.N; n++ {
		got := low + (high-low)/2
		if got != want {
			b.Errorf("got %v want %v", got, want)
		}
	}
}
