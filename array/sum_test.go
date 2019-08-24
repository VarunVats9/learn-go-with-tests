package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("Got %d want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	t.Run("collection of any number of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3, 4}, []int{1, 2, 3, 4})
		want := []int{10, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("sum any number of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{3, 4})
		want := []int{9, 4}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4})
		want := []int{0, 4}

		checkSums(t, got, want)
	})
}
