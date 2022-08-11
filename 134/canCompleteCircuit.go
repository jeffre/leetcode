package leetcode134

/*
	Challenge Description: There are n gas stations along a circular route,
	where the amount of gas at the ith station is gas[i].

	You have a car with an unlimited gas tank and it costs cost[i] of gas to
	travel from the ith station to its next (i + 1)th station. You begin the
	journey with an empty tank at one of the gas stations.

	Given two integer arrays gas and cost, return the starting gas station's
	index if you can travel around the circuit once in the clockwise direction,
	otherwise return -1. If there exists a solution, it is guaranteed to be
	unique.
*/

func canCompleteCircuit(gas []int, cost []int) int {

	start, tank, total := 0, 0, 0

	for i := range gas {
		tank += gas[i] - cost[i]
		total += gas[i] - cost[i]

		// Reset the starting point after each failed traversal
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	// If the total amount of gas collected is greater than the amout needed,
	// then a solution was stumbled into using the above process.
	if total >= 0 {
		return start
	}
	return -1
}
