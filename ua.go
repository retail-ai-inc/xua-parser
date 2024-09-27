// Package xua_parser provides a simple parser for User-Agent strings.
package xua_parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
`[AppName/Version] (DeviceModel, OSDetail, Retailer)`
  - AppName:     Any characters except a slash.
  - AppVersion:  One to three groups of digits separated by periods (e.g., 1, 1.0, 1.0.0).
  - DeviceModel: Any characters except a comma.
  - OSDetail:    OS name, version, and other details except a comma and the first two fields are separated by a space.
  - Others:      Any characters except a closing parenthesis.
*/
var uaRegex = regexp.MustCompile(`^([^/]+)/([\d]+(?:\.[\d]+(?:\.[\d]+)?)?) \(([^,]+), ([^,]+), ([^\)]+)\)$`)

// Parse parses the given user agent string and returns the app-user-agent info if parsed successfully.
func Parse(ua string) (*UserAgent, error) {
	matches := uaRegex.FindStringSubmatch(ua)

	if len(matches) != 6 {
		return nil, fmt.Errorf("failed to parse ua: %s", ua)
	}

	// Split OS field into OSName and OSVersion
	osDetail := strings.Split(matches[4], " ")
	if len(osDetail) < 2 {
		return nil, fmt.Errorf("os detail does not contain name and version: %s", matches[4])
	}
	osName := osDetail[0]
	osVersion := osDetail[1]
	// validate os version is a number
	if _, err := strconv.Atoi(osVersion); err != nil {
		return nil, fmt.Errorf("os version is not a number: %s", osVersion)
	}

	return &UserAgent{
		AppName:     matches[1],
		AppVersion:  matches[2],
		DeviceModel: matches[3],
		OSName:      osName,
		OSVersion:   osVersion,
		Others:      matches[5],
	}, nil
}
