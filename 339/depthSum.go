package leetcode339

/*
	Challenge Description: You are given a nested list of integers nestedList.
	Each element is either an integer or a list whose elements may also be
	integers or other lists.

	The depth of an integer is the number of lists that it is inside of. For
	example, the nested list [1,[2,2],[[3],2],1] has each integer's value set to
	its depth.

	Return the sum of each integer in nestedList multiplied by its depth.
*/

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	i     int
	ni    []*NestedInteger
	isInt bool
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n *NestedInteger) IsInteger() bool {
	return n.isInt
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n *NestedInteger) GetInteger() int {
	return n.i
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.isInt = true
	n.i = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	n.ni = append(n.ni, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n *NestedInteger) GetList() []*NestedInteger {
	return n.ni
}

func depthSum(nestedList []*NestedInteger) int {

	q := nestedList // Repurpose the starting point as our queue
	depth, result := 1, 0

	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {

			// Pop the front off the queue
			ni := q[0]
			q = q[1:]

			// Add the integer or append its children
			if ni.IsInteger() {
				result += ni.GetInteger() * depth
			} else {
				q = append(q, ni.GetList()...)
			}
		}
		depth++
	}
	return result
}
