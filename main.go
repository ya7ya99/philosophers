package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {

	if len(os.Args) < 5 || len(os.Args) > 6 {
        fmt.Println("Usage: go run . number_of_philosophers time_to_die time_to_eat time_to_sleep [number_of_times_each_philosopher_must_eat]")
        return
    }


	args := os.Args[1:]


	philoCount, _ := strconv.Atoi(args[0])
    timeToDie, _ := strconv.Atoi(args[1])
    timeToEat, _ := strconv.Atoi(args[2])
    timeToSleep, _ := strconv.Atoi(args[3])
    
    var mustEatCount int
    if len(args) == 5 {
        mustEatCount, _ = strconv.Atoi(args[4])
    }


	table := NewTable(philoCount, timeToDie, timeToEat, timeToSleep, mustEatCount)
    
    table.startSimulation()
}