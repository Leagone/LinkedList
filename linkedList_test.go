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
