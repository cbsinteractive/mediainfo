package mediainfo

import (
	"context"
	"encoding/json"
	"net/url"
	"os/exec"
	"path"
)

func Analyze(u string) (MediaInfo, error) {
	return AnalyzeWithContext(context.Background(), u)
}

func AnalyzeWithContext(ctx context.Context, urlStr string) (MediaInfo, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return MediaInfo{}, err
	}

	minfo := cliMediaInfo{}

	out, err := exec.CommandContext(ctx, "mediainfo", "--Output=JSON", urlStr).Output()
	if err != nil {
		return MediaInfo{}, err
	}

	if err := json.Unmarshal(out, &minfo); err != nil {
		return MediaInfo{}, err
	}

	info := minfo.toMediaInfo()
	info.File = path.Base(u.Path)
	info.TmpFile = urlStr

	return info, nil
}
