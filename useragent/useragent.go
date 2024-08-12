// useragent/useragent.go
package useragent

import (
	"fmt"
	"regexp"
)

// UserAgent represents the parsed User-Agent information.
type UserAgent struct {
	AppName     string
	VersionName string
	DeviceModel string
	OSVersion   string
	Retailer    string
}

// ParseUserAgent parses the given X-User-Agent string and returns a UserAgent struct.
func ParseUserAgent(userAgent string) (*UserAgent, error) {
	// Regular expression to match the pattern
	re := regexp.MustCompile(`^([\w\s\.]+)/([\d\.]+) \(([^,]+), ([^,]+), ([^\)]+)\)$`)
	matches := re.FindStringSubmatch(userAgent)

	if len(matches) != 6 {
		return nil, fmt.Errorf("invalid user agent format: %s", userAgent)
	}

	return &UserAgent{
		AppName:     matches[1],
		VersionName: matches[2],
		DeviceModel: matches[3],
		OSVersion:   matches[4],
		Retailer:    matches[5],
	}, nil
}

// FormatUserAgent formats the UserAgent struct into a string according to the X-User-Agent format.
func (ua *UserAgent) FormatUserAgent() string {
	return fmt.Sprintf("%s/%s (%s, %s, %s)", ua.AppName, ua.VersionName, ua.DeviceModel, ua.OSVersion, ua.Retailer)
}

