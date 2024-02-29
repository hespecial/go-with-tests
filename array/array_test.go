package main

import "testing"

func TestSum(t *testing.T) {
	assertSum := func(t *testing.T, got, want int, nums []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, nums)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := Sum(nums)
		want := 15

		assertSum(t, got, want, nums)
	})
}
