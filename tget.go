package tget

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"os"
)

const (
	acceptRangeHeader        = "Accept-Ranges"
	contentDispositionHeader = "Content-Disposition"
	contentLengthHeader      = "Content-Length"
)

var (
	ErrCheckSum = errors.New("Checksum is not valid")
)

type Request struct {
	URL      string
	CheckSum string
}

func NewRequest(url string) *Request {
	r := new(Request)
	r.URL = url
	r.CheckSum = ""
	return r
}

func (r *Request) SetCheckSum(checksum string) {
	r.CheckSum = checksum
}

func (r *Request) Download(target string, parallel int) error {
	resp, err := http.Get(r.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if h := resp.Header.Get(acceptRangeHeader); h != "bytes" {
		parallel = 1
	}

	downloadedPath := getDownloadedPath(target, resp)
	out, err := os.Create(downloadedPath)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}

	if r.CheckSum != "" {
		out.Seek(0, io.SeekStart)
		hasher := sha256.New()
		if _, err := io.Copy(hasher, out); err != nil {
			return err
		}

		fCheckSum := hex.EncodeToString(hasher.Sum(nil))
		if r.CheckSum != "" && r.CheckSum != fCheckSum {
			return ErrCheckSum
		}
	}

	return nil
}
