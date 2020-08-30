package mediainfo

import (
	"strconv"
	"time"
)

const (
	typeGeneral = "General"
	typeVideo   = "Video"
	typeAudio   = "Audio"
	typeOther   = "Other"

	typeSecondaryTimecode = "Time code"
)

type cliMediaInfo struct {
	Media Media `json:"media"`
}

type Track struct {
	Type                     string `json:"@type"`
	TypeSecondary            string `json:"Type"`
	VideoCount               string `json:"VideoCount,omitempty"`
	AudioCount               string `json:"AudioCount,omitempty"`
	TextCount                string `json:"TextCount,omitempty"`
	FileExtension            string `json:"FileExtension,omitempty"`
	Format                   string `json:"Format"`
	FormatProfile            string `json:"Format_Profile,omitempty"`
	CodecID                  string `json:"CodecID,omitempty"`
	CodecIDCompatible        string `json:"CodecID_Compatible,omitempty"`
	FileSize                 string `json:"FileSize,omitempty"`
	Duration                 string `json:"Duration"`
	OverallBitRateMode       string `json:"OverallBitRate_Mode,omitempty"`
	OverallBitRate           string `json:"OverallBitRate,omitempty"`
	FrameRate                string `json:"FrameRate,omitempty"`
	FrameCount               string `json:"FrameCount,omitempty"`
	StreamSize               string `json:"StreamSize"`
	HeaderSize               string `json:"HeaderSize,omitempty"`
	DataSize                 string `json:"DataSize,omitempty"`
	FooterSize               string `json:"FooterSize,omitempty"`
	IsStreamable             string `json:"IsStreamable,omitempty"`
	EncodedDate              string `json:"Encoded_Date,omitempty"`
	TaggedDate               string `json:"Tagged_Date,omitempty"`
	LawRating                string `json:"LawRating,omitempty"`
	FileModifiedDate         string `json:"File_Modified_Date,omitempty"`
	FileModifiedDateLocal    string `json:"File_Modified_Date_Local,omitempty"`
	EncodedApplication       string `json:"Encoded_Application,omitempty"`
	StreamOrder              string `json:"StreamOrder,omitempty"`
	ID                       string `json:"ID,omitempty"`
	FormatLevel              string `json:"Format_Level,omitempty"`
	FormatSettingsCABAC      string `json:"Format_Settings_CABAC,omitempty"`
	FormatSettingsRefFrames  string `json:"Format_Settings_RefFrames,omitempty"`
	FormatSettingsGOP        string `json:"Format_Settings_GOP,omitempty"`
	BitRateMode              string `json:"BitRate_Mode,omitempty"`
	BitRate                  string `json:"BitRate,omitempty"`
	Width                    string `json:"Width,omitempty"`
	Height                   string `json:"Height,omitempty"`
	HeightOffset             string `json:"Height_Offset,omitempty"`
	StoredHeight             string `json:"Stored_Height,omitempty"`
	SampledWidth             string `json:"Sampled_Width,omitempty"`
	SampledHeight            string `json:"Sampled_Height,omitempty"`
	PixelAspectRatio         string `json:"PixelAspectRatio,omitempty"`
	DisplayAspectRatio       string `json:"DisplayAspectRatio,omitempty"`
	Rotation                 string `json:"Rotation,omitempty"`
	FrameRateMode            string `json:"FrameRate_Mode,omitempty"`
	FrameRateMinimum         string `json:"FrameRate_Minimum,omitempty"`
	FrameRateMaximum         string `json:"FrameRate_Maximum,omitempty"`
	FrameRateOriginal        string `json:"FrameRate_Original,omitempty"`
	ColorSpace               string `json:"ColorSpace,omitempty"`
	ChromaSubsampling        string `json:"ChromaSubsampling,omitempty"`
	BitDepth                 string `json:"BitDepth,omitempty"`
	ScanType                 string `json:"ScanType,omitempty"`
	BufferSize               string `json:"BufferSize,omitempty"`
	FormatAdditionalFeatures string `json:"Format_AdditionalFeatures,omitempty"`
	Channels                 string `json:"Channels,omitempty"`
	ChannelPositions         string `json:"ChannelPositions,omitempty"`
	ChannelLayout            string `json:"ChannelLayout,omitempty"`
	SamplesPerFrame          string `json:"SamplesPerFrame,omitempty"`
	SamplingRate             string `json:"SamplingRate,omitempty"`
	SamplingCount            string `json:"SamplingCount,omitempty"`
	CompressionMode          string `json:"Compression_Mode,omitempty"`
	StreamSizeProportion     string `json:"StreamSize_Proportion,omitempty"`
	Default                  string `json:"Default,omitempty"`
	AlternateGroup           string `json:"AlternateGroup,omitempty"`
	Typeorder                string `json:"@typeorder,omitempty"`
	MuxingMode               string `json:"MuxingMode,omitempty"`
	MuxingModeMoreInfo       string `json:"MuxingMode_MoreInfo,omitempty"`
	StreamSizeEncoded        string `json:"StreamSize_Encoded,omitempty"`
	FirstFrameTimecode       string `json:"TimeCode_FirstFrame,omitempty"`
	TimecodeSettings         string `json:"TimeCode_Settings,omitempty"`
	Delay                    string `json:"Delay,omitempty"`
}

type Media struct {
	Ref   string  `json:"@ref"`
	Track []Track `json:"track"`
}

func (m cliMediaInfo) toMediaInfo() MediaInfo {
	mi := MediaInfo{}
	for _, track := range m.Media.Track {
		switch track.Type {
		case typeGeneral:
			mi.General = generalInfoFrom(track)
		case typeVideo:
			mi.VideoTracks = append(mi.VideoTracks, videoTrackFrom(track))
		case typeAudio:
			mi.AudioTracks = append(mi.AudioTracks, audioTrackFrom(track))
		case typeOther:
			switch track.TypeSecondary {
			case typeSecondaryTimecode:
				mi.TimecodeTracks = append(mi.TimecodeTracks, timecodeTracksFrom(track))
			}
		}
	}

	return mi
}

func audioTrackFrom(track Track) AudioTrack {
	return AudioTrack{
		StreamOrder:              intParam(track.StreamOrder),
		ID:                       intParam(track.ID),
		Format:                   stringParam(track.Format),
		FormatAdditionalFeatures: stringParam(track.FormatAdditionalFeatures),
		CodecID:                  stringParam(track.CodecID),
		Duration:                 float64Param(track.Duration),
		BitrateMode:              stringParam(track.BitRateMode),
		Bitrate:                  intParam(track.BitRate),
		Channels:                 intParam(track.Channels),
		ChannelPositions:         stringParam(track.ChannelPositions),
		ChannelLayout:            stringParam(track.ChannelLayout),
		SamplesPerFrame:          intParam(track.SamplesPerFrame),
		SamplingRate:             intParam(track.SamplingRate),
		SamplingCount:            intParam(track.SamplingCount),
		FrameRate:                float64Param(track.FrameRate),
		FrameCount:               intParam(track.FrameCount),
		CompressionMode:          stringParam(track.CompressionMode),
		StreamSize:               int64Param(track.StreamSize),
		StreamSizeProportion:     float64Param(track.StreamSizeProportion),
		IsDefault:                boolParam(track.Default),
		AlternateGroup:           stringParam(track.AlternateGroup),
		EncodedDate:              timeParam(track.EncodedDate),
		TaggedDate:               timeParam(track.TaggedDate),
	}

}

func videoTrackFrom(track Track) VideoTrack {
	format := stringParam(track.Format)
	profile := stringParam(track.FormatProfile)

	return VideoTrack{
		StreamOrder:        intParam(track.StreamOrder),
		ID:                 intParam(track.ID),
		Format:             stringParam(track.Format),
		Profile:            stringParam(track.FormatProfile),
		FormatLevel:        stringParam(track.FormatLevel),
		IsCABACEnabled:     boolParam(track.FormatSettingsCABAC),
		RefFrames:          intParam(track.FormatSettingsRefFrames),
		CodecID:            stringParam(track.CodecID),
		Duration:           float64Param(track.Duration),
		Bitrate:            intParam(track.BitRate),
		Width:              intParam(track.Width),
		Height:             intParam(track.Height),
		HeightOffset:       intParam(track.HeightOffset),
		SampledWidth:       intParam(track.SampledWidth),
		SampledHeight:      intParam(track.SampledHeight),
		PixelAspectRatio:   float64Param(track.PixelAspectRatio),
		DisplayAspectRatio: float64Param(track.DisplayAspectRatio),
		Rotation:           float64Param(track.Rotation),
		FrameRateMode:      stringParam(track.FrameRateMode),
		FrameRate:          float64Param(track.FrameRate),
		FrameCount:         intParam(track.FrameCount),
		ColorSpace:         stringParam(track.ColorSpace),
		ChromaSubsampling:  stringParam(track.ChromaSubsampling),
		BitDepth:           bitDepthFrom(track, format, profile),
		ScanType:           stringParam(track.ScanType),
		StreamSize:         int64Param(track.StreamSize),
		EncodedDate:        timeParam(track.EncodedDate),
		TaggedDate:         timeParam(track.TaggedDate),
		FirstFrameTimecode: stringParam(track.FirstFrameTimecode),
		Delay:              float64Param(track.Delay),
	}

}

func generalInfoFrom(track Track) GeneralInfo {
	return GeneralInfo{
		VideoTrackCount:       intParam(track.VideoCount),
		AudioTrackCount:       intParam(track.AudioCount),
		Format:                stringParam(track.Format),
		FormatProfile:         stringParam(track.FormatProfile),
		CodecID:               stringParam(track.CodecID),
		CodecIDCompatible:     stringParam(track.CodecIDCompatible),
		FileSize:              int64Param(track.FileSize),
		Duration:              float64Param(track.Duration),
		BitrateMode:           stringParam(track.BitRateMode),
		Bitrate:               intParam(track.BitRate),
		FrameRate:             float64Param(track.FrameRate),
		FrameCount:            intParam(track.FrameCount),
		StreamSize:            int64Param(track.StreamSize),
		HeaderSize:            int64Param(track.HeaderSize),
		DataSize:              int64Param(track.DataSize),
		FooterSize:            int64Param(track.FooterSize),
		IsStreamable:          boolParam(track.IsStreamable),
		EncodedDate:           timeParam(track.EncodedDate),
		TaggedDate:            timeParam(track.TaggedDate),
		FileModifiedDate:      timeParam(track.FileModifiedDate),
		FileModifiedDateLocal: localTimeParam(track.FileModifiedDateLocal),
		EncodedApplication:    stringParam(track.EncodedApplication),
	}
}

func timecodeTracksFrom(track Track) TimecodeTrack {
	return TimecodeTrack{
		ID:                 stringParam(track.ID),
		Format:             stringParam(track.Format),
		FirstFrameTimecode: stringParam(track.FirstFrameTimecode),
		Settings:           stringParam(track.TimecodeSettings),
		FrameRate:          float64Param(track.FrameRate),
		Delay:              float64Param(track.Delay),
	}
}

func intParam(val string) IntValue {
	if val == "" {
		return IntValue{Val: 0}
	}

	i, _ := strconv.Atoi(val)

	return IntValue{Val: i}
}

func int64Param(val string) Int64Value {
	if val == "" {
		return Int64Value{Val: 0}
	}

	i, _ := strconv.ParseInt(val, 10, 64)

	return Int64Value{Val: i}
}

func float64Param(val string) Float64Value {
	if val == "" {
		return Float64Value{Val: 0}
	}

	f, _ := strconv.ParseFloat(val, 64)

	return Float64Value{Val: f}
}

func stringParam(val string) StringValue {
	return StringValue{Val: val}
}

func boolParam(val string) BoolValue {
	return BoolValue{Val: val != "No"}
}

func timeParam(val string) TimeValue {
	t, _ := time.Parse("MST 2006-01-02 15:04:05", val)

	return TimeValue{Val: t}
}

func localTimeParam(val string) TimeValue {
	t, _ := time.Parse("2006-01-02 15:04:05", val)

	return TimeValue{Val: t}
}

func bitDepthFrom(track Track, format, profile StringValue) IntValue {
	if format.Val == videoFormatProRes && profile.Val == videoProfile4444 {
		return IntValue{Val: bitDepth12}
	} else if format.Val == videoFormatProRes && profile.Val == videoProfile422HQ {
		return IntValue{Val: bitDepth10}
	}

	return intParam(track.BitDepth)
}
