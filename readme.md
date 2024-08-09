# Go Interpreter

A simple interpreter written in Go.

## Features

- Variable assignments
- Basic arithmetic operations
- Strings
- Conditional statements
- Arrays
- Hash Maps
- Functions and Higher-Order Functions
- Closures

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/pavank1301/Interpreter-with-Go
    ```

2. Navigate to the project directory:

    ```bash
    cd Prism
    ```

3. Start:

    ```bash
    go run main.go
    ```

## Usage

You can run the interpreter from the command line:

```bash
./Prism
```


## Example

```bash
> print("hello");
hello
> let x = "string";
> let x = 43;
> let people = [{"name": "Alice", "age": 24}, {"name": "Anna", "age": 28}];
> people[0]["name"];
Alice
> let myArray = ["Thorsten", "Ball", 28, fn(x) { x * x }];
> myArray[0];
Thorsten
> let fullName = fn(first, last) { first + " " + last };
> fullName("John", "Doe");
"John Doe"
```

# Supported Operations

## Arithmetic

```bash
> let a = 5;
> let b = 10;
> a + b;
15
> a * b;
50
> b / a;
2
```

## Strings

```bash
> let str = "Hello, ";
> let name = "World";
> str + name;
"Hello, World"
```

## Arrays

```bash
> let arr = [1, 2, 3, 4];
> arr[2];
3
```

## Hash Maps

```bash
> let hash = {"key1": "value1", "key2": "value2"};
> hash["key1"];
"value1"
```

## Built-in Functions for Arrays

```bash
> let arr = [1, 2, 3, 4];
> len(arr);
4
> first(arr);
1
> last(arr);
4
> rest(arr);
[2, 3, 4]
```

## Variable Assignments

```bash
> let x = 10;
> x = x + 5;
> x;
15
```

## Conditional Statements

```bash
> let num = 10;
> if (num > 5) { "Greater"; } else { "Smaller"; }
"Greater"
```

## Functions and Higher-Order Functions

```bash
> let add = fn(a, b) { a + b };
> add(5, 10);
15
> let apply = fn(f, x) { f(x) };
> apply(fn(x) { x * x }, 5);
25
```

## Closures

```bash
> let newAdder = fn(x) { fn(y) { x + y } };
> let addFive = newAdder(5);
> addFive(10);
15
```

# Project Structure

- main.go: Entry point of the interpreter.
- lexer/: Contains the lexical analyzer for the interpreter.
- parser/: Contains the parser for the interpreter.
- evaluator/: Contains the evaluation logic.
- ast/: Abstract Syntax Tree representation.
- object/: Object system for the interpreter.

# Contributing

Contributions are welcome! Please open an issue or submit a pull request.
  
