# xua-parser

This project is a Go package for parsing a customized user agent HTTP header. It provides functionality to parse and structure the user agent string according to a specified pattern. The package is designed to handle different user agent formats and normalize them for consistent processing.

## Motivation

We built this package to simplify the parsing of customized `X-User-Agent` headers in our system, where multiple user agent formats from various platforms needed to be handled consistently. Existing libraries either lacked the flexibility for custom patterns or introduced unnecessary complexity, so we developed `xua-parser` to address these issues with a streamlined solution specific to our needs.

## Features

- Parse a user agent string into structured data (`UserAgent` struct).
- Includes unit tests and example code to validate functionality with various user agent strings.

## Usage

### Parsing a User Agent

To parse a user agent string:

```go
package main

import (
    "fmt"
    "log"
    "github.com/retail-ai-inc/xua-parser/ua"
)

func main() {
    userAgent := "jp.retailai.App/3.9.3 (Device-Model, Android 10, Other)"
    parsedAgent, err := ua.Parse(userAgent)
    if err != nil {
        log.Fatalf("Error parsing user agent: %v", err)
    }

    fmt.Printf("Parsed User Agent: %+v\n", parsedAgent)
}
```

This will output:

```bash
Parsed User Agent: &{AppName:jp.retailai.App AppVersion:3.9.3 DeviceModel:Device-Model OSName:Android OSVersion:10 Others:Other}
```

## Credits

Some parts of this package are inspired by or adapted from [mssola/useragent](https://github.com/mssola/useragent). We would like to acknowledge and give credit to the original developers of that project.

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue if you have any suggestions or find bugs.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
