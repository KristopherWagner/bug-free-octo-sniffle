package sorting

import (
	"fmt"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("Bubble sort should handle a presorted array", handlePresortedArrayBubble)
	t.Run("Bubble sort should handle a reversed array", handleReversedArrayBubble)
	t.Run("Bubble sort should handle a random array", handleRandomArrayBubble)
}

func handlePresortedArrayBubble(t *testing.T) {
	fmt.Printf("Testing bubble sort with a presorted array")
	psA := generatePresortedArray(500)
	results, err := BubbleSort(psA)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !sort.IntsAreSorted(results) {
		fmt.Printf(" - Failure\n")
		t.Error("Presorted array did not remain sorted")
	} else {
		fmt.Printf(" - Success\n")
	}
}

func handleReversedArrayBubble(t *testing.T) {
	fmt.Printf("Testing bubble sort with a reversed array")
	revA := generateReversedArray(500)
	results, err := BubbleSort(revA)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !sort.IntsAreSorted(results) {
		fmt.Printf(" - Failure\n")
		t.Error("Failed to sort reversed array")
	} else {
		fmt.Printf(" - Success\n")
	}
}

func handleRandomArrayBubble(t *testing.T) {
	fmt.Printf("Testing bubble sort with multiple random arrays")
	var err error
	for i := 0; i < 5 && err == nil; i++ {
		randA := generateRandomArray(500)
		var results []int
		results, err = BubbleSort(randA)
		if err != nil {
			fmt.Printf(" - Failure\n")
			t.Fatalf(err.Error())
		}
		if !sort.IntsAreSorted(results) {
			fmt.Printf(" - Failure\n")
			err = fmt.Errorf("failed to sort random array")
			t.Error(err.Error())
		}
	}
	if err == nil {
		fmt.Printf(" - Success\n")
	}
}
