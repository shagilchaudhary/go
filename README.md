## Author

Name - Prasad Nijsure

UL Student No - 21004528

# Parallel Sorting Algorithms

Heap Sort

![Exercise: Heap Sort Working Sequential]("Heap Sort Working Sequential")

Merge Sort

![Exercise: Merge Sort Working Sequential]("Merge Sort Working Sequential")

## Benchmark Results

| Algorithm  | Max CPUs | InputSize | Sequential Processing Time | Parallel Processing Time | Equality Check
| ------------- | ------------- |
| Heap Sort  | 4 | 1000 | 126.453µs | 1.394877ms | true
| Merge Sort  | 4 | 1000 | 118.847µs | 595.686µs | true  

## References


## Build

Install Golang locally (https://golang.org/doc/install) and also set the environment variables

## Run

GO111MODULE=off is made off during program execution after reading [Reference](https://insujang.github.io/2020-04-04/go-modules/ "GO111MODULE=off reference link")

The arguement is the input size

```
GO111MODULE=off go run main.go 1000
```

## TODO

Writing Golang unit tests