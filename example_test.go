package useragent_test

import (
	"fmt"

	"github.com/retail-ai-inc/useragent"
)

func ExampleUserAgent() {

	userAgent := `jp.retailai.App/3.9.3 (Devicel-Model, Android 10, Other)`
	ua, err := useragent.Parse(userAgent)
	if err != nil {
		panic(err)
	}

	// Output:
	// useragent.UserAgent{
	// 	AppName:     "jp.retailai.App",
	// 	AppVersion:  "3.9.3",
	// 	DeviceModel: "Devicel-Model",
	// 	OSName:      "Android",
	// 	OSVersion:   "10",
	// 	Others:      "Other",
	// }
	fmt.Printf("useragent.UserAgent{\n")
	fmt.Printf("\tAppName:     %q,\n", ua.AppName)
	fmt.Printf("\tAppVersion:  %q,\n", ua.AppVersion)
	fmt.Printf("\tDeviceModel: %q,\n", ua.DeviceModel)
	fmt.Printf("\tOSName:      %q,\n", ua.OSName)
	fmt.Printf("\tOSVersion:   %q,\n", ua.OSVersion)
	fmt.Printf("\tOthers:      %q,\n", ua.Others)
	fmt.Printf("}\n")

}
