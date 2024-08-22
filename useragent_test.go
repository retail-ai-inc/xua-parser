package useragent_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/retail-ai-inc/useragent"
)

func Test_Parse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		ua      string
		want    *useragent.UserAgent
		wantErr bool
	}{
		{
			name: "jp.retailai.App Android Device-Model Other",
			ua:   "jp.retailai.App/3.9.3 (Device-Model, Android 10, Other)",
			want: &useragent.UserAgent{
				AppName:     "jp.retailai.App",
				AppVersion:  "3.9.3",
				DeviceModel: "Device-Model",
				OSName:      "Android",
				OSVersion:   "10",
				Others:      "Other",
			},
			wantErr: false,
		},
		{
			name: "App 1 Android Device Model 1 Other",
			ua:   "App 1/2.3.1 (Device Model 1, Android 11, Other)",
			want: &useragent.UserAgent{
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
			name: "App Windows Windows Device Model Other",
			ua:   "App Windows/4.0.0 (Device Model, Windows 10, Other)",
			want: &useragent.UserAgent{
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
			name: "jp.retailai.App Android Device Model Other",
			ua:   "jp.retailai.App/4.1.0 (Device Model, Android 9, Other)",
			want: &useragent.UserAgent{
				AppName:     "jp.retailai.App",
				AppVersion:  "4.1.0",
				DeviceModel: "Device Model",
				OSName:      "Android",
				OSVersion:   "9",
				Others:      "Other",
			},
			wantErr: false,
		},
		{
			name: "App Windows Device Model 1 Other",
			ua:   "App/3.2.1 (Device Model 1, Windows 10, Other)",
			want: &useragent.UserAgent{
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
			name: "App 1 Android Device Model Other",
			ua:   "App 1/1.1.1 (Device Model, Android 8, Other)",
			want: &useragent.UserAgent{
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
			name: "App Android Device Model Other",
			ua:   "App/2.2.2 (Device Model, Android 11, Other)",
			want: &useragent.UserAgent{
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
			name: "App-1 Windows Device-Model Other",
			ua:   "App-1/10 (Device-Model, Windows 10, Other)",
			want: &useragent.UserAgent{
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
			name: "minimum user agent",
			ua:   "a/1 (b, c 1, d)",
			want: &useragent.UserAgent{
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
			ua:      "",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "app name includes a slash",
			ua:      "App//1.0.0 (Device, OS 1, Other)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "app version is not a number",
			ua:      "App/notNum (Device, OS 1, Other)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "too many app version groups",
			ua:      "App/1.0.0.0 (Device, OS 1, Other)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Device Model includes a comma",
			ua:      "App/1.0.0 (Device,, OS 1, Other)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "os name is not a letter",
			ua:      "App/1.0.0 (Device, 1 1, Other)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "no space between os name and version",
			ua:      "App/1.0.0 (Device, OS1, Other)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "os version is not a number",
			ua:      "App/1.0.0 (Device, OS hoge, Other)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "retailer includes a closing parenthesis",
			ua:      "App/1.0.0 (Device, OS 1, Other))",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := useragent.Parse(tt.ua)
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
