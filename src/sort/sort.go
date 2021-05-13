package sort

import "fmt"

func QuickSort(a []int, lidx, ridx int) {
	if lidx >= ridx {
		return
	}
	oriLidx, oriRidx := lidx, ridx
	sentinel := a[lidx]
	for ; lidx < ridx; {
		for ; lidx < ridx && a[ridx] >= sentinel; {
			ridx--
		}
		a[lidx] = a[ridx]

		for ; lidx < ridx && a[lidx] < sentinel; {
			lidx++
		}
		a[ridx] = a[lidx]

	}
	a[lidx] = sentinel
	QuickSort(a, oriLidx, lidx)
	QuickSort(a, lidx+1, oriRidx)
}

func swap(a []int, l, r int) {
	a[l], a[r] = a[r], a[l]
}

func bubbleSort(a []int) {
	for idx, num := range a {
		for i := idx; i < len(a); i++ {
			if num > a[i] {
				swap(a, idx, i)
			}
		}
	}
}

func MergeSort(a []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l+r)/2 + 1
	MergeSort(a, l, mid-1)
	MergeSort(a, mid, r)
	merge(a, l, mid, r)
}

func merge(a []int, l, m, r int) {
	left, right := append([]int{}, a[l:m]...), append([]int{}, a[m:r+1]...)
	left = append(left, 2147483647)
	right = append(right, 2147483647)
	fmt.Println(left, right)
	var lidx, ridx int
	for i := l; i <= r; i++ {
		if left[lidx] < right[ridx] {
			a[i] = left[lidx]
			if lidx < len(left) {
				lidx++
			}
		} else {
			a[i] = right[ridx]
			if ridx < len(right) {
				ridx++
			}
		}
	}
}
