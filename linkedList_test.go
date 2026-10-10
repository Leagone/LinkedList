package main

import (
	"testing"
)

func toReverseSlice(l *LinkedList) []int {
	var result []int
	if l.Head == nil {
		return result
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	for current != nil {
		result = append(result, current.Value)
		current = current.Prev
	}
	return result
}

func assertListState(t *testing.T, l *LinkedList, expectedLength int, expectedValues []int) {
	t.Helper()

	if l.Length != expectedLength {
		t.Errorf("Length = %d; want %d", l.Length, expectedLength)
	}

	actualForward := toSlice(l)
	if len(actualForward) != len(expectedValues) {
		t.Fatalf("Forward values count = %d; want %d", len(actualForward), len(expectedValues))
	}
	for i := range actualForward {
		if actualForward[i] != expectedValues[i] {
			t.Errorf("Forward traversal node %d = %d; want %d (Next pointer issue)", i, actualForward[i], expectedValues[i])
		}
	}

	actualBackward := toReverseSlice(l)
	if len(actualBackward) != len(expectedValues) {
		t.Fatalf("Backward values count = %d; want %d", len(actualBackward), len(expectedValues))
	}
	for i := range actualBackward {
		expectedIdx := len(expectedValues) - 1 - i
		if actualBackward[i] != expectedValues[expectedIdx] {
			t.Errorf("Backward traversal node %d = %d; want %d (Prev pointer issue)", i, actualBackward[i], expectedValues[expectedIdx])
		}
	}
}

func TestLinkedList_Append(t *testing.T) {
	tests := []struct {
		name           string
		valuesToAppend []int
		expectedLength int
		expectedValues []int
	}{
		{
			name:           "append to empty list",
			valuesToAppend: []int{10},
			expectedLength: 1,
			expectedValues: []int{10},
		},
		{
			name:           "append two nodes",
			valuesToAppend: []int{10, 20},
			expectedLength: 2,
			expectedValues: []int{10, 20},
		},
		{
			name:           "append multiple nodes",
			valuesToAppend: []int{1, 2, 3, 4, 5},
			expectedLength: 5,
			expectedValues: []int{1, 2, 3, 4, 5},
		},
		{
			name:           "append nothing",
			valuesToAppend: []int{},
			expectedLength: 0,
			expectedValues: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.valuesToAppend {
				l.Append(val)
			}
			assertListState(t, l, tt.expectedLength, tt.expectedValues)
		})
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	tests := []struct {
		name            string
		valuesToPrepend []int
		expectedLength  int
		expectedValues  []int
	}{
		{
			name:            "prepend to empty list",
			valuesToPrepend: []int{10},
			expectedLength:  1,
			expectedValues:  []int{10},
		},
		{
			name:            "prepend multiple nodes",
			valuesToPrepend: []int{10, 20, 30},
			expectedLength:  3,
			expectedValues:  []int{30, 20, 10},
		},
		{
			name:            "prepend nothing",
			valuesToPrepend: []int{},
			expectedLength:  0,
			expectedValues:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.valuesToPrepend {
				l.Prepend(val)
			}
			assertListState(t, l, tt.expectedLength, tt.expectedValues)
		})
	}
}

func TestLinkedList_Pop(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		expectedValue  int
		expectError    bool
		expectedLength int
		expectedList   []int
	}{
		{
			name:           "pop from multi-node list",
			initialValues:  []int{10, 20, 30},
			expectedValue:  30,
			expectError:    false,
			expectedLength: 2,
			expectedList:   []int{10, 20},
		},
		{
			name:           "pop from single-node list",
			initialValues:  []int{10},
			expectedValue:  10,
			expectError:    false,
			expectedLength: 0,
			expectedList:   nil,
		},
		{
			name:           "pop from empty list",
			initialValues:  []int{},
			expectedValue:  0,
			expectError:    true,
			expectedLength: 0,
			expectedList:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.initialValues {
				l.Append(val)
			}

			val, err := l.Pop()

			if tt.expectError {
				if err == nil {
					t.Fatalf("Expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
			}

			if val != tt.expectedValue {
				t.Errorf("Returned value = %d; want %d", val, tt.expectedValue)
			}

			assertListState(t, l, tt.expectedLength, tt.expectedList)
		})
	}
}

func TestLinkedList_Get(t *testing.T) {
	tests := []struct {
		name          string
		initialValues []int
		index         int
		expectedValue int
		expectError   bool
	}{
		{
			name:          "get first element",
			initialValues: []int{10, 20, 30},
			index:         0,
			expectedValue: 10,
			expectError:   false,
		},
		{
			name:          "get middle element",
			initialValues: []int{10, 20, 30},
			index:         1,
			expectedValue: 20,
			expectError:   false,
		},
		{
			name:          "get last element",
			initialValues: []int{10, 20, 30},
			index:         2,
			expectedValue: 30,
			expectError:   false,
		},
		{
			name:          "get negative index",
			initialValues: []int{10, 20, 30},
			index:         -1,
			expectedValue: 0,
			expectError:   true,
		},
		{
			name:          "get index equals length",
			initialValues: []int{10, 20, 30},
			index:         3,
			expectedValue: 0,
			expectError:   true,
		},
		{
			name:          "get index out of bounds",
			initialValues: []int{10, 20, 30},
			index:         10,
			expectedValue: 0,
			expectError:   true,
		},
		{
			name:          "get from empty list",
			initialValues: []int{},
			index:         0,
			expectedValue: 0,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.initialValues {
				l.Append(val)
			}

			val, err := l.Get(tt.index)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}

			if val != tt.expectedValue {
				t.Errorf("Get(%d) = %d; want %d", tt.index, val, tt.expectedValue)
			}
		})
	}
}

func TestLinkedList_InsertAt(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		insertValue    int
		insertIndex    int
		expectError    bool
		expectedLength int
		expectedList   []int
	}{
		{
			name:           "insert at 0 in empty list",
			initialValues:  []int{},
			insertValue:    99,
			insertIndex:    0,
			expectError:    false,
			expectedLength: 1,
			expectedList:   []int{99},
		},
		{
			name:           "insert at 0 in populated list",
			initialValues:  []int{10, 20},
			insertValue:    99,
			insertIndex:    0,
			expectError:    false,
			expectedLength: 3,
			expectedList:   []int{99, 10, 20},
		},
		{
			name:           "insert in the middle",
			initialValues:  []int{10, 30},
			insertValue:    20,
			insertIndex:    1,
			expectError:    false,
			expectedLength: 3,
			expectedList:   []int{10, 20, 30},
		},
		{
			name:           "insert at the end",
			initialValues:  []int{10, 20},
			insertValue:    30,
			insertIndex:    2,
			expectError:    false,
			expectedLength: 3,
			expectedList:   []int{10, 20, 30},
		},
		{
			name:           "insert negative index",
			initialValues:  []int{10, 20},
			insertValue:    99,
			insertIndex:    -1,
			expectError:    true,
			expectedLength: 2,
			expectedList:   []int{10, 20},
		},
		{
			name:           "insert out of bounds positive",
			initialValues:  []int{10, 20},
			insertValue:    99,
			insertIndex:    3,
			expectError:    true,
			expectedLength: 2,
			expectedList:   []int{10, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.initialValues {
				l.Append(val)
			}

			err := l.InsertAt(tt.insertIndex, tt.insertValue)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}

			assertListState(t, l, tt.expectedLength, tt.expectedList)
		})
	}
}

func TestLinkedList_RemoveAt(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		removeIndex    int
		expectError    bool
		expectedLength int
		expectedList   []int
	}{
		{
			name:           "remove at 0",
			initialValues:  []int{10, 20, 30},
			removeIndex:    0,
			expectError:    false,
			expectedLength: 2,
			expectedList:   []int{20, 30},
		},
		{
			name:           "remove in the middle",
			initialValues:  []int{10, 20, 30},
			removeIndex:    1,
			expectError:    false,
			expectedLength: 2,
			expectedList:   []int{10, 30},
		},
		{
			name:           "remove at the end",
			initialValues:  []int{10, 20, 30},
			removeIndex:    2,
			expectError:    false,
			expectedLength: 2,
			expectedList:   []int{10, 20},
		},
		{
			name:           "remove only element",
			initialValues:  []int{10},
			removeIndex:    0,
			expectError:    false,
			expectedLength: 0,
			expectedList:   nil,
		},
		{
			name:           "remove negative index",
			initialValues:  []int{10, 20},
			removeIndex:    -1,
			expectError:    true,
			expectedLength: 2,
			expectedList:   []int{10, 20},
		},
		{
			name:           "remove out of bounds positive",
			initialValues:  []int{10, 20},
			removeIndex:    2,
			expectError:    true,
			expectedLength: 2,
			expectedList:   []int{10, 20},
		},
		{
			name:           "remove from empty list",
			initialValues:  []int{},
			removeIndex:    0,
			expectError:    true,
			expectedLength: 0,
			expectedList:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.initialValues {
				l.Append(val)
			}

			err := l.RemoveAt(tt.removeIndex)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}

			assertListState(t, l, tt.expectedLength, tt.expectedList)
		})
	}
}

func TestLinkedList_Contains(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		searchValue    int
		expectedResult bool
	}{
		{
			name:           "contains first element",
			initialValues:  []int{10, 20, 30},
			searchValue:    10,
			expectedResult: true,
		},
		{
			name:           "contains middle element",
			initialValues:  []int{10, 20, 30},
			searchValue:    20,
			expectedResult: true,
		},
		{
			name:           "contains last element",
			initialValues:  []int{10, 20, 30},
			searchValue:    30,
			expectedResult: true,
		},
		{
			name:           "does not contain element",
			initialValues:  []int{10, 20, 30},
			searchValue:    99,
			expectedResult: false,
		},
		{
			name:           "contains in empty list",
			initialValues:  []int{},
			searchValue:    10,
			expectedResult: false,
		},
		{
			name:           "contains duplicate element",
			initialValues:  []int{10, 20, 20, 30},
			searchValue:    20,
			expectedResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.initialValues {
				l.Append(val)
			}

			result := l.Contains(tt.searchValue)

			if result != tt.expectedResult {
				t.Errorf("Contains(%d) = %v; want %v", tt.searchValue, result, tt.expectedResult)
			}
		})
	}
}

func TestLinkedList_IndexOf(t *testing.T) {
	tests := []struct {
		name          string
		initialValues []int
		searchValue   int
		expectedIndex int
		expectError   bool
	}{
		{
			name:          "index of first element",
			initialValues: []int{10, 20, 30},
			searchValue:   10,
			expectedIndex: 0,
			expectError:   false,
		},
		{
			name:          "index of middle element",
			initialValues: []int{10, 20, 30},
			searchValue:   20,
			expectedIndex: 1,
			expectError:   false,
		},
		{
			name:          "index of last element",
			initialValues: []int{10, 20, 30},
			searchValue:   30,
			expectedIndex: 2,
			expectError:   false,
		},
		{
			name:          "index of non-existent element",
			initialValues: []int{10, 20, 30},
			searchValue:   99,
			expectedIndex: -1,
			expectError:   true,
		},
		{
			name:          "index in empty list",
			initialValues: []int{},
			searchValue:   10,
			expectedIndex: -1,
			expectError:   true,
		},
		{
			name:          "index of duplicate element",
			initialValues: []int{10, 20, 20, 30},
			searchValue:   20,
			expectedIndex: 1,
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.initialValues {
				l.Append(val)
			}

			idx, err := l.IndexOf(tt.searchValue)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}

			if idx != tt.expectedIndex {
				t.Errorf("IndexOf(%d) = %d; want %d", tt.searchValue, idx, tt.expectedIndex)
			}
		})
	}
}

func TestLinkedList_RemoveByValue(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		removeValue    int
		expectError    bool
		expectedLength int
		expectedList   []int
	}{
		{
			name:           "remove head",
			initialValues:  []int{10, 20, 30},
			removeValue:    10,
			expectError:    false,
			expectedLength: 2,
			expectedList:   []int{20, 30},
		},
		{
			name:           "remove middle element",
			initialValues:  []int{10, 20, 30},
			removeValue:    20,
			expectError:    false,
			expectedLength: 2,
			expectedList:   []int{10, 30},
		},
		{
			name:           "remove tail",
			initialValues:  []int{10, 20, 30},
			removeValue:    30,
			expectError:    false,
			expectedLength: 2,
			expectedList:   []int{10, 20},
		},
		{
			name:           "remove only element",
			initialValues:  []int{10},
			removeValue:    10,
			expectError:    false,
			expectedLength: 0,
			expectedList:   nil,
		},
		{
			name:           "remove non-existent element",
			initialValues:  []int{10, 20, 30},
			removeValue:    99,
			expectError:    true,
			expectedLength: 3,
			expectedList:   []int{10, 20, 30},
		},
		{
			name:           "remove from empty list",
			initialValues:  []int{},
			removeValue:    10,
			expectError:    true,
			expectedLength: 0,
			expectedList:   nil,
		},
		{
			name:           "remove first of duplicates",
			initialValues:  []int{10, 20, 20, 30},
			removeValue:    20,
			expectError:    false,
			expectedLength: 3,
			expectedList:   []int{10, 20, 30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.initialValues {
				l.Append(val)
			}

			err := l.RemoveByValue(tt.removeValue)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}

			assertListState(t, l, tt.expectedLength, tt.expectedList)
		})
	}
}

func TestLinkedList_Filter(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		condition      func(value int) bool
		expectedLength int
		expectedList   []int
	}{
		{
			name:          "filter all evens",
			initialValues: []int{1, 2, 3, 4, 5, 6},
			condition: func(value int) bool {
				return value%2 == 0
			},
			expectedLength: 3,
			expectedList:   []int{2, 4, 6},
		},
		{
			name:          "filter keep all",
			initialValues: []int{10, 20, 30},
			condition: func(value int) bool {
				return true
			},
			expectedLength: 3,
			expectedList:   []int{10, 20, 30},
		},
		{
			name:          "filter remove all",
			initialValues: []int{10, 20, 30},
			condition: func(value int) bool {
				return false
			},
			expectedLength: 0,
			expectedList:   nil,
		},
		{
			name:          "filter from empty list",
			initialValues: []int{},
			condition: func(value int) bool {
				return true
			},
			expectedLength: 0,
			expectedList:   nil,
		},
		{
			name:          "filter keep only head",
			initialValues: []int{10, 20, 30},
			condition: func(value int) bool {
				return value == 10
			},
			expectedLength: 1,
			expectedList:   []int{10},
		},
		{
			name:          "filter keep only tail",
			initialValues: []int{10, 20, 30},
			condition: func(value int) bool {
				return value == 30
			},
			expectedLength: 1,
			expectedList:   []int{30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.initialValues {
				l.Append(val)
			}

			l.Filter(tt.condition)

			assertListState(t, l, tt.expectedLength, tt.expectedList)
		})
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		expectedLength int
		expectedList   []int
	}{
		{
			name:           "reverse multi-node list",
			initialValues:  []int{10, 20, 30, 40},
			expectedLength: 4,
			expectedList:   []int{40, 30, 20, 10},
		},
		{
			name:           "reverse two-node list",
			initialValues:  []int{10, 20},
			expectedLength: 2,
			expectedList:   []int{20, 10},
		},
		{
			name:           "reverse single-node list",
			initialValues:  []int{10},
			expectedLength: 1,
			expectedList:   []int{10},
		},
		{
			name:           "reverse empty list",
			initialValues:  []int{},
			expectedLength: 0,
			expectedList:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.initialValues {
				l.Append(val)
			}

			l.Reverse()

			assertListState(t, l, tt.expectedLength, tt.expectedList)
		})
	}
}

func TestLinkedList_Middle(t *testing.T) {
	tests := []struct {
		name          string
		initialValues []int
		expectedValue int
	}{
		{
			name:          "middle of odd length list",
			initialValues: []int{10, 20, 30},
			expectedValue: 20,
		},
		{
			name:          "middle of even length list",
			initialValues: []int{10, 20, 30, 40},
			expectedValue: 30,
		},
		{
			name:          "middle of single element list",
			initialValues: []int{10},
			expectedValue: 10,
		},
		{
			name:          "middle of empty list",
			initialValues: []int{},
			expectedValue: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.initialValues {
				l.Append(val)
			}

			result := l.Middle()

			if result != tt.expectedValue {
				t.Errorf("Middle() = %d; want %d", result, tt.expectedValue)
			}
		})
	}
}

func TestLinkedList_RemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		n              int
		expectError    bool
		expectedLength int
		expectedList   []int
	}{
		{
			name:           "remove 1st from end (last element)",
			initialValues:  []int{10, 20, 30},
			n:              1,
			expectError:    false,
			expectedLength: 2,
			expectedList:   []int{10, 20},
		},
		{
			name:           "remove 2nd from end (middle element)",
			initialValues:  []int{10, 20, 30},
			n:              2,
			expectError:    false,
			expectedLength: 2,
			expectedList:   []int{10, 30},
		},
		{
			name:           "remove 3rd from end (first element)",
			initialValues:  []int{10, 20, 30},
			n:              3,
			expectError:    false,
			expectedLength: 2,
			expectedList:   []int{20, 30},
		},
		{
			name:           "remove from single-node list",
			initialValues:  []int{10},
			n:              1,
			expectError:    false,
			expectedLength: 0,
			expectedList:   nil,
		},
		{
			name:           "remove out of bounds (too large)",
			initialValues:  []int{10, 20},
			n:              3,
			expectError:    true,
			expectedLength: 2,
			expectedList:   []int{10, 20},
		},
		{
			name:           "remove out of bounds (zero)",
			initialValues:  []int{10, 20},
			n:              0,
			expectError:    true,
			expectedLength: 2,
			expectedList:   []int{10, 20},
		},
		{
			name:           "remove out of bounds (negative)",
			initialValues:  []int{10, 20},
			n:              -1,
			expectError:    true,
			expectedLength: 2,
			expectedList:   []int{10, 20},
		},
		{
			name:           "remove from empty list",
			initialValues:  []int{},
			n:              1,
			expectError:    true,
			expectedLength: 0,
			expectedList:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.initialValues {
				l.Append(val)
			}

			err := l.RemoveNthFromEnd(tt.n)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}

			assertListState(t, l, tt.expectedLength, tt.expectedList)
		})
	}
}

func TestLinkedList_HasCycle(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		cycleIndex     int
		expectedResult bool
	}{
		{
			name:           "empty list",
			initialValues:  []int{},
			cycleIndex:     -1,
			expectedResult: false,
		},
		{
			name:           "single node no cycle",
			initialValues:  []int{10},
			cycleIndex:     -1,
			expectedResult: false,
		},
		{
			name:           "single node with cycle",
			initialValues:  []int{10},
			cycleIndex:     0,
			expectedResult: true,
		},
		{
			name:           "multi node no cycle",
			initialValues:  []int{10, 20, 30, 40},
			cycleIndex:     -1,
			expectedResult: false,
		},
		{
			name:           "multi node cycle to head",
			initialValues:  []int{10, 20, 30, 40},
			cycleIndex:     0,
			expectedResult: true,
		},
		{
			name:           "multi node cycle to middle",
			initialValues:  []int{10, 20, 30, 40},
			cycleIndex:     2,
			expectedResult: true,
		},
		{
			name:           "multi node cycle to tail (self loop)",
			initialValues:  []int{10, 20, 30, 40},
			cycleIndex:     3,
			expectedResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.initialValues {
				l.Append(val)
			}

			if tt.cycleIndex >= 0 && l.Head != nil {
				var targetNode *Node
				tail := l.Head

				for i := 0; tail.Next != nil; i++ {
					if i == tt.cycleIndex {
						targetNode = tail
					}
					tail = tail.Next
				}

				if tt.cycleIndex == l.Length-1 {
					targetNode = tail
				}

				if targetNode != nil {
					tail.Next = targetNode
				}
			}

			result := l.HasCycle()

			if result != tt.expectedResult {
				t.Errorf("HasCycle() = %v; want %v", result, tt.expectedResult)
			}
		})
	}
}

func TestLinkedList_FindCycleStart(t *testing.T) {
	tests := []struct {
		name               string
		initialValues      []int
		cycleIndex         int
		expectedCycleIndex int
	}{
		{
			name:               "empty list",
			initialValues:      []int{},
			cycleIndex:         -1,
			expectedCycleIndex: -1,
		},
		{
			name:               "single node no cycle",
			initialValues:      []int{10},
			cycleIndex:         -1,
			expectedCycleIndex: -1,
		},
		{
			name:               "single node with cycle",
			initialValues:      []int{10},
			cycleIndex:         0,
			expectedCycleIndex: 0,
		},
		{
			name:               "multi node no cycle",
			initialValues:      []int{10, 20, 30, 40},
			cycleIndex:         -1,
			expectedCycleIndex: -1,
		},
		{
			name:               "multi node cycle to head",
			initialValues:      []int{10, 20, 30, 40},
			cycleIndex:         0,
			expectedCycleIndex: 0,
		},
		{
			name:               "multi node cycle to middle",
			initialValues:      []int{10, 20, 30, 40},
			cycleIndex:         2,
			expectedCycleIndex: 2,
		},
		{
			name:               "multi node cycle to tail (self loop)",
			initialValues:      []int{10, 20, 30, 40},
			cycleIndex:         3,
			expectedCycleIndex: 3,
		},
		{
			name:               "meeting point differs from start (larger list)",
			initialValues:      []int{10, 20, 30, 40, 50, 60, 70, 80},
			cycleIndex:         2,
			expectedCycleIndex: 2,
		},
		{
			name:               "long tail, small cycle",
			initialValues:      []int{1, 2, 3, 4, 5, 6, 7},
			cycleIndex:         5,
			expectedCycleIndex: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, val := range tt.initialValues {
				l.Append(val)
			}

			var expectedNode *Node

			if tt.cycleIndex >= 0 && l.Head != nil {
				tail := l.Head
				for i := 0; tail.Next != nil; i++ {
					if i == tt.cycleIndex {
						expectedNode = tail
					}
					tail = tail.Next
				}

				if tt.cycleIndex == l.Length-1 {
					expectedNode = tail
				}

				if expectedNode != nil {
					tail.Next = expectedNode
				}
			}

			result := l.FindCycleStart()

			if tt.expectedCycleIndex == -1 {
				if result != nil {
					t.Errorf("FindCycleStart() = %v; want nil", result)
				}
			} else {
				if result == nil {
					t.Fatalf("FindCycleStart() = nil; want node at index %d", tt.expectedCycleIndex)
				}
				if result != expectedNode {
					t.Errorf("FindCycleStart() returned wrong node. Got value %d; want %d", result.Value, expectedNode.Value)
				}
			}
		})
	}
}

func TestLinkedList_Zip(t *testing.T) {
	tests := []struct {
		name            string
		initialL1Values []int
		initialL2Values []int
		expectedLength  int
		expectedList    []int
	}{
		{
			name:            "equal length lists",
			initialL1Values: []int{1, 3, 5},
			initialL2Values: []int{2, 4, 6},
			expectedLength:  6,
			expectedList:    []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:            "l1 is longer",
			initialL1Values: []int{1, 3, 5, 6, 7},
			initialL2Values: []int{2, 4},
			expectedLength:  7,
			expectedList:    []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:            "l2 is longer",
			initialL1Values: []int{1, 3},
			initialL2Values: []int{2, 4, 5, 6, 7},
			expectedLength:  7,
			expectedList:    []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:            "l1 is empty",
			initialL1Values: []int{},
			initialL2Values: []int{1, 2, 3},
			expectedLength:  3,
			expectedList:    []int{1, 2, 3},
		},
		{
			name:            "l2 is empty",
			initialL1Values: []int{1, 2, 3},
			initialL2Values: []int{},
			expectedLength:  3,
			expectedList:    []int{1, 2, 3},
		},
		{
			name:            "both lists empty",
			initialL1Values: []int{},
			initialL2Values: []int{},
			expectedLength:  0,
			expectedList:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := &LinkedList{}
			for _, val := range tt.initialL1Values {
				l1.Append(val)
			}

			l2 := &LinkedList{}
			for _, val := range tt.initialL2Values {
				l2.Append(val)
			}

			l1.Zip(l2)

			assertListState(t, l1, tt.expectedLength, tt.expectedList)
		})
	}
}
