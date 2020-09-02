# mediainfo
A wrapper for [libmediainfo](https://mediaarea.net/en/MediaInfo) -- returns metadata about media streams

### Dependencies

- mediainfo CLI

On macOS (via [homebrew](https://brew.sh/)):

    brew install mediainfo

Or install mediainfo from [here](https://mediaarea.net/en/MediaInfo/Download).

Source and other platform downloads can be found [here](https://mediaarea.net/en/MediaInfo/Download).

---

### Usage

```go
info, err := mediainfo.Analyze("http://domain.com/file.mp4")
if err != nil {
  // handle err
}

fmt.Println(info) // access and use metadata from this filled info struct
```

Another easy way to get started is to run the examples in `mediainfo_test.go`

### Tests

Run:

    go test ./... -race -v -cover -covermode=atomic

### Documentation

`import "github.com/cbsinteractive/mediabench/pkg/mediainfo"`

## <a name="pkg-index">Index</a>
* [type MediaInfo](#MediaInfo)
  * [func New(url string) (*MediaInfo, error)](#New)
* [type GeneralInfo](#GeneralInfo)
* [type VideoTrack](#VideoTrack)
* [type AudioTrack](#AudioTrack)
* [type TextTrack](#TextTrack)
* [type TimecodeTrack](#TimecodeTrack)
* [type BoolValue](#BoolValue)
* [type StringValue](#StringValue)
* [type IntValue](#IntValue)
* [type Int64Value](#Int64Value)
* [type Float64Value](#Float64Value)
* [type TimeValue](#TimeValue)
* [type Extra](#Extra)

#### <a name="pkg-files">Package files</a>
[mediainfo.go](mediainfo.go) [models.go](models.go) [params.go](params.go) [wrapper.go](wrapper.go) 

## <a name="MediaInfo">type</a> [MediaInfo](models.go?s=92:194#L6)
``` go
type MediaInfo struct {
    General     GeneralInfo
    VideoTracks []VideoTrack
    AudioTracks []AudioTrack
}

```
MediaInfo is the root container for all media metadata

### <a name="New">func</a> [New](mediainfo.go?s=115:175#L8)
``` go
func New(url string) (*MediaInfo, error)
```
New creates and returns MediaInfo from a url, optionally returns an error

## <a name="GeneralInfo">type</a> [GeneralInfo](models.go?s=258:1065#L17)
``` go
type GeneralInfo struct {
    VideoTrackCount       IntValue
    AudioTrackCount       IntValue
    TextTrackCount        IntValue
    FileExtension         StringValue
    Format                StringValue
    FormatProfile         StringValue
    CodecID               StringValue
    CodecIDCompatible     StringValue
    FileSize              Int64Value
    Duration              Float64Value
    BitrateMode           StringValue
    Bitrate               IntValue
    FrameRate             Float64Value
    FrameCount            IntValue
    StreamSize            Int64Value
    HeaderSize            Int64Value
    DataSize              Int64Value
    FooterSize            Int64Value
    IsStreamable          BoolValue
    EncodedDate           TimeValue
    TaggedDate            TimeValue
    FileModifiedDate      TimeValue
    FileModifiedDateLocal TimeValue
    EncodedApplication    StringValue
}

```
GeneralInfo contains all stream metadata tagged as general

## <a name="VideoTrack">type</a> [VideoTrack](models.go?s=1135:2109#L44)
``` go
type VideoTrack struct {
    StreamOrder           IntValue
    ID                    IntValue
    Format                StringValue
    Profile               StringValue
    Level                 StringValue
    IsCABACEnabled        BoolValue
    RefFrames             IntValue
    CodecID               StringValue
    Duration              Float64Value
    Bitrate               IntValue
    Width                 IntValue
    Height                IntValue
    SampledWidth          IntValue
    SampledHeight         IntValue
    PixelAspectRatio      Float64Value
    DisplayAspectRatio    Float64Value
    Rotation              Float64Value
    FrameRateMode         StringValue
    FrameRateModeOriginal StringValue
    FrameRate             Float64Value
    FrameCount            IntValue
    ColorSpace            StringValue
    ChromaSubsampling     StringValue
    BitDepth              IntValue
    ScanType              StringValue
    StreamSize            Int64Value
    EncodedDate           TimeValue
    TaggedDate            TimeValue
}

```
VideoTrack contains all stream metadata for a single video track

## <a name="AudioTrack">type</a> [AudioTrack](models.go?s=2179:3086#L79)
``` go
type AudioTrack struct {
    StreamOrder              IntValue
    ID                       IntValue
    Format                   StringValue
    FormatAdditionalFeatures StringValue
    CodecID                  StringValue
    Duration                 Float64Value
    BitrateMode              StringValue
    Bitrate                  IntValue
    BitrateMaximum           IntValue
    Channels                 IntValue
    ChannelPositions         StringValue
    ChannelLayout            StringValue
    SamplesPerFrame          IntValue
    SamplingRate             IntValue
    SamplingCount            IntValue
    FrameRate                Float64Value
    FrameCount               IntValue
    CompressionMode          StringValue
    StreamSize               Int64Value
    StreamSizeProportion     Float64Value
    IsDefault                BoolValue
    AlternateGroup           StringValue
    EncodedDate              TimeValue
    TaggedDate               TimeValue
}

```
AudioTrack contains all stream metadata for a single video track

## <a name="TextTrack">type</a> [TextTrack](models.go?s=2179:3086#L107)
```go
type TextTrack struct {
	ID              IntValue
	Format          StringValue
	Duration        Float64Value
	Width           IntValue
	Height          IntValue
	FrameRate       Float64Value
	CompressionMode StringValue
	Language        StringValue
	Delay           Float64Value
}
```
TextTrack contains all stream metadata for a single text track

## <a name="TimecodeTrack">type</a> [TimecodeTrack](models.go?s=2179:3086#L119)
```go
type TimecodeTrack struct {
	ID, Format, FirstFrameTimecode, Settings StringValue
	FrameRate, Delay                         Float64Value
}
```
TimecodeTrack contains all stream metadata for a single timecode track

## <a name="BoolValue">type</a> [BoolValue](models.go?s=3762:3812#L156)
``` go
type BoolValue struct {
    Val   bool
    Extra Extra
}

```
BoolValue is a custom wrapper for a bool + Extra metadata

## <a name="StringValue">type</a> [StringValue](models.go?s=3645:3699#L150)
``` go
type StringValue struct {
    Val   string
    Extra Extra
}

```
StringValue is a custom wrapper for an string + Extra metadata

## <a name="IntValue">type</a> [IntValue](models.go?s=3285:3333#L132)
``` go
type IntValue struct {
    Val   int
    Extra Extra
}

```
IntValue is a custom wrapper for an int + Extra metadata

## <a name="Int64Value">type</a> [Int64Value](models.go?s=3399:3451#L138)
``` go
type Int64Value struct {
    Val   int64
    Extra Extra
}

```
Int64Value is a custom wrapper for an int64 + Extra metadata

## <a name="Float64Value">type</a> [Float64Value](models.go?s=3521:3577#L144)
type Float64Value struct {
    Val   float64
    Extra Extra
}

```
Float64Value is a custom wrapper for an float64 + Extra metadata

## <a name="TimeValue">type</a> [TimeValue](models.go?s=3880:3935#L162)
``` go
type TimeValue struct {
    Val   time.Time
    Extra Extra
}

```
TimeValue is a custom wrapper for a time.Time + Extra metadata

## <a name="Extra">type</a> [Extra](models.go?s=3151:3223#L125)
``` go
type Extra struct {
    Measure  string
    NameText string
    Info     string
}

```
Extra is the set of values that enrich a raw property value

