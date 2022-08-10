package leetcode339

import "testing"

type given struct {
	nestedList []*NestedInteger
}

type testCase struct {
	name  string
	given given
	want  int
}

var cases []testCase

func init() {

	nl11 := NestedInteger{}
	nl11.SetInteger(1)
	nl12 := NestedInteger{}
	nl12.SetInteger(1)

	nl31 := NestedInteger{}
	nl31.SetInteger(1)
	nl32 := NestedInteger{}
	nl32.SetInteger(1)

	nl1 := NestedInteger{}
	nl1.Add(nl11)
	nl1.Add(nl12)

	nl2 := NestedInteger{}
	nl2.SetInteger(2)

	nl3 := NestedInteger{}
	nl3.Add(nl31)
	nl3.Add(nl32)

	nl := []*NestedInteger{&nl1, &nl2, &nl3}

	cases = append(cases, testCase{
		name:  "example1",
		given: given{nl},
		want:  10,
	})
}

func init() {

	nl1 := NestedInteger{}
	nl1.SetInteger(1)

	nl21 := NestedInteger{}
	nl21.SetInteger(4)

	nl221 := NestedInteger{}
	nl221.SetInteger(6)

	nl22 := NestedInteger{}
	nl22.Add(nl221)

	nl2 := NestedInteger{}
	nl2.Add(nl21)
	nl2.Add(nl22)

	nl := []*NestedInteger{&nl1, &nl2}

	cases = append(cases, testCase{
		name:  "example2",
		given: given{nl},
		want:  27,
	})
}

func init() {

	nl1 := NestedInteger{}
	nl1.SetInteger(0)

	nl := []*NestedInteger{&nl1}

	cases = append(cases, testCase{
		name:  "example3",
		given: given{nl},
		want:  0,
	})
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := depthSum(tt.given.nestedList)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			depthSum(tt.given.nestedList)
		}
	}
}
