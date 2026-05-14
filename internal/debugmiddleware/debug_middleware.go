package debugmiddleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"reflect"
	"strings"
)

// For the time being these type definitions are duplicated here so that we can
// test this file in a non-generated context.
type (
	Middleware     = func(*http.Request, MiddlewareNext) (*http.Response, error)
	MiddlewareNext = func(*http.Request) (*http.Response, error)
)

const redactedPlaceholder = "<REDACTED>"

// Headers known to contain sensitive information like an API key. Note that this exclude `Authorization`,
// which is handled specially in `redactRequest` below.
var sensitiveHeaders = []string{
	"api-key",
	"x-api-key",
	"cookie",
	"set-cookie",
}

// RequestLogger is a middleware that logs HTTP requests and responses.
type RequestLogger struct {
	logger           interface{ Printf(string, ...any) } // field for testability; usually log.Default()
	sensitiveHeaders []string                            // field for testability; usually sensitiveHeaders
}

// NewRequestLogger returns a new RequestLogger instance with default options.
func NewRequestLogger() *RequestLogger {
	return &RequestLogger{
		logger:           log.Default(),
		sensitiveHeaders: sensitiveHeaders,
	}
}

func (m *RequestLogger) Middleware() Middleware {
	return func(req *http.Request, mn MiddlewareNext) (*http.Response, error) {
		redacted, err := m.redactRequest(req)
		if err != nil {
			return nil, err
		}
		if reqBytes, err := httputil.DumpRequest(redacted, true); err == nil {
			m.logger.Printf("Request Content:\n%s\n", reqBytes)
		}

		resp, err := mn(req)
		if err != nil {
			return resp, err
		}

		if err := m.dumpRedactedResponse(resp); err != nil {
			return resp, err
		}

		return resp, err
	}
}

// redactHeaders returns a clone of h with sensitive headers replaced by the
// placeholder. Authorization is redacted with its token-kind prefix preserved
// (e.g. "Bearer <REDACTED>"). Every header in m.sensitiveHeaders is replaced
// in full. The clone is always allocated so callers can compare the result to
// the original to detect whether any redaction actually occurred.
func (m *RequestLogger) redactHeaders(h http.Header) http.Header {
	redactedHeaders := h.Clone()

	// Notably, the clauses below are written so they can redact multiple
	// headers of the same name if necessary.
	if values := redactedHeaders.Values("Authorization"); len(values) > 0 {
		redactedHeaders.Del("Authorization")

		for _, value := range values {
			// In case we're using something like a bearer token (e.g. `Bearer
			// <my_token>`), keep the `Bearer` part for more debugging
			// information.
			if authKind, _, ok := strings.Cut(value, " "); ok {
				redactedHeaders.Add("Authorization", authKind+" "+redactedPlaceholder)
			} else {
				redactedHeaders.Add("Authorization", redactedPlaceholder)
			}
		}
	}

	for _, header := range m.sensitiveHeaders {
		values := redactedHeaders.Values(header)
		if len(values) == 0 {
			continue
		}

		redactedHeaders.Del(header)

		for range values {
			redactedHeaders.Add(header, redactedPlaceholder)
		}
	}

	return redactedHeaders
}

// redactRequest redacts sensitive information from the request for logging
// purposes. If redaction is necessary, the request is cloned before mutating
// the original and that clone is returned. As a small optimization, the
// original is request is returned unchanged if no redaction is necessary.
func (m *RequestLogger) redactRequest(req *http.Request) (*http.Request, error) {
	redactedHeaders := m.redactHeaders(req.Header)

	if reflect.DeepEqual(req.Header, redactedHeaders) {
		return req, nil
	}

	redacted := req.Clone(req.Context())
	redacted.Header = redactedHeaders
	var err error
	redacted.Body, req.Body, err = cloneBody(req.Body)
	return redacted, err
}

// dumpRedactedResponse logs the response with sensitive headers redacted.
// `set-cookie` is on `sensitiveHeaders` because Mercury's API session cookie
// (HttpOnly+Secure+SameSite=Strict, e.g. `_SESSION=...`) appears on every
// authenticated response. Without redaction the cookie surfaces verbatim in
// `--debug` output, which is commonly archived in CI logs and developer
// shell history. The previous implementation only redacted request headers;
// this mirrors the same redaction on the response. Body bytes are buffered
// via cloneBody so the downstream caller can still consume `resp.Body` after
// the dump, matching the original behavior.
func (m *RequestLogger) dumpRedactedResponse(resp *http.Response) error {
	redactedHeaders := m.redactHeaders(resp.Header)

	if reflect.DeepEqual(resp.Header, redactedHeaders) {
		if respBytes, err := httputil.DumpResponse(resp, true); err == nil {
			m.logger.Printf("Response Content:\n%s\n", respBytes)
		}
		return nil
	}

	respClone := *resp
	respClone.Header = redactedHeaders
	var err error
	respClone.Body, resp.Body, err = cloneBody(resp.Body)
	if err != nil {
		return err
	}
	if respBytes, dumpErr := httputil.DumpResponse(&respClone, true); dumpErr == nil {
		m.logger.Printf("Response Content:\n%s\n", respBytes)
	}
	return nil
}

// This function returns two copies of an HTTP request body that can each be
// read independently without affecting the other.
// This logic is taken from `drainBody` in net/http/httputil.
func cloneBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == nil || b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return io.NopCloser(&buf), io.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
