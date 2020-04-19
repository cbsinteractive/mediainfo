package mediainfo

import "fmt"

func ExampleAnalyze() {
	mi, _ := Analyze(
		"./testdata/test_bbb_360x240_1mb.mp4",
	)

	fmt.Println(mi.General.FileSize)
	fmt.Println(mi.File)
	// Output:
	// {1053651 {  }}
	// test_bbb_360x240_1mb.mp4
}
