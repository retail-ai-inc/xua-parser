// Package xua_parser provides a simple parser for User-Agent strings.
package xua_parser

import (
	"fmt"
	"regexp"
)

// UserAgent represents the parsed User-Agent information.
type UserAgent struct {
	AppName     string
	AppVersion  string
	DeviceModel string
	OSName      string
	OSVersion   string
	Others      string
}

/*
uaRegex extracts the following information from the User-Agent string:
`[AppName/Version] (DeviceModel, OSName OSVersion, Retailer)`
  - AppName:     Any characters except a slash.
  - AppVersion:  One to three groups of digits separated by periods (e.g., 1, 1.0, 1.0.0).
  - DeviceModel: Any characters except a comma.
  - OSName:      One or more letters.
  - OSVersion:   One or more digits.
  - Others:      Any characters except a closing parenthesis.
*/
var uaRegex = regexp.MustCompile(`^([^/]+)/([\d]+(?:\.[\d]+(?:\.[\d]+)?)?) \(([^,]+), ([a-zA-Z]+) (\d+), ([^\)]+)\)$`)

// Parse parses the given user agent string and returns the app-user-agent info if parsed successfully.
func Parse(ua string) (*UserAgent, error) {

	matches := uaRegex.FindStringSubmatch(ua)

	if len(matches) != 7 {
		return nil, fmt.Errorf("failed to parse ua: %s", ua)
	}

	return &UserAgent{
		AppName:     matches[1],
		AppVersion:  matches[2],
		DeviceModel: matches[3],
		OSName:      matches[4],
		OSVersion:   matches[5],
		Others:      matches[6],
	}, nil
}
