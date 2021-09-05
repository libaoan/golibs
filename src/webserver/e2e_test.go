package webserver

import "testing"

func TestClient(t *testing.T) {
	go RunServer()

	testPassedCases := []struct{
		description string
		input []string
		expectedOutput []bool
	}{
		{
			"server initial, expected all false",
			[]string{"1", "2", "3"},
			 []bool{ false, false, false},
		},{
			"empty input",
			[]string{},
			[]bool{},
		},{
			"expected all true",
			[]string{"1", "2", "3"},
			[]bool{ true, true, true},
		},{
			"expected 1/4 true",
			[]string{"3", "4", "5", "6"},
			[]bool{ true, false, false, false},
		},
	}

	for _, testcase := range testPassedCases {
		result, err := Client(testcase.input)
		if err != nil{
			t.Fatal(err)
		}
		if len(result) != len(testcase.expectedOutput) {
			t.Errorf("the length of the output[%d] is not matched with input[%d]", len(result), len(testcase.input))
		}
		for i := range result {
			if result[i] != testcase.expectedOutput[i] {
				t.Errorf("result[%#v] and expected result[%#v] is not matched", result, testcase.expectedOutput)
				break
			}
		}
	}
	t.Log("All tests has passed")
}
