package twobucket

import "fmt"

type State struct {
	SizeBucketOne, SizeBucketTwo int
	ParentsState                 []State
}

func (s State) String() string {
	return fmt.Sprintf("(%d,%d)", s.SizeBucketOne, s.SizeBucketTwo)
}

func (s State) IsGoal(goalAmount int) (goalBucket int, otherBucketLevel int, err error) {
	if s.SizeBucketOne != goalAmount && s.SizeBucketTwo != goalAmount {
		return 0, 0, fmt.Errorf("not reached")
	}

	if s.SizeBucketOne == goalAmount {

		return 1, s.SizeBucketTwo, nil
	}

	return 2, s.SizeBucketOne, nil
}

func (s State) generateNextStates(initialSizeBucketOne, initialSizeBucketTwo int) []State {
	states := make([]State, 0)

	parentsState := append([]State{}, s.ParentsState...)

	//a possible state is to fill the first bucket (if not already filled)
	if s.SizeBucketOne != initialSizeBucketOne {
		states = append(states, State{initialSizeBucketOne, s.SizeBucketTwo, append(parentsState, s)})
	}

	//a possible state is to fill the second bucket (if not already filled AND bucket one not empty)
	if s.SizeBucketTwo != initialSizeBucketTwo && s.SizeBucketOne > 0 {
		states = append(states, State{s.SizeBucketOne, initialSizeBucketTwo, append(parentsState, s)})
	}

	//a possible state is to empty the first bucket (if not already empty AND bucket two not filled)
	if s.SizeBucketOne != 0 && s.SizeBucketTwo != initialSizeBucketTwo {
		states = append(states, State{0, s.SizeBucketTwo, append(parentsState, s)})
	}

	//a possible state is to empty the second bucjet (if not already empty)
	if s.SizeBucketTwo != 0 {
		states = append(states, State{s.SizeBucketOne, 0, append(parentsState, s)})
	}

	//a possible state is to pour the first bucket in the second (if the first is not empty and the second is not full and if after this)
	if s.SizeBucketOne != 0 && s.SizeBucketTwo != initialSizeBucketTwo {

		maxSizeToFill := initialSizeBucketTwo - s.SizeBucketTwo
		if maxSizeToFill > s.SizeBucketOne { //the first will be empty but the second won't be full, so, it is accepted
			states = append(states, State{0, s.SizeBucketOne + s.SizeBucketTwo, append(parentsState, s)})
		} else if maxSizeToFill < s.SizeBucketOne { //the second will be filled but the first won't be empty, so, it is accepted too
			states = append(states, State{s.SizeBucketOne - maxSizeToFill, initialSizeBucketTwo, append(parentsState, s)})
		} //third case is not accepted : the first will be empty and the second will be filled !!!!
	}

	//a possible state is to pour the second bucket in the first (if the first is not full and the second is not empty)
	if s.SizeBucketOne != initialSizeBucketOne && s.SizeBucketTwo != 0 {

		maxSizeToFill := initialSizeBucketOne - s.SizeBucketOne
		if maxSizeToFill >= s.SizeBucketTwo { //the second will be empty
			states = append(states, State{s.SizeBucketOne + s.SizeBucketTwo, 0, append(parentsState, s)})
		} else {
			states = append(states, State{s.SizeBucketOne + maxSizeToFill, s.SizeBucketTwo - maxSizeToFill, append(parentsState, s)})
		}
	}

	return states

}

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {

	//simply returns an error if startBucket neigther "one" or "two"
	if startBucket != "one" && startBucket != "two" {
		return "", 0, 0, fmt.Errorf("invalid start bucket name")
	}

	//error handling when buckets size invalid (0 or negative)
	if sizeBucketOne <= 0 {
		return "", 0, 0, fmt.Errorf("invalid first bucket size")
	}

	if sizeBucketTwo <= 0 {
		return "", 0, 0, fmt.Errorf("invalid second bucket size")
	}

	//error handling when goal is 0 or negative
	if goalAmount <= 0 {
		return "", 0, 0, fmt.Errorf("invalid goal amount")
	}

	//error handling when goal is larger than buckets size
	if goalAmount > sizeBucketOne && goalAmount > sizeBucketTwo {
		return "", 0, 0, fmt.Errorf("impossible")
	}

	//revert bucket one and two if the first to be fill is the second
	if startBucket == "two" {

		sizeBucketOne, sizeBucketTwo = sizeBucketTwo, sizeBucketOne
	}

	//nextStatesToTest
	nextStatesToTest := []State{
		{0, 0, []State{}},
	}

	numSteps := 0

	//existingStates will be populate by the states already found
	existingStates := map[string]struct{}{
		"(0,0)": {},
	}

	for len(nextStatesToTest) != 0 {
		nextStates := []State{}

		for _, state := range nextStatesToTest {
			//test if this one is THE one
			if goalBucket, otherBucketLevel, err := state.IsGoal(goalAmount); err == nil { //We find it !!!
				//check for reversing
				if startBucket == "two" {
					if goalBucket == 1 {
						goalBucket = 2
					} else {
						goalBucket = 1
					}
				}
				if goalBucket == 1 {
					return "one", numSteps, otherBucketLevel, nil
				}
				return "two", numSteps, otherBucketLevel, nil
			}
			for _, nState := range state.generateNextStates(sizeBucketOne, sizeBucketTwo) {
				//test if the state not already exists
				if _, ok := existingStates[nState.String()]; !ok { //not tested yet !
					existingStates[nState.String()] = struct{}{}
					nextStates = append(nextStates, nState)
				}
			}
		}
		nextStatesToTest = nextStates
		numSteps++
	}

	return "", 0, 0, fmt.Errorf("impossible")
}
