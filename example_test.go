package useragent_test

import (
	"fmt"

	"github.com/retail-ai-inc/useragent"
)

func ExampleUserAgent() {

	userAgent := `jp.retailai.raicart/3.9.3 (S-500, Android 10, trial)`
	ua, err := useragent.Parse(userAgent)
	if err != nil {
		panic(err)
	}

	// Output:
	// useragent.UserAgent{
	// 	AppName:     "jp.retailai.raicart",
	// 	AppVersion:  "3.9.3",
	// 	DeviceModel: "S-500",
	// 	OSName:      "Android",
	// 	OSVersion:   "10",
	// 	Retailer:    "trial",
	// }
	fmt.Printf("useragent.UserAgent{\n")
	fmt.Printf("\tAppName:     %q,\n", ua.AppName)
	fmt.Printf("\tAppVersion:  %q,\n", ua.AppVersion)
	fmt.Printf("\tDeviceModel: %q,\n", ua.DeviceModel)
	fmt.Printf("\tOSName:      %q,\n", ua.OSName)
	fmt.Printf("\tOSVersion:   %q,\n", ua.OSVersion)
	fmt.Printf("\tRetailer:    %q,\n", ua.Retailer)
	fmt.Printf("}\n")

}
