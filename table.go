package main

import (
	"fmt"
	"sync"
	"time"
)

type Table struct {
	philosophersCount int
	timeToDie         int
	timeToEat         int
	timeToSleep       int
	mustEatCount      int
	forks             []*sync.Mutex
	logMutex          sync.Mutex
}

func NewTable(philoCount, timeToDie, timeToEat, timeToSleep, mustEatCount int) *Table {
	forks := make([]*sync.Mutex, philoCount)
	for i := range forks {
		forks[i] = &sync.Mutex{}
	}
	return &Table{
		philosophersCount: philoCount,
		timeToDie:         timeToDie,
		timeToEat:         timeToEat,
		timeToSleep:       timeToSleep,
		mustEatCount:      mustEatCount,
		forks:             forks,
	}
}

func (t *Table) startSimulation() {
	philosophers := make([]*Philosopher, t.philosophersCount)
	for i := 0; i < t.philosophersCount; i++ {
		leftFork := t.forks[i]
		rightFork := t.forks[(i+1)%t.philosophersCount]
		philosophers[i] = NewPhilosopher(i+1, leftFork, rightFork)
	}

	var wg sync.WaitGroup
	wg.Add(t.philosophersCount)

	for _, p := range philosophers {
		go func(philo *Philosopher) {
			defer wg.Done()
			for {
				philo.think(t)
				philo.eat(t)
				philo.sleep(t)

				if t.mustEatCount > 0 && philo.eatCount >= t.mustEatCount {
					return
				}

				if time.Since(philo.lastMeal).Milliseconds() > int64(t.timeToDie) {
					t.logState(philo.id, "died")
					return
				}
			}
		}(p)
	}

	wg.Wait()
}

func (t *Table) logState(id int, state string) {
	t.logMutex.Lock()
	defer t.logMutex.Unlock()
	fmt.Printf("Philosophe %d %s\n", id, state)
}
