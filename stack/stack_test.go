package stack

import (
	"testing"

	"github.com/NimaC/go-merge-intervals/interval"
)

func TestBasicFunctionality(t *testing.T) {
	intervalStack := Stack{}
	intervalStack = intervalStack.Push(interval.Interval{0, 0})
	_, element, hasElements := intervalStack.Pop()
	if !hasElements {
		t.Errorf("Nonempty stack returns no Elements flag: %t", hasElements)
	}
	if element.Start != 0 || element.End != 0 {
		t.Errorf("Stack element Values invalid: [%d, %d]", element.Start, element.End)
	}
}

func TestEmptyStack(t *testing.T) {
	intervalStack := Stack{}
	_, _, hasElements := intervalStack.Pop()
	if hasElements {
		t.Errorf("Empty stack returns positive hasElements flag: %t", hasElements)
	}
}
