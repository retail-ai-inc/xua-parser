
# RAI Agent Parser

This project is a Go package for parsing RAI customized `X-User-Agent` HTTP header. It provides functionality to parse and format user agent strings according to a specified pattern. The package is designed to handle different platforms' user agent formats and normalize them for consistent processing.

## Features

- Parse `X-User-Agent` strings into structured data.
- Format structured data back into a `X-User-Agent` string.
- Includes unit tests to validate functionality with various user agent strings.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/useragent-parser.git
   cd useragent-parser
   ```

2. Initialize the Go module:

   ```bash
   go mod init myproject
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

## Usage

### Parsing a User Agent

To parse a `X-User-Agent` string, use the `ParseUserAgent` function:

```go
package main

import (
    "fmt"
    "log"
    "myproject/useragent"
)

func main() {
    xUserAgent := "jp.retailai.raicart/3.9.3 (S-500, Android 10, trial)"
    parsedAgent, err := useragent.ParseUserAgent(xUserAgent)
    if err != nil {
        log.Fatalf("Error parsing user agent: %v", err)
    }

    fmt.Printf("Parsed User Agent: %+v\n", parsedAgent)
}
```

### Formatting a User Agent

To format a `UserAgent` struct back into a string, use the `FormatUserAgent` method:

```go
func main() {
    ua := useragent.UserAgent{
        AppName:     "jp.retailai.raicart",
        VersionName: "3.9.3",
        DeviceModel: "S-500",
        OSVersion:   "Android 10",
        Retailer:    "trial",
    }

    formatted := ua.FormatUserAgent()
    fmt.Println("Formatted User Agent:", formatted)
}
```

## Project Structure

```
myproject/
├── go.mod
├── main.go
└── useragent/
    ├── useragent.go
    └── useragent_test.go
```

- `main.go`: Example usage of the `useragent` package.
- `useragent/useragent.go`: The main package that includes the logic for parsing and formatting user agents.
- `useragent/useragent_test.go`: Unit tests for the `useragent` package.

## Running Tests

To run the unit tests:

```bash
go test ./useragent
```

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue if you have any suggestions or find any bugs.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
