package leetcode21

import (
	"reflect"
	"testing"
)

func TestListNodeToInts(t *testing.T) {
	given := []int{0, 1, 2}
	got := ListNodeToInts(IntsToListNode(given))
	if !reflect.DeepEqual(given, got) {
		t.Errorf("Wanted %+v but got %+v", given, got)
	}
}
