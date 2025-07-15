package _973

import (
	"reflect"
	"testing"
)

func TestKClosest(t *testing.T) {
	givenPoints := [][]int{{1, 3}, {-2, 2}}
	givenK := 1

	got := kClosest(givenPoints, givenK)
	want := [][]int{{-2, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
