package mediainfo_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/cbsinteractive/mediainfo"
)

func ExampleNew() {
	mi, _ := mediainfo.New("./testdata/test_bbb_360x240_1mb.mp4")

	fmt.Println(mi.General.FileSize)
	// Output: {1053651 { byte File size File size in bytes}}
}

func ExampleNewWithLogger() {
	mi, _ := mediainfo.NewWithLogger(
		"./testdata/test_bbb_360x240_1mb.mp4",
		log.New(ioutil.Discard, "", 0),
	)

	fmt.Println(mi.General.FileSize)
	fmt.Println(mi.File)
	// Output:
	// {1053651 { byte File size File size in bytes}}
	// test_bbb_360x240_1mb.mp4
}

func TestNew_errors(t *testing.T) {
	_, err := mediainfo.New("./testdata/non_existent_file.mp4")
	if err == nil {
		t.Error("expected New to return an error if the file location is invalid")
	}
}
