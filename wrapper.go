package mediainfo

/*
#cgo linux LDFLAGS: -ldl
#cgo darwin LDFLAGS: -framework CoreFoundation
#include <stdlib.h>
#include "clib/mediainfo.c"
*/
import "C"

import (
	"errors"
	"log"
	"runtime"
	"strconv"
	"time"
	"unsafe"
)

// wrapper is an instance of a c_mediainfo accessor.
type wrapper struct {
	cptr   unsafe.Pointer
}

// streamKind is used to specify the type of stream (audio, video, chapters, etc) when getting information.
type streamKind int

// infoKind is used to specify the aspect of information (name, value, unit of measure) when retrieving information.
type infoKind int

const (
	streamGeneral streamKind = 0
	streamVideo   streamKind = 1
	streamAudio   streamKind = 2
)

const (
	infoName     infoKind = 0
	infoText     infoKind = 1
	infoMeasure  infoKind = 2
	infoNameText infoKind = 4
	infoInfo     infoKind = 6
)

func toCInfo(i infoKind) C.MediaInfo_info_C {
	return C.MediaInfo_info_C(i)
}

func toCStream(s streamKind) C.MediaInfo_stream_C {
	return C.MediaInfo_stream_C(s)
}

var errOpenFailed = errors.New("file open failed")

func init() {
	C.MediaInfoDLL_Load()
}

func newWrapper() *wrapper {
	cmi := C.g_MediaInfo_New()
	w := &wrapper{cmi}
	runtime.SetFinalizer(w, func(w *wrapper) {
		if w.cptr != nil {
			C.g_MediaInfo_Delete(w.cptr)
		}
	})
	return w
}

func (w *wrapper) open(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	rc := C.g_MediaInfo_Open(w.cptr, cpath)

	if rc != 1 {
		return errOpenFailed
	}
	return nil
}

func (w *wrapper) get(section streamKind, param string, streamNum int) string {
	return w.getKind(section, streamNum, param, infoText)
}

func (w *wrapper) getKind(streamKind streamKind, streamNumber int, parameter string, kindOfInfo infoKind) string {
	cparameter := C.CString(parameter)
	val := C.g_MediaInfo_Get(w.cptr, toCStream(streamKind), C.size_t(streamNumber), cparameter, toCInfo(kindOfInfo), toCInfo(infoName))
	C.free(unsafe.Pointer(cparameter))

	return C.GoString(val)
}

func (w *wrapper) count(streamKind streamKind) int {
	return int(C.g_MediaInfo_Count_Get(w.cptr, toCStream(streamKind)))
}

func (w *wrapper) close() {
	C.g_MediaInfo_Close(w.cptr)
}

func (w *wrapper) intParam(section streamKind, param string, streamNum int, logger *log.Logger) IntValue {
	i, err := strconv.Atoi(w.get(section, param, streamNum))
	if err != nil {
		logger.Printf("converting param %q to an int: %v", param, err)
	}

	return IntValue{
		Val:   i,
		Extra: w.extraFor(section, param, streamNum),
	}
}

func (w *wrapper) int64Param(section streamKind, param string, streamNum int, logger *log.Logger) Int64Value {
	i, err := strconv.ParseInt(w.get(section, param, streamNum), 10, 64)
	if err != nil {
		logger.Printf("converting param %q to an int64: %v", param, err)
	}

	return Int64Value{
		Val:   i,
		Extra: w.extraFor(section, param, streamNum),
	}
}

func (w *wrapper) float64Param(section streamKind, param string, streamNum int, logger *log.Logger) Float64Value {
	f, err := strconv.ParseFloat(w.get(section, param, streamNum), 64)
	if err != nil {
		logger.Printf("converting param %q to an float64: %v", param, err)
	}

	return Float64Value{
		Val:   f,
		Extra: w.extraFor(section, param, streamNum),
	}
}

func (w *wrapper) stringParam(section streamKind, param string, streamNum int) StringValue {
	return StringValue{
		Val:   w.get(section, param, streamNum),
		Extra: w.extraFor(section, param, streamNum),
	}
}

func (w *wrapper) boolParam(section streamKind, param string, streamNum int) BoolValue {
	return BoolValue{
		Val:   w.get(section, param, streamNum) != "No",
		Extra: w.extraFor(section, param, streamNum),
	}
}

func (w *wrapper) timeParam(section streamKind, param string, streamNum int, logger *log.Logger) TimeValue {
	time, err := time.Parse("MST 2006-01-02 15:04:05", w.get(section, param, streamNum))
	if err != nil {
		logger.Printf("converting param %q to time.Time: %v", param, err)
	}

	return TimeValue{
		Val:   time,
		Extra: w.extraFor(section, param, streamNum),
	}
}

func (w *wrapper) localTimeParam(section streamKind, param string, streamNum int, logger *log.Logger) TimeValue {
	time, err := time.Parse("2006-01-02 15:04:05", w.get(section, param, streamNum))
	if err != nil {
		logger.Printf("converting param %q to time.Time: %v", param, err)
	}

	return TimeValue{
		Val:   time,
		Extra: w.extraFor(section, param, streamNum),
	}
}

func (w *wrapper) extraFor(section streamKind, param string, streamNum int) Extra {
	return Extra{
		Measure:  w.getKind(section, streamNum, param, infoMeasure),
		NameText: w.getKind(section, streamNum, param, infoNameText),
		Info:     w.getKind(section, streamNum, param, infoInfo),
	}
}
