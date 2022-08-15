#!/bin/bash

if [ "$#" -ne 2 ]; then
  printf "$0 requires 2 parameters: (Problem Number) and (Function Name)\n" >&2
  exit 1
fi

id="$1"
func="$2"
go="${id}/${func}.go"
gotest="${id}/${func}_test.go"

mkdir "$id" || exit


## main .go file
test -f "$go" || cat <<EOF > "${go}"
package leetcode$id

/*
	Challenge Description: 
*/

func $func() {
	
}

EOF


## _test.go
test -f "$gotest" || cat <<EOF > "${gotest}"
package leetcode$id

import "testing"

type given struct {
	nums   []int
	target int
}

type want struct {
	int
}

func runTest(given given) want {
	got := ${func}(given.nums, given.target)
	return want{got}
}

var tests = []test {
	{
		"example1",
		given{
			nums:   []int{4, 5},
			target: 0,
		},
		want{4},
	},
}

type test struct {
	name  string
	given given
	want  want
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := runTest(tt.given)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			runTest(tt.given)
		}
	}
}

EOF


## Open new files for editing
code "$go" "$gotest"
