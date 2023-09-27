package main

import "fmt"

/*
    Problem statement:
    You are given start time and end time for a task.
    They are provided in slices in this format:
        start = []int32{task1startTime, task2startTime, task3startTime, ...}
        end = []int32{task1endTime, task2endTime, task3endTime, ...}
    Length of both slices will always be equal.

    There are machines which can perform these task.
    Each machine can only do one task at a time.
    Efficiently allocate tasks to the machines so that least numbers of machines
    are required to perform all the tasks.

    Eg:
        start = []int32{2, 1, 5, 5, 8}
	    end = []int32{5, 3, 8, 6, 12}

        Machine 1 = [(2, 5), (8, 12)]
        Machine 2 = [(1, 3), (5, 8)]
        Machine 3 = [(5, 6)]

        So, the least number of machines required is 3.
*/

type TimeSlot struct {
	Start int32
	End   int32
}

type Machine struct {
	TimeSlots []TimeSlot
}

func getLastTimeSlot(machine Machine) TimeSlot {
	if len(machine.TimeSlots) == 0 {
		return TimeSlot{}
	}
	return machine.TimeSlots[len(machine.TimeSlots)-1]
}

func chooseAvailableMachine(machines []Machine) int {
	var leastTime int32 = 100000 // Set to some imaginary big number
	var availableMachineIndex int

	for index, machine := range machines {
		timeSlot := getLastTimeSlot(machine)
		if timeSlot.End < leastTime {
			leastTime = timeSlot.End
			availableMachineIndex = index
		}
	}
	return availableMachineIndex
}

func pushAndUpdateTimeSlotInMachines(timeslot TimeSlot, totalMachines []Machine) []Machine {
	availableMachineIndex := chooseAvailableMachine(totalMachines)
	
    // For the first allocation
    if len(totalMachines) == 0 {
		return []Machine{
			{TimeSlots: []TimeSlot{
				timeslot,
			}},
		}
	}

	availableMachineTimeSlots := totalMachines[availableMachineIndex].TimeSlots

    // Allocate task to a new machine
	if timeslot.Start <= availableMachineTimeSlots[len(availableMachineTimeSlots)-1].End {
		totalMachines = append(totalMachines, Machine{
			TimeSlots: []TimeSlot{timeslot},
		})
		return totalMachines
	}

    // Allocate task to an existing machine
	totalMachines[availableMachineIndex].TimeSlots = append(totalMachines[availableMachineIndex].TimeSlots, timeslot)
	return totalMachines
}

func getLeastRequiredMachine(start, end []int32) int32 {
	var totalMachines []Machine
	for i := 0; i < len(start); i++ {
		totalMachines = pushAndUpdateTimeSlotInMachines(TimeSlot{
			Start: start[i],
			End:   end[i],
		}, totalMachines)
	}
	return int32(len(totalMachines))
}

func main() {
    // Test case 1
	start := []int32{2, 1, 5, 5, 8}
	end := []int32{5, 3, 8, 6, 12}

    // Test case 2
	// start := []int32{10, 9, 8, 10}
	// end := []int32{11, 11, 9, 14}

	fmt.Println(getLeastRequiredMachine(start, end))
}
