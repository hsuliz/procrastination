package _57

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	givenIntervals := [][]int{
		{1, 7}, {8, 10}, {2, 6}, {15, 18}, {5, 7},
	}
	got := merge(givenIntervals)
	want := [][]int{
		{1, 7}, {8, 10}, {15, 18},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
