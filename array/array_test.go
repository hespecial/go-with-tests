package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	got := Sum(nums)
	want := 15
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, nums)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
