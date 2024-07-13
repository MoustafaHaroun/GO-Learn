package Chapter_4

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Collection of two slices summed together", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{1, 2, 4})
		want := []int{6, 7}

		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Collection sum with only the tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 2, 4})
		want := []int{5, 6}

		checkSums(t, got, want)
	})

	t.Run("Collection with empty tail", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{1, 3, 4})
		want := []int{7}

		checkSums(t, got, want)
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Println(sum) //Output: 15
}

func ExampleSumAll() {
	sum := SumAll([]int{1, 2, 3}, []int{1, 2, 4})
	fmt.Println(sum) //Output: [6 7]
}

func ExampleSumAllTails() {
	sum := SumAllTails([]int{1}, []int{1, 3, 4})
	fmt.Println(sum) //Output: [7]
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3}, []int{1, 2, 4})
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{1, 2, 3}, []int{1, 2, 4})
	}
}
