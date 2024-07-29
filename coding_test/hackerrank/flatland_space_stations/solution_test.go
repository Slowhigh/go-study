package flatlandSpaceStations

import "testing"

type testCase struct {
	n      int32
	c      []int32
	result int32
}

func Test_flatlandSpaceStations(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:      5,
			c:      []int32{0,1,2},
			result: 2,
			// 0 1 2 3 4
			// 0 0 0 1 2
		},
		{
			n:      5,
			c:      []int32{0, 4},
			result: 2,
		},
	}

	for i, tc := range testCaseArr {
		res := flatlandSpaceStations(tc.n, tc.c)
		if res == tc.result {
			t.Logf("PASS - %d (%d)", i, res)
		} else {
			t.Errorf("FAIL - %d (%d)", i, res)
		}
	}
}
