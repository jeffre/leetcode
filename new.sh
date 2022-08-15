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
package leetcode${id}

import "testing"

type given struct {
	nums   []int
	target int
}

type result struct {
	int
}

func (tt *test) run() result {
	got := ${func}(given.nums, given.target)
	return result{got}
}

var tests []test

// Add example1
func init() {
	g := given{nums: []int{4, 5}, target: 0}
	w := 4
	tests = append(tests, test{
		name: "example1",
		given: g,
		want: w,
	})
}

type test struct {
	name  string
	given given
	want  result
}

func TestCases(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.run()
			if got != tt.want {
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

EOF


## Open new files for editing
code "$go" "$gotest"
