package lark_docs_md

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/chyroc/lark"
)

func downloadFile(ctx context.Context, fileToken, name string, opt *FormatOpt) string {
	if opt.LarkClient == nil {
		return ""
	}
	if opt.StaticAsURL {
		resp, _, err := opt.LarkClient.Drive.BatchGetDriveMediaTmpDownloadURL(ctx, &lark.BatchGetDriveMediaTmpDownloadURLReq{
			FileTokens: []string{fileToken},
		})
		if err != nil {
			log.Printf("lark get drive media tmp url %s fail: %s", fileToken, err)
			return ""
		}
		for _, v := range resp.TmpDownloadURLs {
			return v.TmpDownloadURL
		}
		return ""
	} else {
		resp, _, err := opt.LarkClient.Drive.DownloadDriveMedia(ctx, &lark.DownloadDriveMediaReq{
			FileToken: fileToken,
		})
		if err != nil {
			log.Printf("lark download drive media %s fail: %s", fileToken, err)
			return ""
		}
		filename := fmt.Sprintf("%s/%s", opt.StaticDir, name)
		mdname := fmt.Sprintf("%s/%s", opt.FilePrefix, name)
		_ = os.MkdirAll(filepath.Dir(filename), 0o755)
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0o666)
		if err != nil {
			log.Printf("open file %s fail: %s", filename, err)
			return ""
		}
		defer f.Close()

		_, _ = io.Copy(f, resp.File)
		return mdname
	}
}
