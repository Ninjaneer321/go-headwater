package headwater

import (
	"fmt"
	"testing"
)

func TestForEach(t *testing.T) {
	items := []int64{0, 1, 2}
	result := 0

	ForEach(items, func(item int64) {
		result++
	})

	want := 3

	if result != want {
		t.Errorf("Count is wrong.  Expected: %#v, Received: %#v", want, result)
	}
}

// Map should create an array with a new value for each value of the original array
func TestMapEach(t *testing.T) {
	items := []int64{0, 1, 2}

	result := Map(items, func(item int64) string {
		return "" + fmt.Sprint(item)
	})

	want := []string{"0", "1", "2"}

	if !Equal(want, result) {
		t.Errorf("Count is wrong.  Expected: %#v, Received: %#v", want, result)
	}
}

func TestNestedMap(t *testing.T) {
	items := []int64{0, 1, 2}

	result := Map(Map(Map(
		items,
		func(item int64) float64 {
			return float64(item) / 2
		}),
		func(item float64) string {
			return "" + fmt.Sprint(item)
		}),
		func(item string) string {
			return "a" + item
		})

	want := []string{"a0", "a0.5", "a1"}

	if !Equal(want, result) {
		t.Errorf("Count is wrong.  Expected: %#v, Received: %#v", want, result)
	}
}

func TestReduce(t *testing.T) {
	items := []int{0, 1, 2}
	target := 0

	result := Reduce(items, func(target int, item int) int {
		return target + item
	}, target)

	want := 3

	if result != want {
		t.Errorf("Count is wrong.  Expected: %#v, Received: %#v", want, result)
	}
}

func TestFilter(t *testing.T) {
	items := []int{0, 3, 1, 4, 2, 5}

	result := Filter(items, func(item int) bool {
		if item >= 3 {
			return true
		} else {
			return false
		}
	})

	out := fmt.Sprintf("%v, %v", items, result)
	fmt.Println(out)

	if !Equal(result, []int{3, 4, 5}) {
		t.Error("Count is wrong")
	}
}
