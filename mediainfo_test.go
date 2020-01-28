package mediainfo_test

import (
	"fmt"

	"github.com/cbsinteractive/mediainfo"
)

func ExampleAnalyze() {
	mi, _ := mediainfo.Analyze(
		"./testdata/test_bbb_360x240_1mb.mp4",
	)

	fmt.Println(mi.General.FileSize)
	fmt.Println(mi.File)
	// Output:
	// {1053651 {  }}
	// test_bbb_360x240_1mb.mp4
}
