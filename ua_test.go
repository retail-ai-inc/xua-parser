package xua_parser_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	ua "github.com/retail-ai-inc/xua-parser"
)

func Test_Parse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    *ua.UserAgent
		wantErr bool
	}{
		{
			name:  "com.example.app Android Device-Model Other",
			input: "com.example.app/3.9.3 (Device-Model, Android 10, Other)",
			want: &ua.UserAgent{
				AppName:     "com.example.app",
				AppVersion:  "3.9.3",
				DeviceModel: "Device-Model",
				OSName:      "Android",
				OSVersion:   "10",
				Others:      "Other",
			},
			wantErr: false,
		},
		{
			name:  "App 1 Android Device Model 1 Other",
			input: "App 1/2.3.1 (Device Model 1, Android 11, Other)",
			want: &ua.UserAgent{
				AppName:     "App 1",
				AppVersion:  "2.3.1",
				DeviceModel: "Device Model 1",
				OSName:      "Android",
				OSVersion:   "11",
				Others:      "Other",
			},
			wantErr: false,
		},
		{
			name:  "App Windows Windows Device Model Other",
			input: "App Windows/4.0.0 (Device Model, Windows 10, Other)",
			want: &ua.UserAgent{
				AppName:     "App Windows",
				AppVersion:  "4.0.0",
				DeviceModel: "Device Model",
				OSName:      "Windows",
				OSVersion:   "10",
				Others:      "Other",
			},
			wantErr: false,
		},
		{
			name:  "com.example.app Android Device Model Other",
			input: "com.example.app/4.1.0 (Device Model, Android 9, Other)",
			want: &ua.UserAgent{
				AppName:     "com.example.app",
				AppVersion:  "4.1.0",
				DeviceModel: "Device Model",
				OSName:      "Android",
				OSVersion:   "9",
				Others:      "Other",
			},
			wantErr: false,
		},
		{
			name:  "App Windows Device Model 1 Other",
			input: "App/3.2.1 (Device Model 1, Windows 10, Other)",
			want: &ua.UserAgent{
				AppName:     "App",
				AppVersion:  "3.2.1",
				DeviceModel: "Device Model 1",
				OSName:      "Windows",
				OSVersion:   "10",
				Others:      "Other",
			},
			wantErr: false,
		},
		{
			name:  "App 1 Android Device Model Other",
			input: "App 1/1.1.1 (Device Model, Android 8, Other)",
			want: &ua.UserAgent{
				AppName:     "App 1",
				AppVersion:  "1.1.1",
				DeviceModel: "Device Model",
				OSName:      "Android",
				OSVersion:   "8",
				Others:      "Other",
			},
			wantErr: false,
		},
		{
			name:  "App Android Device Model Other",
			input: "App/2.2.2 (Device Model, Android 11, Other)",
			want: &ua.UserAgent{
				AppName:     "App",
				AppVersion:  "2.2.2",
				DeviceModel: "Device Model",
				OSName:      "Android",
				OSVersion:   "11",
				Others:      "Other",
			},
			wantErr: false,
		},
		{
			name:  "App-1 Windows Device-Model Other",
			input: "App-1/10 (Device-Model, Windows 10, Other)",
			want: &ua.UserAgent{
				AppName:     "App-1",
				AppVersion:  "10",
				DeviceModel: "Device-Model",
				OSName:      "Windows",
				OSVersion:   "10",
				Others:      "Other",
			},
			wantErr: false,
		},
		{
			name:  "App-1 Windows with OS detail, Device-Model Other",
			input: "App-1/0.1.1 (Device-Model, Windows 10 Pro 10.0 (Build 19045), Other)",
			want: &ua.UserAgent{
				AppName:     "App-1",
				AppVersion:  "0.1.1",
				DeviceModel: "Device-Model",
				OSName:      "Windows",
				OSVersion:   "10",
				Others:      "Other",
			},
		},
		{
			name:  "minimum user agent",
			input: "a/1 (b, c 1, d)",
			want: &ua.UserAgent{
				AppName:     "a",
				AppVersion:  "1",
				DeviceModel: "b",
				OSName:      "c",
				OSVersion:   "1",
				Others:      "d",
			},
			wantErr: false,
		},
		{
			name:    "empty user agent",
			input:   "",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "app name includes a slash",
			input:   "App//1.0.0 (Device, OS 1, Other)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "app version is not a number",
			input:   "App/notNum (Device, OS 1, Other)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "too many app version groups",
			input:   "App/1.0.0.0 (Device, OS 1, Other)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Device Model includes a comma",
			input:   "App/1.0.0 (Device,, OS 1, Other)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "no space between os name and version",
			input:   "App/1.0.0 (Device, OS1, Other)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "os version is not a number",
			input:   "App/1.0.0 (Device, OS hoge, Other)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "retailer includes a closing parenthesis",
			input:   "App/1.0.0 (Device, OS 1, Other))",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ua.Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, but got: %v", tt.wantErr, err)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("unexpected result (-want +got):\n%s", diff)
			}
		})
	}
}
