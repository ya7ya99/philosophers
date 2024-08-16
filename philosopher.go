package main

import (
	"sync"
	"time"
)

type Philosopher struct {
	id        int
	leftFork  *sync.Mutex
	rightFork *sync.Mutex
	eatCount  int
	lastMeal  time.Time
}

func NewPhilosopher(id int, left, right *sync.Mutex) *Philosopher {
	return &Philosopher{
		id:        id,
		leftFork:  left,
		rightFork: right,
		lastMeal:  time.Now(),
	}
}

func (p *Philosopher) eat(table *Table) {
	p.leftFork.Lock()
	p.rightFork.Lock()

	table.logState(p.id, "is eating")
	time.Sleep(time.Duration(table.timeToEat) * time.Millisecond)

	p.eatCount++
	p.lastMeal = time.Now()

	p.rightFork.Unlock()
	p.leftFork.Unlock()
}

func (p *Philosopher) think(table *Table) {
	table.logState(p.id, "is thinking")
	time.Sleep(time.Duration(table.timeToSleep) * time.Millisecond)
}

func (p *Philosopher) sleep(table *Table) {
	table.logState(p.id, "is sleeping")
	time.Sleep(time.Duration(table.timeToSleep) * time.Millisecond)
}
