package leetcode94

import (
	"reflect"
	"testing"
)

var tests []test

type test struct {
	name  string
	given *TreeNode
	want  result
}

type result []int

func (t *test) run() result {
	return inorderTraversal(t.given)
}

// Example1
func init() {
	root := &TreeNode{1,
		nil,
		&TreeNode{2,
			&TreeNode{3, nil, nil},
			nil,
		},
	}
	tests = append(tests, test{
		name:  "example1",
		given: root,
		want:  []int{1, 3, 2},
	})
}

// Example2
func init() {
	tests = append(tests, test{
		name:  "example2",
		given: nil,
		want:  []int{},
	})
}

// Example3
func init() {
	root := &TreeNode{1, nil, nil}
	tests = append(tests, test{
		name:  "example3",
		given: root,
		want:  []int{1},
	})
}

func TestCases(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.run()
			if !reflect.DeepEqual(got, tt.want) {
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
