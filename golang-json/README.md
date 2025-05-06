# Golang JSON Course Project

A demonstration of JSON handling capabilities in Go, showcasing various techniques for encoding, decoding, and manipulating JSON data.

## Overview

This project serves as a practical guide to working with JSON in Go, using the standard `encoding/json` package. It demonstrates common JSON operations through executable test files.

## Requirements

- Go 1.24 or later

## Project Structure

- `go.mod` - Module definition and dependencies
- `customer.json` - Sample JSON data for customer objects
- `customer_out.json` - Output file for JSON encoding operations
- Test files demonstrating various JSON functionalities:
  - `decode_test.go` - JSON decoding operations
  - `encode_json_test.go` - JSON encoding operations
  - `json_array_test.go` - Working with JSON arrays
  - `json_object_test.go` - Working with JSON objects
  - `json_tag_test.go` - Custom field naming with struct tags
  - `map_test.go` - Dynamic JSON handling with maps
  - `stream_test.go` - Streaming JSON operations

## Features

### Struct Mapping

Define Go structs that map to JSON data:

```go
type Customer struct {
    FirstName  string
    MiddleName string
    LastName   string
    Age        int
    Married    bool
}
```

### Custom Field Names with Tags

Use struct tags to customize JSON field names:

```go
type Product struct {
    Id       string `json:"id"`
    Name     string `json:"name"`
    Price    int    `json:"price"`
    ImageURL string `json:"image_url"`
}
```

### Basic Encoding/Decoding

```go
// Decoding JSON to struct
json.Unmarshal([]byte(jsonString), &customer)

// Encoding struct to JSON
bytes, _ := json.Marshal(customer)
```

### Dynamic JSON with Maps

Work with JSON when you don't know the structure in advance:

```go
var data map[string]interface{}
json.Unmarshal([]byte(jsonString), &data)
```

### File Operations

Read from and write to JSON files:

```go
// Reading JSON from file
bytes, _ := os.ReadFile("customer.json")
json.Unmarshal(bytes, &customer)

// Writing JSON to file
bytes, _ := json.Marshal(customer)
os.WriteFile("customer_out.json", bytes, 0644)
```

### JSON Streaming

Process large JSON files efficiently with streaming:

```go
file, _ := os.Open("large_file.json")
decoder := json.NewDecoder(file)
for decoder.More() {
    decoder.Decode(&item)
    // Process item
}
```

## Running Tests

To run a specific test file:

```bash
go test -v json_tag_test.go
```

To run all tests:

```bash
go test -v ./...
```

## License

This project is created for educational purposes as part of a Go programming course.