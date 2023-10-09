package main

type TimeSlot struct {
	Start int32
	End   int32
}

type Machine struct {
	TimeSlots []TimeSlot
}

const maxInt32 = int32(^uint32(0) >> 1)

// getLastTimeSlot returns the last time slot in a machine's schedule.
func getLastTimeSlot(machine Machine) TimeSlot {
	return machine.TimeSlots[len(machine.TimeSlots)-1]
}

// chooseAvailableMachine selects the machine with the earliest available time slot.
func chooseAvailableMachine(machines []Machine) int {
	var leastTime int32 = maxInt32 // Max int32 value
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

// allocateTask allocates a task to a machine based on the provided time slot.
func allocateTask(timeslot TimeSlot, machines []Machine) []Machine {
	// For the first allocation
	if len(machines) == 0 {
		return []Machine{
			{TimeSlots: []TimeSlot{
				timeslot,
			}},
		}
	}

	availableMachineIndex := chooseAvailableMachine(machines)
	availableMachineTimeSlots := machines[availableMachineIndex].TimeSlots

	// Allocate task to a new machine
	if timeslot.Start <= availableMachineTimeSlots[len(availableMachineTimeSlots)-1].End {
		machines = append(machines, Machine{
			TimeSlots: []TimeSlot{timeslot},
		})
		return machines
	}

	// Allocate task to an existing machine
	machines[availableMachineIndex].TimeSlots = append(machines[availableMachineIndex].TimeSlots, timeslot)
	return machines
}

// GetLeastRequiredMachine calculates the least number of machines required to perform all tasks.
func GetLeastRequiredMachine(start, end []int32) int32 {
	var machines []Machine
	for i := 0; i < len(start); i++ {
		machines = allocateTask(TimeSlot{
			Start: start[i],
			End:   end[i],
		}, machines)
	}
	return int32(len(machines))
}
