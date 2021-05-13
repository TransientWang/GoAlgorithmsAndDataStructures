package sort

import (
	"fmt"
	"testing"
)

var a = []int{44, 3, 243, 1, 98, 50, 45, -1}

func TestQuickSort(t *testing.T) {

	fmt.Println(a)
	QuickSort(a, 0, len(a)-1)
	t.Log(a)
}

func TestBubbleSort(t *testing.T) {

	fmt.Println(a)
	bubbleSort(a)
	t.Log(a)
}

func TestMergeSort(t *testing.T) {
	//x := []int{6,5,4,3,2,1}
	MergeSort(a, 0, len(a)-1)
	t.Log(a)
}
