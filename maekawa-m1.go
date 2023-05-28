package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Process represents a process in the system.
type Process struct {
    id int
    state string
    q   []int
}

// NewProcess creates a new process.
func NewProcess(id int) *Process {
    return &Process{
        id: id,
        state: "ready",
        q:   make([]int, 0),
    }
}

// RequestCS requests permission to enter the critical section.
func (p *Process) RequestCS() {
    fmt.Printf("Process %d requests CS\n", p.id)

    // Send request messages to all processes in the quorum.
    for _, pid := range p.q {
        fmt.Printf("Process %d sends request message to %d\n", p.id, pid)
    }

    // Wait for reply messages from a majority of processes in the quorum.
    granted := false
    for i := 0; i < len(p.q); i++ {
        time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
        pid := p.q[i]
        fmt.Printf("Process %d waits for reply message from %d\n", p.id, pid)

        // If a majority of processes have granted permission, enter the critical section.
        if pid != p.id && p.state == "ready" {
            fmt.Printf("Process %d receives reply message from %d\n", p.id, pid)
            granted = granted || true
        }
    }

    // If permission was not granted, block.
    if !granted {
        fmt.Printf("Process %d is blocked\n", p.id)
        p.state = "blocked"
    } else {
        fmt.Printf("Process %d enters CS\n", p.id)
        p.state = "critical"
    }
}

// ReleaseCS releases the critical section.
func (p *Process) ReleaseCS() {
    fmt.Printf("Process %d releases CS\n", p.id)

    // Send release messages to all processes in the quorum.
    for _, pid := range p.q {
        fmt.Printf("Process %d sends release message to %d\n", p.id, pid)
    }

    // Change state to ready.
    p.state = "ready"
}

func main() {
    // Create a set of processes.
    numProcs := 5
    processes := make([]*Process, numProcs)
    for i := 0; i < numProcs; i++ {
        processes[i] = NewProcess(i)
    }

    // Set the quorums for each process.
    for i := 0; i < numProcs; i++ {
        processes[i].q = []int{(i + 1) % numProcs, (i + 2) % numProcs}
    }

    // Run the simulation for 1000 iterations.
    for i := 0; i < 3; i++ {
        // Randomly select a process to request the critical section.
        pid := rand.Intn(numProcs)
        processes[pid].RequestCS()

        // Wait for all processes to finish with the critical section.
        for _, p := range processes {
            if p.state == "critical" {
                p.ReleaseCS()
            }
        }
    }
}
