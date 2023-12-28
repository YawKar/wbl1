package _16

import (
	"sync"

	"golang.org/x/exp/constraints"
)

// Sorts the `slice` in-place using [QuickSort] algorithm.
//
// Choice of pivot is [median-of-three] based on Robert Sedgewick's recommendation.
//
// This implementation takes advantage of the [repeated elements optimization].
//
// [QuickSort]: https://en.wikipedia.org/wiki/Quicksort
// [median-of-three]: https://en.wikipedia.org/wiki/Quicksort#Choice_of_pivot
// [repeated elements optimization]: https://en.wikipedia.org/wiki/Quicksort#Repeated_elements
func QuickSort[T constraints.Ordered](slice []T) {
	if len(slice) < 2 {
		// there's nothing to do if the range has fewer than two elements
		return
	}
	// median-of-three will place pivot in [high]
	medianOfThree(slice)
	pivot := slice[len(slice)-1]
	lt, gt := partition(pivot, slice)
	QuickSort(slice[:lt])
	QuickSort(slice[gt+1:])
}

// Paralleled version of QuickSort that uses goroutines and wait groups
func QuickSortParalleled[T constraints.Ordered](slice []T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go quickSortParalleled(slice, &wg)
	wg.Wait()
}

func quickSortParalleled[T constraints.Ordered](slice []T, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(slice) < 2 {
		// there's nothing to do if the range has fewer than two elements
		return
	}
	// median-of-three will place pivot in [high]
	medianOfThree(slice)
	pivot := slice[len(slice)-1]
	lt, gt := partition(pivot, slice)
	wg.Add(1)
	go quickSortParalleled(slice[:lt], wg)
	wg.Add(1)
	go quickSortParalleled(slice[gt+1:], wg)
}

func medianOfThree[T constraints.Ordered](slice []T) {
	mid := len(slice) / 2
	if slice[mid] < slice[0] {
		slice[0], slice[mid] = slice[mid], slice[0]
	}
	if slice[len(slice)-1] < slice[0] {
		slice[0], slice[len(slice)-1] = slice[len(slice)-1], slice[0]
	}
	if slice[mid] < slice[len(slice)-1] {
		slice[mid], slice[len(slice)-1] = slice[len(slice)-1], slice[mid]
	}
}

func partition[T constraints.Ordered](pivot T, slice []T) (lessThanIx, greatherThanIx int) {
	lessThanIx = 0
	equalIx := 0
	greatherThanIx = len(slice) - 1

	for equalIx <= greatherThanIx {
		if slice[equalIx] < pivot {
			slice[equalIx], slice[lessThanIx] = slice[lessThanIx], slice[equalIx]
			lessThanIx++
			equalIx++
		} else if slice[equalIx] > pivot {
			slice[equalIx], slice[greatherThanIx] = slice[greatherThanIx], slice[equalIx]
			greatherThanIx--
		} else {
			equalIx++
		}
	}
	return
}
