package mediainfo

import "time"

// MediaInfo is the root container for all media metadata
type MediaInfo struct {
	File        string
	TmpFile     string `json:"-"`
	General     GeneralInfo
	VideoTracks []VideoTrack
	AudioTracks []AudioTrack
}

// GeneralInfo contains all stream metadata tagged as general
type GeneralInfo struct {
	VideoTrackCount       IntValue
	AudioTrackCount       IntValue
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

// VideoTrack contains all stream metadata for a single video track
type VideoTrack struct {
	StreamOrder           IntValue
	ID                    IntValue
	Format                StringValue
	Profile               StringValue
	FormatLevel           StringValue
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

// AudioTrack contains all stream metadata for a single video track
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

// Extra is the set of values that enrich a raw property value
type Extra struct {
	Measure  string
	NameText string
	Info     string
}

// IntValue is a custom wrapper for an int + Extra metadata
type IntValue struct {
	Val   int
	Extra Extra
}

// Int64Value is a custom wrapper for an int64 + Extra metadata
type Int64Value struct {
	Val   int64
	Extra Extra
}

// Float64Value is a custom wrapper for an float64 + Extra metadata
type Float64Value struct {
	Val   float64
	Extra Extra
}

// StringValue is a custom wrapper for an string + Extra metadata
type StringValue struct {
	Val   string
	Extra Extra
}

// BoolValue is a custom wrapper for a bool + Extra metadata
type BoolValue struct {
	Val   bool
	Extra Extra
}

// TimeValue is a custom wrapper for a time.Time + Extra metadata
type TimeValue struct {
	Val   time.Time
	Extra Extra
}
