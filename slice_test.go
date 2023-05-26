package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestApply(t *testing.T) {
	s := []int{1, 2, 3, 4}

	fmt.Println(s)

	want := []int{2, 4, 6, 8}

	fn := func(i int) int {
		return i * 2
	}

	if got := Apply(s, fn); !reflect.DeepEqual(got, want) {
		t.Errorf("Apply() = %v, want %v", got, want)
	}
}

func TestReduce(t *testing.T) {
	s := []int{1, 2, 3, 4}
	want := 10

	fn := func(i, initVal int) int {
		return initVal + i
	}

	if got := Reduce(s, fn, 0); !reflect.DeepEqual(got, want) {
		t.Errorf("Reduce() = %v, want %v", got, want)
	}
}

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4}
	want := []int{2, 3, 4}

	fn := func(i int) bool {
		return i > 1
	}

	if got := Filter(s, fn); !reflect.DeepEqual(got, want) {
		t.Errorf("Filter() = %v, want %v", got, want)
	}
}
