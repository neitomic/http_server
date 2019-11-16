package response

import (
	"io"
	"net/http"
	"time"
)

type BinaryResp struct {
	Name string
	ModTime time.Time
	ContentType string
	HttpHeader http.Header
	Reader io.ReadSeeker
}