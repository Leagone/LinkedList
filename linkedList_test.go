package main

import (
	"testing"
)

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
			if l.Length != tt.expectedLength {
				t.Errorf("Length = %d; want %d", l.Length, tt.expectedLength)
			}
			actualValues := toSlice(l)

			if len(actualValues) != len(tt.expectedValues) {
				t.Fatalf("List values count = %d; want %d", len(actualValues), len(tt.expectedValues))
			}

			for i := range actualValues {
				if actualValues[i] != tt.expectedValues[i] {
					t.Errorf("Value at node %d = %d; want %d", i, actualValues[i], tt.expectedValues[i])
				}
			}
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

			if l.Length != tt.expectedLength {
				t.Errorf("Length = %d; want %d", l.Length, tt.expectedLength)
			}

			actualValues := toSlice(l)

			if len(actualValues) != len(tt.expectedValues) {
				t.Fatalf("List values count = %d; want %d", len(actualValues), len(tt.expectedValues))
			}

			for i := range actualValues {
				if actualValues[i] != tt.expectedValues[i] {
					t.Errorf("Value at node %d = %d; want %d", i, actualValues[i], tt.expectedValues[i])
				}
			}
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

			if l.Length != tt.expectedLength {
				t.Errorf("Length after pop = %d; want %d", l.Length, tt.expectedLength)
			}

			actualValues := toSlice(l)
			if len(actualValues) != len(tt.expectedList) {
				t.Fatalf("Remaining nodes count = %d; want %d", len(actualValues), len(tt.expectedList))
			}

			for i := range actualValues {
				if actualValues[i] != tt.expectedList[i] {
					t.Errorf("Value at node %d = %d; want %d", i, actualValues[i], tt.expectedList[i])
				}
			}
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

			if l.Length != tt.expectedLength {
				t.Errorf("Length = %d; want %d", l.Length, tt.expectedLength)
			}

			actualValues := toSlice(l)
			if len(actualValues) != len(tt.expectedList) {
				t.Fatalf("List values count = %d; want %d", len(actualValues), len(tt.expectedList))
			}

			for i := range actualValues {
				if actualValues[i] != tt.expectedList[i] {
					t.Errorf("Value at node %d = %d; want %d", i, actualValues[i], tt.expectedList[i])
				}
			}
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

			if l.Length != tt.expectedLength {
				t.Errorf("Length = %d; want %d", l.Length, tt.expectedLength)
			}

			actualValues := toSlice(l)
			if len(actualValues) != len(tt.expectedList) {
				t.Fatalf("List values count = %d; want %d", len(actualValues), len(tt.expectedList))
			}

			for i := range actualValues {
				if actualValues[i] != tt.expectedList[i] {
					t.Errorf("Value at node %d = %d; want %d", i, actualValues[i], tt.expectedList[i])
				}
			}
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

			if l.Length != tt.expectedLength {
				t.Errorf("Length = %d; want %d", l.Length, tt.expectedLength)
			}

			actualValues := toSlice(l)
			if len(actualValues) != len(tt.expectedList) {
				t.Fatalf("List values count = %d; want %d", len(actualValues), len(tt.expectedList))
			}

			for i := range actualValues {
				if actualValues[i] != tt.expectedList[i] {
					t.Errorf("Value at node %d = %d; want %d", i, actualValues[i], tt.expectedList[i])
				}
			}
		})
	}
}
