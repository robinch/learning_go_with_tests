package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collecion of size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want but got %d, %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumOfAllTails(t *testing.T) {
 checkSums := func(t testing.TB, got []int, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}


	t.Run("make the sums of slices", func(t *testing.T) {
		got := SumOfAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("sums empty slices", func(t *testing.T) {
		got := SumOfAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

