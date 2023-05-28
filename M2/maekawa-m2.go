package main

import (
	"fmt"
	"math/rand"
	"time"
)

var colorReset = "\033[0m"

var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"
var colorWhite = "\033[37m"

// Process represents a process in the system.
type Process struct {
	id int
	state string
	queue   []int
}

// NewProcess creates a new process.
func NewProcess(id int) *Process {
	return &Process{
		id: id,
		state: "ready",
		queue: make([]int, 0),
	}
}

// RequestCS requests permission to enter the critical section.
func (process *Process) RequestCS() {
	fmt.Println(string(colorCyan), "Process", process.id, "requests CS", string(colorReset))

	// Send request messages to all processes in the quorum.
	for _, pid := range process.queue {
		fmt.Println("Process", process.id, "sends request message to", pid)
	}

	// Wait for reply messages from a majority of processes in the quorum.
	granted := false
	for i := 0; i < len(process.queue); i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		pid := process.queue[i]
		fmt.Println("Process", process.id, "waits for reply message from", pid)

		// If a majority of processes have granted permission, enter the critical section.
		if pid != process.id && process.state == "ready" {
			fmt.Println("Process", process.id, "receives reply message from", pid)
			granted = granted || true
		}
	}

	// If permission was not granted, block.
	if !granted {
		fmt.Println(string(colorRed), "Process", process.id, "is blocked", string(colorReset))
		process.state = "blocked"
	} else {
		fmt.Println(string(colorRed), "Process", process.id, "enters CS", string(colorReset))
		process.state = "critical"
	}
}

// ReleaseCS releases the critical section.
func (process *Process) ReleaseCS() {
	fmt.Println(string(colorRed), "Process", process.id, "releases CS", string(colorReset))

	// Send release messages to all processes in the quorum.
	for _, pid := range process.queue {
		fmt.Println("Process", process.id, "sends release message to", pid)
	}

	// Change state to ready.
	process.state = "ready"
}

func main() {

    fmt.Printf("\n")
    fmt.Println(string(colorRed), "-------------------", string(colorReset))
    fmt.Println(string(colorRed), "Maekawa's Algorithm", string(colorReset))
    fmt.Println(string(colorRed), "-------------------", string(colorReset))
    fmt.Printf("\n")
	
	// Create a set of processes.
	var numProcs int
	fmt.Println(string(colorCyan), "Enter the number of processes: ", string(colorReset))
	fmt.Scanln(&numProcs)

	processes := make([]*Process, numProcs)
	for i := 0; i < numProcs; i++ {
		processes[i] = NewProcess(i)
	}

	// Set the quorums for each process.
	for i := 0; i < numProcs; i++ {
		processes[i].queue = []int{(i + 1) % numProcs, (i + 2) % numProcs}
	}

	var numberOfIterations int
	fmt.Println(string(colorCyan), "Enter the number of iterations: ", string(colorReset))
	fmt.Scanln(&numberOfIterations)

    // Print the quorums.
    fmt.Printf("\n")
    fmt.Println(string(colorPurple), "-------------------", string(colorReset))
    fmt.Println(string(colorPurple), "Quorums", string(colorReset))
    fmt.Println(string(colorPurple), "-------------------", string(colorReset))
    for _, process := range processes {
        fmt.Printf("Process %d: ", process.id)
        for _, pid := range process.queue {
            fmt.Printf("%d ", pid)
        }
        fmt.Printf("\n")
    }
    fmt.Println(string(colorPurple), "-------------------", string(colorReset))

    fmt.Printf("\n")

	// Run the simulation for 1000 iterations.
	for i := 0; i < numberOfIterations; i++ {
        fmt.Println(string(colorYellow), "Iteration", i+1, string(colorReset))
		// Randomly select a process to request the critical section.
		pid := rand.Intn(numProcs)
		processes[pid].RequestCS()

		// Wait for all processes to finish with the critical section.
		for _, process := range processes {
			if process.state == "critical" {
				process.ReleaseCS()
            }
        }
        fmt.Printf("\n")
    }
}