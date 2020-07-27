package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// It's important to note that reflect.DeepEqual is not "type safe" 
	if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
			t.Helper()
			if !reflect.DeepEqual(got, want) {
					t.Errorf("got %v want %v", got, want)
			}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
			got := SumAllTails([]int{1, 2}, []int{0, 9})
			want := []int{2, 9}
			checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
			got := SumAllTails([]int{}, []int{3, 4, 5})
			want := []int{0, 9}
			checkSums(t, got, want)
	})

}