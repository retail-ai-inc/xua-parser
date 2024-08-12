// useragent/useragent_test.go
package useragent

import (
	"testing"
)

func TestParseUserAgent(t *testing.T) {
	testCases := []struct {
		input    string
		expected UserAgent
		hasError bool
	}{
		{
			input: "jp.retailai.raicart/3.9.3 (S-500, Android 10, trial)",
			expected: UserAgent{
				AppName:     "jp.retailai.raicart",
				VersionName: "3.9.3",
				DeviceModel: "S-500",
				OSVersion:   "Android 10",
				Retailer:    "trial",
			},
			hasError: false,
		},
		{
			input: "SCT/1.0.0 (Toshiba T380, Windows 11, Aeon)",
			expected: UserAgent{
				AppName:     "SCT",
				VersionName: "1.0.0",
				DeviceModel: "Toshiba T380",
				OSVersion:   "Windows 11",
				Retailer:    "Aeon",
			},
			hasError: false,
		},
		{
			input: "ISM Flutter/2.3.1 (Samsung S8 Pro, Android 11, Sugi)",
			expected: UserAgent{
				AppName:     "ISM Flutter",
				VersionName: "2.3.1",
				DeviceModel: "Samsung S8 Pro",
				OSVersion:   "Android 11",
				Retailer:    "Sugi",
			},
			hasError: false,
		},
		{
			input: "SC Windows/4.0.0 (HP EliteBook, Windows 10, TRIAL)",
			expected: UserAgent{
				AppName:     "SC Windows",
				VersionName: "4.0.0",
				DeviceModel: "HP EliteBook",
				OSVersion:   "Windows 10",
				Retailer:    "TRIAL",
			},
			hasError: false,
		},
		{
			input: "jp.retailai.raicart/4.1.0 (Huawei P30, Android 9, Taiyo)",
			expected: UserAgent{
				AppName:     "jp.retailai.raicart",
				VersionName: "4.1.0",
				DeviceModel: "Huawei P30",
				OSVersion:   "Android 9",
				Retailer:    "Taiyo",
			},
			hasError: false,
		},
		{
			input: "ISM/3.2.1 (Dell XPS 13, Windows 10, Aeon)",
			expected: UserAgent{
				AppName:     "ISM",
				VersionName: "3.2.1",
				DeviceModel: "Dell XPS 13",
				OSVersion:   "Windows 10",
				Retailer:    "Aeon",
			},
			hasError: false,
		},
		{
			input: "SCT Flutter/1.1.1 (Sony Xperia, Android 8, TRIAL)",
			expected: UserAgent{
				AppName:     "SCT Flutter",
				VersionName: "1.1.1",
				DeviceModel: "Sony Xperia",
				OSVersion:   "Android 8",
				Retailer:    "TRIAL",
			},
			hasError: false,
		},
		{
			input: "SC Windows/2.5.0 (Lenovo ThinkPad, Windows 8, Taiyo)",
			expected: UserAgent{
				AppName:     "SC Windows",
				VersionName: "2.5.0",
				DeviceModel: "Lenovo ThinkPad",
				OSVersion:   "Windows 8",
				Retailer:    "Taiyo",
			},
			hasError: false,
		},
		{
			input: "ISM Flutter/3.0.0 (Samsung Galaxy S10, Android 10, Sugi)",
			expected: UserAgent{
				AppName:     "ISM Flutter",
				VersionName: "3.0.0",
				DeviceModel: "Samsung Galaxy S10",
				OSVersion:   "Android 10",
				Retailer:    "Sugi",
			},
			hasError: false,
		},
		{
			input: "SCT/2.2.2 (Google Pixel, Android 11, TRIAL)",
			expected: UserAgent{
				AppName:     "SCT",
				VersionName: "2.2.2",
				DeviceModel: "Google Pixel",
				OSVersion:   "Android 11",
				Retailer:    "TRIAL",
			},
			hasError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual, err := ParseUserAgent(tc.input)
			if tc.hasError {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if actual == nil {
					t.Errorf("expected %+v, got nil", tc.expected)
				} else if *actual != tc.expected {
					t.Errorf("expected %+v, got %+v", tc.expected, actual)
				}
			}
		})
	}
}

