# Hello Golang

## Concept Goal
The goal of this project is to understand the absolute basics of a Go application: defining an entry point, using the standard library for output, and importing and calling an external third-party package for the first time.

## Exercise Requirements
To recreate this project from scratch as a learning exercise:

1. Create a `main.go` file inside this folder.
2. Initialize the file with the strict `main` package declaration.
3. Import the standard library `"fmt"` package.
4. Import the external package `"github.com/google/uuid"`.
   > **Hint:** You will need to fetch this dependency using your terminal before it will build! (`go get github.com/google/uuid`)
5. Create the primary `main()` function entrypoint.
6. Print the text `"Hello Golang"` to the standard output.
7. Generate a new UUID using the imported `uuid` package, convert it to a string, and print it below the greeting (e.g. `"UUID:  <generated-uuid>"`).

## How to Run
Once you believe you have successfully met the requirements, test your code securely across any terminal by running:
```bash
cd 1_hello-golang && go run .
```
