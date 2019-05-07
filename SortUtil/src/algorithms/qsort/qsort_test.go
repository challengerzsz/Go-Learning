package qsort

import "testing"

func TestQuickSort(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	QuickSort(values)
	result := []int{1, 2, 3, 4, 5}
	for i, v := range values {
		if result[i] != v {
			t.Error("QuickSort failed. Got", values)
		}
	}
}
