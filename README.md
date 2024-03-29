# Gofuzzer
This is a simple Go fuzzer designed to read directories from a file, send HTTP requests to specified URLs, and report the status codes and response lengths. It can be used for web directory fuzzing and testing.

## Prerequisites

Before running the Go fuzzer, ensure you have the following installed:

- Go (Golang): [Download Go](https://golang.org/dl/)

## Getting Started

1. Clone or download this repository to your local machine.
2. You can replace the URL and directory name in the code
3. Go in to the directory and go build ```go build``` to build the package on your machine
4. Run the Go fuzzer by executing the following command:

      ```
       gofuzz or go run main.go
      ```



  



## Example output

```
URL: https://example.com/directory1
Status Code: 200
Response Length: 1234 bytes

URL: https://example.com/directory2
Status Code: 404
Response Length: 0 bytes

```