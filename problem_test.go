package main

import (
	"strconv"
	"testing"
)

type testTimeSlot struct {
	start    []int32
	end      []int32
	expected int32
}

var testTimeSlots = []testTimeSlot{
	{
		start:    []int32{2, 1, 5, 5, 8},
		end:      []int32{5, 3, 8, 6, 12},
		expected: 3,
	},
	{
		start:    []int32{10, 9, 8, 10},
		end:      []int32{11, 11, 9, 14},
		expected: 3,
	},
}

func TestLeastRequiredMachine(t *testing.T) {
	for index, test := range testTimeSlots {
		testName := "Test" + strconv.Itoa(index)
		t.Run(testName, func(t *testing.T) {
			if result := GetLeastRequiredMachine(test.start, test.end); result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

func BenchmarkLeastRequiredMachine(b *testing.B){
	for i:=0; i<b.N; i++{
		GetLeastRequiredMachine(testTimeSlots[0].start, testTimeSlots[0].end)
	}
}
