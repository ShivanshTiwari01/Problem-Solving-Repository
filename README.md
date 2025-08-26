# Practice Problems Repository

Welcome to my practice problems repository! Here, I share important coding problems along with Go implementations of common data structures and algorithms (DSA).

## About

This repository is a collection of important algorithms I've encountered during my coding journey. I've carefully selected unique problems to help others improve their problem-solving skills. Originally implemented in C++, this repository has been migrated to Go for simplicity, and modern development practices.

## Data Structures and Algorithms

- Go implementations of common data structures and algorithms, including:
  - Arrays and Slices
  - Linked Lists
  - Stacks
  - Queues
  - Trees (Binary Tree, BST)
  - Graphs (DFS, BFS)
  - Sorting Algorithms (Quick Sort, Merge Sort, etc.)
  - Searching Algorithms (Binary Search, etc.)
  - Dynamic Programming
  - Backtracking
  - And more...

## Repository Structure

```
DSA Implementation/
├── Go implementation of all important Data Structures and Algorithms

```

## Running the Code

Make sure you have Go installed (version 1.19 or later recommended):

```bash
# Run a specific implementation
go run DSA\ Implementation/arrays/reverse_array.go

# Run with go modules (recommended)
go mod init problem-solving
go run ./DSA\ Implementation/arrays/reverse_array.go

# Build an executable
go build -o reverse_array ./DSA\ Implementation/arrays/reverse_array.go
```

## Key Differences from C++ Version

- **Memory Management**: Automatic garbage collection eliminates manual memory management
- **Package System**: Better modularity with Go's package system
- **Built-in Data Structures**: Slices, maps, and channels as first-class citizens
- **Concurrency**: Goroutines for concurrent algorithm implementations where applicable
- **Testing**: Built-in testing framework with `_test.go` files

## Contributing

- Contributions are welcome! If you have a unique coding problem or want to improve existing solutions, feel free to open a pull request.
- Please follow Go conventions and include tests for new implementations.
- Run `go fmt` and `go vet` before submitting.

## License

This repository is licensed under the [MIT License](LICENSE).
