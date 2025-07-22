package _150

import "testing"

func TestEvalRPN(t *testing.T) {
	testCases := []struct {
		name  string
		given []string
		want  int
	}{
		{"2 1 + 3 *", []string{"2", "1", "+", "3", "*"}, 9},
		{"4 13 5 / +", []string{"4", "13", "5", "/", "+"}, 6},
		{"big ass", []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := evalRPN(tc.given)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
