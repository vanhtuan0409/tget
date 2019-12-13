package tget

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path"
	"time"
)

func getDownloadedPath(target string, resp *http.Response) string {
	st, err := os.Stat(target)
	if err == nil && st.IsDir() {
		contentDisposition := resp.Header.Get(contentDispositionHeader)
		_, params, err := mime.ParseMediaType(contentDisposition)
		if err != nil {
			return path.Join(target, fmt.Sprintf("tget-file-%d", time.Now().Unix()))
		}
		fName, ok := params["filename"]
		if !ok {
			return path.Join(target, fmt.Sprintf("tget-file-%d", time.Now().Unix()))
		}
		return path.Join(target, fName)
	}
	return target
}
