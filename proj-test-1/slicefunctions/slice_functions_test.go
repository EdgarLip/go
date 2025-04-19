package slicefunctions

import "testing"


func TestJoinStringsInSlice(t *testing.T) {
	//input
	sliceOfStrings := []string{"one", "2", "3"}
	seperatortor := ", "

	//execute
    result := JoinStringsInSlice(sliceOfStrings, seperatortor)

	// test
    expected := "one, 2, 3"
    if result != expected {
        t.Errorf("expected %v, got %v", expected, result)
    }
}

func TestJoinStringsInSlice2(t *testing.T) {
	//input
	subTests := []struct {
		sliceOfStrings  []string
		seperator string
		expectedResault string
	}{
		{
			sliceOfStrings: []string{"one", "2", "3"},
			seperator: ", ",
			expectedResault: "one, 2, 3",
		},
		{
			sliceOfStrings:  []string{"one", "two", "333"},
			seperator: ": ",
			expectedResault: "one: two: 333",
		},
		{
			sliceOfStrings: []string{"2", "2", "2"},
			seperator: "| ",
			expectedResault: "2| 2| 2",
		},
	}

	//execute
	for _, subTest := range subTests {
		//test
		if subTestResualt := JoinStringsInSlice(subTest.sliceOfStrings, subTest.seperator); subTestResualt != subTest.expectedResault {
			t.Errorf("got %s + %v, gave %v, wanted %s", subTest.sliceOfStrings, subTest.seperator,
			                                            subTestResualt, subTest.expectedResault)
		}
	}
}