package elasticsearch

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/americanas-go/log"
)

// Logger implements the estransport.Logger interface.
type Logger struct {
}

// LogRoundTrip prints the information about request and response.
func (l *Logger) LogRoundTrip(
	req *http.Request,
	res *http.Response,
	err error,
	start time.Time,
	dur time.Duration,
) error {
	var (
		nReq int64
		nRes int64
	)

	// Count number of bytes in request and response.
	//
	if req != nil && req.Body != nil && req.Body != http.NoBody {
		nReq, _ = io.Copy(ioutil.Discard, req.Body)
	}
	if res != nil && res.Body != nil && res.Body != http.NoBody {
		nRes, _ = io.Copy(ioutil.Discard, res.Body)
	}

	// Log event.
	//
	logger := log.WithFields(log.Fields{
		"method":      req.Method,
		"status_code": res.StatusCode,
		"duration":    dur,
		"req_bytes":   nReq,
		"res_bytes":   nRes,
		"url":         req.URL.String(),
	})

	switch {
	case err != nil:
		logger.Error(err)
	case res != nil && res.StatusCode > 0 && res.StatusCode < 300:
		logger.Info("success")
	case res != nil && res.StatusCode > 299 && res.StatusCode < 500:
		logger.Warn("warning")
	default:
		logger.Error("error")
	}

	return nil
}

// RequestBodyEnabled makes the client pass request body to logger.
func (l *Logger) RequestBodyEnabled() bool { return true }

// RequestBodyEnabled makes the client pass response body to logger.
func (l *Logger) ResponseBodyEnabled() bool { return true }
