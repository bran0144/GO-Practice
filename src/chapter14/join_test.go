package chapter14

import (
	"fmt"
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestJointWithCommas(t *testing.T) {
	tests := []testData{
		testData{list: []string{}, want: ""},
		testData{list: []string{"apple"}, want: "apple"},
		testData{list: []string{"apple", "orange"}, want: "apple and orange"},
		testData{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinwithCommas(%#v) = \"%s\", want \"%s\"", test.list, got, test.want)
		}
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinwithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}
