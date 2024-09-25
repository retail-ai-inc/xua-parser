package xua_parser_test

import (
	"fmt"

	ua "github.com/retail-ai-inc/xua-parser"
)

func ExampleUserAgent() {

	userAgent := `jp.retailai.App/3.9.3 (Devicel-Model, Android 10, Other)`
	result, err := ua.Parse(userAgent)
	if err != nil {
		panic(err)
	}

	// Output:
	// ua.UserAgent{
	// 	AppName:     "jp.retailai.App",
	// 	AppVersion:  "3.9.3",
	// 	DeviceModel: "Devicel-Model",
	// 	OSName:      "Android",
	// 	OSVersion:   "10",
	// 	Others:      "Other",
	// }
	fmt.Printf("ua.UserAgent{\n")
	fmt.Printf("\tAppName:     %q,\n", result.AppName)
	fmt.Printf("\tAppVersion:  %q,\n", result.AppVersion)
	fmt.Printf("\tDeviceModel: %q,\n", result.DeviceModel)
	fmt.Printf("\tOSName:      %q,\n", result.OSName)
	fmt.Printf("\tOSVersion:   %q,\n", result.OSVersion)
	fmt.Printf("\tOthers:      %q,\n", result.Others)
	fmt.Printf("}\n")

}
