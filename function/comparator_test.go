package function

import "testing"

func TestIntComparator(t *testing.T) {
	// i1,i2,expected
	tests := [][]interface{}{
		{1, 1, 0},
		{nil, 0, 0},
	}

	for _, test := range tests {
		actual := IntComparator(test[0], test[1])
		expected := test[2]
		if actual != expected {
			t.Errorf("Got %v expected %v", actual, expected)
		}
	}
}
