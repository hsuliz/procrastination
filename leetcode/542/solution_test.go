package _542

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	given := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	want := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 2, 1},
	}
	got := updateMatrix(given)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
