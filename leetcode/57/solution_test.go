package _57

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	givenIntervals := [][]int{
		{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	}
	givenNewIntervals := []int{4, 8}
	got := insert(givenIntervals, givenNewIntervals)
	want := [][]int{
		{1, 2}, {3, 8}, {4, 8}, {6, 7}, {8, 10}, {12, 16},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
