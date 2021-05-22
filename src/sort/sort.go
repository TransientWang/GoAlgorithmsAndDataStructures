package sort

import (
	"fmt"
)

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

func maxHeapify(a []int, i, size int) {
	var (
		max   int
		left  = i << 1
		right = i<<1 + 1
	)

	if left < size && a[i] < a[left] {
		max = left
	} else {
		max = i
	}
	if right < size && a[max] < a[right] {
		max = right
	}

	if max != i {
		swap(a, i, max)
		maxHeapify(a, max, size)
	}
}

func buildMaxHeadp(a []int) {
	for i := len(a)/2 + 1; i >= 0; i-- {
		maxHeapify(a, i, len(a))
	}
}

func heapSort(a []int) {
	buildMaxHeadp(a)

	for i := 0; i < len(a); i++ {
		swap(a, 0, len(a)-i-1)
		maxHeapify(a, 0, len(a)-i-1)
	}

}

func quickPartition(a []int, l, r int) int {
	x := a[r]
	i := l
	for j := l; j < r; j++ {

		if a[j] < x {
			swap(a, i, j)
			i++
		}
	}
	swap(a, i, r)
	return i
}

func quickSort(a []int, l, r int) {
	if l >= r {
		return
	}
	i := quickPartition(a, l, r)
	fmt.Println(a, i, l, r)
	quickSort(a, l, i-1)
	quickSort(a, i, r)
}

func countingSort(a []int) []int {
	var max int
	for _, i2 := range a {
		if i2 > max {
			max = i2
		}
	}
	b := make([]int, len(a)+1, len(a)+1)
	c := make([]int, max+1, max+1)
	for _, val := range a {
		c[val]++
	}
	for i := range c {
		if i == 0 {
			continue
		}
		c[i] += c[i-1]
	}
	fmt.Println(c)

	for i := len(a) - 1; i >= 0; i-- {
		fmt.Println(a[i],c[a[i]])
		b[c[a[i]]] = a[i]
		c[a[i]]--
	}
	return b
}
