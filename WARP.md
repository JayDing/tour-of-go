# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project Overview

This is a Go learning project (`tour-of-go`) that demonstrates fundamental Go programming concepts. The repository contains a single main package with example functions showcasing:

- Basic function declarations and multiple return values
- Named return parameters  
- Mathematical computations (including a Newton's method implementation for square root)
- String manipulation and variable swapping

The project structure is intentionally simple for learning purposes, with all code contained in `main.go`.

## Development Commands

### Build and Run
```bash
# Run the program directly
go run main.go

# Build executable
go build -o tour-of-go main.go

# Run the built executable
./tour-of-go
```

### Testing
```bash
# Run tests (currently no test files exist)
go test -v ./...
```

### Code Quality
```bash
# Format Go code
go fmt ./...

# Check for potential issues
go vet ./...

# Run Go modules commands
go mod tidy    # Clean up dependencies
go mod verify  # Verify dependencies
```

## Architecture Notes

- **Single Package Structure**: All code is in the `main` package within `main.go`
- **Educational Focus**: Functions demonstrate progressively complex Go concepts:
  - `add()`: Basic function with typed parameters
  - `swap()`: Multiple return values
  - `split()`: Named return parameters
  - `sqrt()`: Iterative algorithm with mathematical precision
- **Module Name**: `first-project` (Go 1.25)
- **No External Dependencies**: Uses only standard library (`fmt`, `math`)

## Key Files

- `main.go`: Contains all example functions and the main execution logic
- `go.mod`: Module definition with Go version 1.25
- `README.md`: Minimal project description
- `tour-of-go.code-workspace`: VS Code workspace configuration

## Development Notes

- This is a learning repository focused on Go fundamentals
- The Newton's method square root implementation demonstrates floating-point precision handling
- Code output shows function results: addition, string swapping, integer splitting, and square root comparison
- No tests are currently implemented but could be added for educational purposes
