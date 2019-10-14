package mediainfo

import (
	"io/ioutil"
	"log"
	"net/url"
	"path"
)

// New creates and returns MediaInfo from a url, optionally returns an error
func New(urlStr string) (*MediaInfo, error) {
	silentLogger := log.New(ioutil.Discard, "", 0)
	return NewWithLogger(urlStr, silentLogger)
}

// NewWithLogger creates and returns MediaInfo from a url and logger, optionally returns an error
func NewWithLogger(urlStr string, logger *log.Logger) (*MediaInfo, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	return newWithLoggerAndFilename(urlStr, logger, path.Base(u.Path))
}

func newWithLoggerAndFilename(url string, logger *log.Logger, filename string) (*MediaInfo, error) {
	w := newWrapper()
	defer w.close()

	err := w.open(url)
	if err != nil {
		return nil, err
	}

	generalInfo := GeneralInfo{
		VideoTrackCount:       w.intParam(streamGeneral, genParamVideoCount, 0, logger),
		AudioTrackCount:       w.intParam(streamGeneral, genParamAudioCount, 0, logger),
		Format:                w.stringParam(streamGeneral, genParamFormat, 0),
		FormatProfile:         w.stringParam(streamGeneral, genParamFormatProfile, 0),
		CodecID:               w.stringParam(streamGeneral, genParamCodecID, 0),
		CodecIDCompatible:     w.stringParam(streamGeneral, genParamCodecIDCompatible, 0),
		FileSize:              w.int64Param(streamGeneral, genParamFileSize, 0, logger),
		Duration:              w.float64Param(streamGeneral, genParamDuration, 0, logger),
		BitrateMode:           w.stringParam(streamGeneral, genParamOverallBitRateMode, 0),
		Bitrate:               w.intParam(streamGeneral, genParamOverallBitRate, 0, logger),
		FrameRate:             w.float64Param(streamGeneral, genParamFrameRate, 0, logger),
		FrameCount:            w.intParam(streamGeneral, genParamFrameCount, 0, logger),
		StreamSize:            w.int64Param(streamGeneral, genParamStreamSize, 0, logger),
		HeaderSize:            w.int64Param(streamGeneral, genParamHeaderSize, 0, logger),
		DataSize:              w.int64Param(streamGeneral, genParamDataSize, 0, logger),
		FooterSize:            w.int64Param(streamGeneral, genParamFooterSize, 0, logger),
		IsStreamable:          w.boolParam(streamGeneral, genParamIsStreamable, 0),
		EncodedDate:           w.timeParam(streamGeneral, genParamEncodedDate, 0, logger),
		TaggedDate:            w.timeParam(streamGeneral, genParamTaggedDate, 0, logger),
		FileModifiedDate:      w.timeParam(streamGeneral, genParamFileModifiedDate, 0, logger),
		FileModifiedDateLocal: w.localTimeParam(streamGeneral, genParamFileModifiedDateLocal, 0, logger),
		EncodedApplication:    w.stringParam(streamGeneral, genParamEncodedApplication, 0),
	}

	return &MediaInfo{
		File:        filename,
		TmpFile:     url,
		General:     generalInfo,
		VideoTracks: videoTracksFrom(w, generalInfo.VideoTrackCount.Val, logger),
		AudioTracks: audioTracksFrom(w, generalInfo.AudioTrackCount.Val, logger),
	}, nil
}

func videoTracksFrom(w *wrapper, numTracks int, logger *log.Logger) []VideoTrack {
	var tracks []VideoTrack

	for i := 0; i < numTracks; i++ {
		format := w.stringParam(streamVideo, videoParamFormat, i)
		profile := w.stringParam(streamVideo, videoParamFormatProfile, i)

		tracks = append(tracks, VideoTrack{
			StreamOrder:           w.intParam(streamVideo, videoParamStreamOrder, i, logger),
			ID:                    w.intParam(streamVideo, videoParamID, i, logger),
			Format:                format,
			Profile:               profile,
			FormatLevel:           w.stringParam(streamVideo, videoParamFormatLevel, i),
			IsCABACEnabled:        w.boolParam(streamVideo, videoParamFormatSettingsCABAC, i),
			RefFrames:             w.intParam(streamVideo, videoParamFormatSettingsRefFrames, i, logger),
			CodecID:               w.stringParam(streamVideo, videoParamCodecID, i),
			Duration:              w.float64Param(streamVideo, videoParamDuration, i, logger),
			Bitrate:               w.intParam(streamVideo, videoParamBitRate, i, logger),
			Width:                 w.intParam(streamVideo, videoParamWidth, i, logger),
			Height:                w.intParam(streamVideo, videoParamHeight, i, logger),
			SampledWidth:          w.intParam(streamVideo, videoParamSampledWidth, i, logger),
			SampledHeight:         w.intParam(streamVideo, videoParamSampledHeight, i, logger),
			PixelAspectRatio:      w.float64Param(streamVideo, videoParamPixelAspectRatio, i, logger),
			DisplayAspectRatio:    w.float64Param(streamVideo, videoParamDisplayAspectRatio, i, logger),
			Rotation:              w.float64Param(streamVideo, videoParamRotation, i, logger),
			FrameRateMode:         w.stringParam(streamVideo, videoParamFrameRateMode, i),
			FrameRateModeOriginal: w.stringParam(streamVideo, videoParamFrameRateModeOriginal, i),
			FrameRate:             w.float64Param(streamVideo, videoParamFrameRate, i, logger),
			FrameCount:            w.intParam(streamVideo, videoParamFrameCount, i, logger),
			ColorSpace:            w.stringParam(streamVideo, videoParamColorSpace, i),
			ChromaSubsampling:     w.stringParam(streamVideo, videoParamChromaSubsampling, i),
			BitDepth:              bitDepthFrom(w, format, profile, i, logger),
			ScanType:              w.stringParam(streamVideo, videoParamScanType, i),
			StreamSize:            w.int64Param(streamVideo, videoParamStreamSize, i, logger),
			EncodedDate:           w.timeParam(streamVideo, videoParamEncodedDate, i, logger),
			TaggedDate:            w.timeParam(streamVideo, videoParamTaggedDate, i, logger),
		})
	}

	return tracks
}

func audioTracksFrom(w *wrapper, numTracks int, logger *log.Logger) []AudioTrack {
	var tracks []AudioTrack

	for i := 0; i < numTracks; i++ {
		tracks = append(tracks, AudioTrack{
			StreamOrder:              w.intParam(streamAudio, audioParamStreamOrder, i, logger),
			ID:                       w.intParam(streamAudio, audioParamID, i, logger),
			Format:                   w.stringParam(streamAudio, audioParamFormat, i),
			FormatAdditionalFeatures: w.stringParam(streamAudio, audioParamFormatAdditionalFeatures, i),
			CodecID:                  w.stringParam(streamAudio, audioParamCodecID, i),
			Duration:                 w.float64Param(streamAudio, audioParamDuration, i, logger),
			BitrateMode:              w.stringParam(streamAudio, audioParamBitRateMode, i),
			Bitrate:                  w.intParam(streamAudio, audioParamBitRate, i, logger),
			BitrateMaximum:           w.intParam(streamAudio, audioParamBitRateMaximum, i, logger),
			Channels:                 w.intParam(streamAudio, audioParamChannels, i, logger),
			ChannelPositions:         w.stringParam(streamAudio, audioParamChannelPositions, i),
			ChannelLayout:            w.stringParam(streamAudio, audioParamChannelLayout, i),
			SamplesPerFrame:          w.intParam(streamAudio, audioParamSamplesPerFrame, i, logger),
			SamplingRate:             w.intParam(streamAudio, audioParamSamplingRate, i, logger),
			SamplingCount:            w.intParam(streamAudio, audioParamSamplingCount, i, logger),
			FrameRate:                w.float64Param(streamAudio, audioParamFrameRate, i, logger),
			FrameCount:               w.intParam(streamAudio, audioParamFrameCount, i, logger),
			CompressionMode:          w.stringParam(streamAudio, audioParamCompressionMode, i),
			StreamSize:               w.int64Param(streamAudio, audioParamStreamSize, i, logger),
			StreamSizeProportion:     w.float64Param(streamAudio, audioParamStreamSizeProportion, i, logger),
			IsDefault:                w.boolParam(streamAudio, audioParamDefault, i),
			AlternateGroup:           w.stringParam(streamAudio, audioParamAlternateGroup, i),
			EncodedDate:              w.timeParam(streamAudio, audioParamEncodedDate, i, logger),
			TaggedDate:               w.timeParam(streamAudio, audioParamTaggedDate, i, logger),
		})
	}

	return tracks
}

func bitDepthFrom(w *wrapper, format, profile StringValue, i int, logger *log.Logger) IntValue {
	if format.Val == videoFormatProRes && profile.Val == videoProfile4444 {
		return IntValue{Val: bitDepth12}
	} else if format.Val == videoFormatProRes && profile.Val == videoProfile422HQ {
		return IntValue{Val: bitDepth10}
	}

	return w.intParam(streamVideo, videoParamBitDepth, i, logger)
}
