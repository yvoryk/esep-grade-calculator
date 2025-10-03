# ESEP Grade Calculator

A Go library for calculating final course grades based on weighted averages of assignments, exams, and essays.

## Overview

This library computes final grades by:
1. Calculating the average grade for each grade type (assignments, exams, essays)
2. Computing a weighted average using configurable weights
3. Converting the numerical grade to a letter grade (A-F) or Pass/Fail

**Default Weights:**
- Assignments: 50%
- Exams: 35%
- Essays: 15%

**Default Grade Scale:**
- A: ≥ 90
- B: ≥ 80
- C: ≥ 70 (minimum passing)
- D: ≥ 60
- F: < 60

## Requirements

**Runtime:** Go 1.24+

## Installation

```bash
go get esep/grade-calculator
```

## Usage

```go
import "esep/grade-calculator"

// Create a new grade calculator
gc := esepunittests.NewGradeCalculator()

// Add grades
gc.AddGrade("Homework 1", 95, esepunittests.Assignment)
gc.AddGrade("Midterm Exam", 88, esepunittests.Exam)
gc.AddGrade("Research Paper", 92, esepunittests.Essay)

// Get final letter grade
finalGrade := gc.GetFinalGrade() // Returns "A", "B", "C", "D", or "F"
```

## Running Tests

Run all unit tests:
```bash
cd grade-calculator/
go test ./...
```

Run tests with verbose output:
```bash
cd grade-calculator/
go test -v ./...
```

Run tests with code coverage:
```bash
cd grade-calculator/
go test -coverprofile=coverage.txt ./...
go tool cover -html=coverage.txt  # View coverage in browser
```

## Project Structure

```
esep-grade-calculator/
├── README.MD                      # This file
└── grade-calculator/
    ├── go.mod                     # Go module definition
    ├── grade_calculator.go        # Main implementation
    └── grade_calculator_test.go   # Unit tests
```

## Development

### Running Tests During Development

```bash
cd grade-calculator/
go test -v ./...
```

### Checking Code Coverage

```bash
cd grade-calculator/
go test -cover ./...
```

## API Reference

### Types

- `GradeType`: Enum for grade types (Assignment, Exam, Essay)
- `Grade`: Struct containing Name (string), Grade (int), and Type (GradeType)
- `GradeCalculator`: Main calculator struct

### Functions

- `NewGradeCalculator() *GradeCalculator`: Creates a new grade calculator instance
- `AddGrade(name string, grade int, gradeType GradeType)`: Adds a grade to the calculator
- `GetFinalGrade() string`: Returns the final letter grade (A-F)
