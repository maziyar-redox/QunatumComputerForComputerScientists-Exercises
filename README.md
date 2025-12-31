
# 🧮 Quantum Computing Solutions

A collection of Golang implementations for programming drills from "Quantum Computing for Computer Scientists" by Mirco A. Mannucci and Noson S. Yanofsky.

## 📚 About the Book

**Title**: Quantum Computing for Computer Scientists

**Authors**: Mirco A. Mannucci and Noson S. Yanofsky

**Focus**: Bridging the gap between traditional computer science concepts and quantum computing fundamentals.

This repository contains my solutions to the programming exercises and drills from the book, implemented in **Go (Golang)**.

## 📁 Project Structure

```sh
quantum-computing-drills/
│
├── README.md                    # This file
├── go.mod                       # Go module file
├── LICENSE                      # MIT License
│
├── pkg/                         # Core quantum computing packages
│   ├── complex/                 # Complex number utilities
│   ├── matrix/                  # Matrix operations for quantum states
│   ├── qubit/                   # Qubit representation and operations
│   ├── gates/                   # Quantum gate implementations
│   └── circuits/                # Quantum circuit simulation
│
├── chapters/                    # Solutions organized by book chapters
│   ├── chapter01/              # Introduction & Mathematical Preliminaries
│   │   ├── main.go
│   │   ├── exercises.go
│   │   └── README.md
│   ├── chapter02/              # Quantum Mechanics Basics
│   ├── chapter03/              # Quantum Bits (Qubits)
│   ├── chapter04/              # Quantum Gates
│   ├── chapter05/              # Quantum Circuits
│   ├── chapter06/              # Quantum Algorithms
│   └── ...
│
├── examples/                    # Complete working examples
│   ├── bell_state/             # Bell state creation
│   ├── deutsch_jozsa/          # Deutsch-Jozsa algorithm
│   ├── grovers_search/         # Grover's search algorithm
│   └── quantum_fourier/        # Quantum Fourier Transform
│
├── tests/                       # Unit tests for all implementations
│   ├── complex_test.go
│   ├── matrix_test.go
│   ├── qubit_test.go
│   └── gates_test.go
│
└── docs/                        # Documentation and notes
    ├── concepts.md
    ├── book_notes.md
    └── algorithms.md
```

## 🧬 Implementation Details

### Core Quantum Concepts in Go

1. Complex Numbers (Chapter 1)

```sh
// Complex number representation
type Complex struct {
    Real float64
    Imag float64
}

// Operations: addition, multiplication, conjugation, norm
```

2. Qubit Representation (Chapter 3)

```sh
type Qubit struct {
    Alpha complex128  // |0⟩ coefficient
    Beta  complex128  // |1⟩ coefficient
}

// Methods: measure, applyGate, tensorProduct
```

3. Quantum Gates (Chapter 4)

```sh
// Pauli gates, Hadamard, CNOT, etc.
type Gate interface {
    Apply(Qubit) Qubit
    Matrix() [][]complex128
}
```

4. Quantum Circuits (Chapter 5)

type Circuit struct {
    Qubits   []Qubit
    Gates    []Gate
    Steps    int
}

## 🚀 Getting Started

### Prerequisites

- Go 1.20+ installed
- Basic understanding of linear algebra
- Familiarity with complex numbers

### Installation

```sh
# Clone the repository
git clone https://github.com/maziyar-redox/QuantumComputerForComputerScientists-Exercises.git
cd QuantumComputerForComputerScientists-Exercises

# Install dependencies
go mod download

# Run tests to verify installation
go test ./...
```

### Running Examples

```sh
# Run a specific chapter's solutions
cd chapters/chapter03
go run main.go

# Run a complete algorithm example
cd examples/bell_state
go run main.go
```

## 📖 Chapter-wise Solutions

### Chapter 1: Mathematical Preliminaries

- Complex number operations
- Linear algebra basics
- Matrix multiplication
- Inner products and norms

### Chapter 2: Quantum Mechanics Basics

- State vectors
- Observables
- Measurement
- Tensor products

### Chapter 3: Quantum Bits

- Qubit representation
- Bloch sphere
- Single qubit operations
- Computational basis

### Chapter 4: Quantum Gates

- Pauli matrices
- Hadamard gate
- Phase gates
- Controlled gates

### Chapter 5: Quantum Circuits

- Circuit diagrams
- Gate sequences
- Multiple qubit systems
- Circuit optimization

### Chapter 6: Quantum Algorithms

- Deutsch-Jozsa algorithm
- Bernstein-Vazirani
- Simon's algorithm
- Grover's search

*(Additional chapters will be added as progress continues)*

## 🧪 Testing

Run comprehensive tests to verify quantum operations:

```sh
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./pkg/qubit

# Run benchmark tests
go test -bench=. ./pkg/matrix
```

## 📈 Progress Tracking

| Chapter       | Status   | Exercises Completed | 
| ----------------- | --------------------------|---- |
| 1 |  ✅ Complete| 15/15|
| 2 |  🔄 In Progress| 15/15|
|  3|  ⏳ Pending| 15/15|
|  4|  ✅ Complete| 15/15|


## 🤝 Contributing

While this is primarily a personal learning repository, suggestions are welcome:

1. Issues: Report bugs or suggest improvements
2. Discussion: Share alternative approaches or optimizations
3. Educational Value: Focus on clarity and learning resources

*Note: Please don't submit direct solutions to book exercises.*

## 📄 License
MIT License - See [LICENSE](https://github.com/maziyar-redox/QunatumComputerForComputerScientists-Exercises/blob/main/LICENSE) file for details.

All code implementations are original work by Mr_Red0x.

Book exercises copyright © by Mirco A. Mannucci and Noson S. Yanofsky.

*Last Updated: 12/31/2025*