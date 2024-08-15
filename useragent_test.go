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
			name: "cart android trial",
			ua:   "jp.retailai.raicart/3.9.3 (S-500, Android 10, trial)",
			want: &useragent.UserAgent{
				AppName:     "jp.retailai.raicart",
				AppVersion:  "3.9.3",
				DeviceModel: "S-500",
				OSName:      "Android",
				OSVersion:   "10",
				Retailer:    "trial",
			},
			wantErr: false,
		},
		{
			name: "sct windows aeon",
			ua:   "SCT/1.0.0 (Toshiba T380, Windows 11, Aeon)",
			want: &useragent.UserAgent{
				AppName:     "SCT",
				AppVersion:  "1.0.0",
				DeviceModel: "Toshiba T380",
				OSName:      "Windows",
				OSVersion:   "11",
				Retailer:    "Aeon",
			},
			wantErr: false,
		},
		{
			name: "ism flutter sugi",
			ua:   "ISM Flutter/2.3.1 (Samsung S8 Pro, Android 11, Sugi)",
			want: &useragent.UserAgent{
				AppName:     "ISM Flutter",
				AppVersion:  "2.3.1",
				DeviceModel: "Samsung S8 Pro",
				OSName:      "Android",
				OSVersion:   "11",
				Retailer:    "Sugi",
			},
			wantErr: false,
		},
		{
			name: "cart windows trial",
			ua:   "SC Windows/4.0.0 (HP EliteBook, Windows 10, TRIAL)",
			want: &useragent.UserAgent{
				AppName:     "SC Windows",
				AppVersion:  "4.0.0",
				DeviceModel: "HP EliteBook",
				OSName:      "Windows",
				OSVersion:   "10",
				Retailer:    "TRIAL",
			},
			wantErr: false,
		},
		{
			name: "cart android taiyo",
			ua:   "jp.retailai.raicart/4.1.0 (Huawei P30, Android 9, Taiyo)",
			want: &useragent.UserAgent{
				AppName:     "jp.retailai.raicart",
				AppVersion:  "4.1.0",
				DeviceModel: "Huawei P30",
				OSName:      "Android",
				OSVersion:   "9",
				Retailer:    "Taiyo",
			},
			wantErr: false,
		},
		{
			name: "ism windows aeon",
			ua:   "ISM/3.2.1 (Dell XPS 13, Windows 10, Aeon)",
			want: &useragent.UserAgent{
				AppName:     "ISM",
				AppVersion:  "3.2.1",
				DeviceModel: "Dell XPS 13",
				OSName:      "Windows",
				OSVersion:   "10",
				Retailer:    "Aeon",
			},
			wantErr: false,
		},
		{
			name: "sct flutter trial",
			ua:   "SCT Flutter/1.1.1 (Sony Xperia, Android 8, TRIAL)",
			want: &useragent.UserAgent{
				AppName:     "SCT Flutter",
				AppVersion:  "1.1.1",
				DeviceModel: "Sony Xperia",
				OSName:      "Android",
				OSVersion:   "8",
				Retailer:    "TRIAL",
			},
			wantErr: false,
		},
		{
			name: "cart windows taiyo",
			ua:   "SC Windows/2.5.0 (Lenovo ThinkPad, Windows 8, Taiyo)",
			want: &useragent.UserAgent{
				AppName:     "SC Windows",
				AppVersion:  "2.5.0",
				DeviceModel: "Lenovo ThinkPad",
				OSName:      "Windows",
				OSVersion:   "8",
				Retailer:    "Taiyo",
			},
			wantErr: false,
		},
		{
			name: "ism flutter sugi",
			ua:   "ISM Flutter/3.0.0 (Samsung Galaxy S10, Android 10, Sugi)",
			want: &useragent.UserAgent{
				AppName:     "ISM Flutter",
				AppVersion:  "3.0.0",
				DeviceModel: "Samsung Galaxy S10",
				OSName:      "Android",
				OSVersion:   "10",
				Retailer:    "Sugi",
			},
			wantErr: false,
		},
		{
			name: "sct android trial",
			ua:   "SCT/2.2.2 (Google Pixel, Android 11, TRIAL)",
			want: &useragent.UserAgent{
				AppName:     "SCT",
				AppVersion:  "2.2.2",
				DeviceModel: "Google Pixel",
				OSName:      "Android",
				OSVersion:   "11",
				Retailer:    "TRIAL",
			},
			wantErr: false,
		},
		{
			name: "ism flutter trial",
			ua:   "ism-flutter/1.3.0 (ESYP-10142789, Windows 10, trial)",
			want: &useragent.UserAgent{
				AppName:     "ism-flutter",
				AppVersion:  "1.3.0",
				DeviceModel: "ESYP-10142789",
				OSName:      "Windows",
				OSVersion:   "10",
				Retailer:    "trial",
			},
			wantErr: false,
		},
		{
			name: "ism windows trial",
			ua:   "ISM-Vue/10 (ESYP-10142789, Windows 10, trial)",
			want: &useragent.UserAgent{
				AppName:     "ISM-Vue",
				AppVersion:  "10",
				DeviceModel: "ESYP-10142789",
				OSName:      "Windows",
				OSVersion:   "10",
				Retailer:    "trial",
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
				Retailer:    "d",
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
			ua:      "App//1.0.0 (Device, OS 1, Retailer)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "app version is not a number",
			ua:      "App/notNum (Device, OS 1, Retailer)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "too many app version groups",
			ua:      "App/1.0.0.0 (Device, OS 1, Retailer)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "device model includes a comma",
			ua:      "App/1.0.0 (Device,, OS 1, Retailer)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "os name is not a letter",
			ua:      "App/1.0.0 (Device, 1 1, Retailer)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "no space between os name and version",
			ua:      "App/1.0.0 (Device, OS1, Retailer)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "os version is not a number",
			ua:      "App/1.0.0 (Device, OS hoge, Retailer)",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "retailer includes a closing parenthesis",
			ua:      "App/1.0.0 (Device, OS 1, Retailer))",
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
