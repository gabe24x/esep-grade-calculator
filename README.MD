# ESEP Grade Calculator

A tiny Go library that computes final course grades from assignments, exams, and essays.

## Runtime

- Go 1.24+ (as specified in `go.mod`). Any recent Go 1.x should work.

## Project Layout

- `go.mod` – module name and Go version
- `grade_calculator.go` – library code (`GradeCalculator`, `Grade`, `GradeType`)
- `grade_calculator_test.go` – tests

## How Final Grades Are Computed

1. Compute the **average** for each category:
   - Assignments (50%)
   - Exams (35%)
   - Essays (15%)
2. Compute the **weighted** sum of those averages.
3. Convert to a **letter** grade using:
   - A ≥ 90
   - B ≥ 80
   - C ≥ 70
   - D ≥ 60
   - F < 60

> Empty categories are treated as average = 0.  
> Numeric → letter uses integer conversion (floors the weighted average).

## Run Tests

```bash
cd grade_calculator
go test ./...
