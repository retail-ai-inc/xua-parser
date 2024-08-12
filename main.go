// main.go
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

    formattedAgent := parsedAgent.FormatUserAgent()
    fmt.Printf("Formatted User Agent: %s\n", formattedAgent)
}

