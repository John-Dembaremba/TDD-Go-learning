# Practising TDD concepts
I am learning TDD using the book "Test-Driven Development with Go."


## Introduction to TDD

Test-Driven Design (TDD) is a software development approach in which tests are written before writing the code that needs to be tested. The process follows a simple cycle: write a test, write code to pass the test, and then refactor the code.

## Benefits of TDD

- **Improved Code Quality**: Writing tests first ensures that the code meets the requirements from the start.
- **Faster Debugging**: Early detection of defects makes them easier to fix.
- **Better Design**: Encourages simple, modular, and reusable code.
- **Documentation**: Tests serve as documentation for how the code is supposed to work.

## TDD Cycle

The TDD cycle consists of three main steps, often referred to as **Red-Green-Refactor**:

1. **Red**: Write a test that fails.
2. **Green**: Write the minimum amount of code to make the test pass.
3. **Refactor**: Refactor the code to improve its structure and readability without changing its behavior.

## Setting Up Your Go Environment

1. **Install Go**: Download and install Go from the [official website](https://golang.org/dl/).
2. **Set up your workspace**: Create a directory for your Go projects. Set the `GOPATH` environment variable to this directory.
3. **Install a code editor**: Use an editor like VSCode or GoLand that supports Go.

## Writing Your First Test

1. **Create a Go module**:
    ```sh
    mkdir tdd-example
    cd tdd-example
    go mod init tdd-example
    ```

2. **Write a test file**:
    ```go
    // calculator_test.go
    package main

    import (
        "testing"
    )

    func TestAdd(t *testing.T) {
        result := Add(2, 3)
        expected := 5
        if result != expected {
            t.Errorf("Expected %d but got %d", expected, result)
        }
    }
    ```

3. **Run the test**:
    ```sh
    go test
    ```

## Implementing the Code

1. **Write the implementation to pass the test**:
    ```go
    // calculator.go
    package main

    func Add(a, b int) int {
        return a + b
    }
    ```

2. **Run the test again**:
    ```sh
    go test
    ```

## Refactoring

Refactor your code to improve its structure and maintainability. Ensure that all tests still pass after refactoring.

## Best Practices

- **Write small tests**: Each test should verify a single behavior.
- **Keep tests independent**: Tests should not depend on each other.
- **Use meaningful names**: Both for tests and functions.
- **Test edge cases**: Consider edge cases and error conditions.

## Resources

- [Go Official Documentation](https://golang.org/doc/)
- [Test-Driven Development by Example by Kent Beck](https://www.amazon.com/Test-Driven-Development-Kent-Beck/dp/0321146530)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Testing Package](https://golang.org/pkg/testing/)

