
```markdown
# Philosopher Dining Simulation

This project simulates the classic "Dining Philosophers" problem, a concurrency problem that demonstrates the challenges of avoiding deadlock and ensuring fairness among competing processes.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Overview

In this simulation, a group of philosophers sit around a table with a fork placed between each pair of adjacent philosophers. Each philosopher needs both forks to eat, but can only pick up one fork at a time. The philosophers alternate between thinking, eating, and sleeping. If a philosopher does not eat within a certain amount of time, they die.

This project uses Go's concurrency primitives to simulate the actions of each philosopher and manage access to shared resources (forks).

## Features

- Simulates the Dining Philosophers problem with adjustable parameters.
- Allows configuration of the number of philosophers, time to eat, sleep, and maximum time before a philosopher dies.
- Implements concurrency using goroutines and mutexes to manage shared resources.

## Installation

To get started with this project, make sure you have [Go installed](https://golang.org/dl/) on your machine.

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/philosopher-dining-simulation.git
   ```
2. Navigate to the project directory:
   ```bash
   cd philosopher-dining-simulation
   ```
3. Build the project:
   ```bash
   go build
   ```

## Usage

To run the simulation, use the following command:

```bash
go run . number_of_philosophers time_to_die time_to_eat time_to_sleep [number_of_times_each_philosopher_must_eat]
```

### Parameters

- `number_of_philosophers`: The number of philosophers sitting at the table.
- `time_to_die`: The maximum time (in milliseconds) a philosopher can live without eating.
- `time_to_eat`: The time (in milliseconds) it takes for a philosopher to eat.
- `time_to_sleep`: The time (in milliseconds) it takes for a philosopher to sleep.
- `number_of_times_each_philosopher_must_eat`: (Optional) The number of times each philosopher needs to eat before the simulation ends.

### Example

To run a simulation with 5 philosophers, a `time_to_die` of 800ms, `time_to_eat` of 200ms, and `time_to_sleep` of 100ms:

```bash
go run . 5 800 200 100
```

## Configuration

You can adjust the parameters mentioned above to customize the simulation according to your needs.

## Contributing

Contributions are welcome! Please fork this repository, make your changes, and submit a pull request. 

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```
