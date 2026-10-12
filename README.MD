# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

The grade calculator stores assignment across three categories: **Assignments (50%)**, **Exams (35%)**, and **Essays (15%)**. Each category is calculated independently but then later computed using a weighted average across all categories. Where, the final grade is calculated using:

Final Grade = (Assignment Average * Assignment Weight) + (Exam Average * Exam Weight) + (Essay Average * Essay Weight)

Finally, grades are assigned based on the following thresholds:

Score  | Grade
-------|-------
`>= 90`| A
`>= 80`| B
`>= 70`| C
`>= 60`| F

Assignment category weights and grade thresholds can also be changed, but the ones provided are for the current example.

## Tech Stack

Runtime requirements:
Golang 1.24+

This project uses Golang 1.24.4. Additionally, any version above 1.24.4 is sufficent.

Install Golang here: https://go.dev/doc/install

Verify your Golang version using
```
go version
```

Golang also manages dependency management and the test environment used.

## File Summary
`README.MD`:
Summary of the project's code, tests, and structure. Covers technical requirements and how to run test cases.

`go.mod`: defines the module path to identify packages within the project. (ie. the grade calculator) and Golang version used.

`grade_calculator.go`: defines the implementation of the grade calculator, such as class object structure and calculation operations.

`grade_calculator_test.go` holds all of the test cases for the grade calculator implementation. 

## Running Test Cases
Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```