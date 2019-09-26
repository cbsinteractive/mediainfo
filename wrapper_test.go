package mediainfo

import (
	"fmt"
)

func ExampleCount() {
	mi := newWrapper()
	err := mi.open("./testdata/test_bbb_360x240_1mb.mp4")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mi.count(streamVideo), mi.count(streamAudio))
	// Output: 1 1
}

func ExampleGet() {
	mi := newWrapper()
	err := mi.open("./testdata/test_bbb_360x240_1mb.mp4")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mi.get(streamGeneral, genParamFormat, 0))
	// Output: MPEG-4
}
