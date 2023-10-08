# Least Machine Problem

## Problem Statement
You are given start time and end time for a task. They are provided in slices in this format:

    start = []int32{task1startTime, task2startTime, task3startTime, ...}
    end = []int32{task1endTime, task2endTime, task3endTime, ...}

There can be several machines which perform these tasks. Each machine can only do one task at a time. Efficiently allocate tasks to the machines so that `least` numbers of machines are required to perform all the tasks.

## Constraints
- Length of both slices will always be `equal`.
- `start[i]` will always be less than `end[i]`

## Example
    start = []int32{2, 1, 5, 5, 8}
    end = []int32{5, 3, 8, 6, 12}

> Machine 1 = [(2, 5), (8, 12)]

> Machine 2 = [(1, 3), (5, 8)]

> Machine 3 = [(5, 6)]

So, the `least` number of machines required, `count(Machine)` is `3`.