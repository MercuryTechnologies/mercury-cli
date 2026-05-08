package debugmiddleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDebugMiddleware(t *testing.T) {
	t.Parallel()

	setup := func() (*RequestLogger, *bytes.Buffer) {
		var (
			logBuf     bytes.Buffer
			middleware = NewRequestLogger()
		)
		middleware.logger = log.New(&logBuf, "", 0)
		return middleware, &logBuf
	}

	t.Run("DoesNotRedactMostHeaders", func(t *testing.T) {
		t.Parallel()

		middleware, logBuf := setup()

		const stainlessUserAgent = "Stainless"

		req := httptest.NewRequest("GET", "https://example.com", nil)
		req.Header.Set("User-Agent", stainlessUserAgent)

		var nextMiddlewareRan bool
		middleware.Middleware()(req, func(req *http.Request) (*http.Response, error) {
			nextMiddlewareRan = true

			// The request sent down through middleware shouldn't be mutated.
			require.Equal(t, stainlessUserAgent, req.Header.Get("User-Agent"))

			return &http.Response{}, nil
		})

		require.True(t, nextMiddlewareRan)
		require.Contains(t, logBuf.String(), "User-Agent: "+stainlessUserAgent)
	})

	const secretToken = "secret-token"

	t.Run("RedactsAuthorizationHeader", func(t *testing.T) {
		t.Parallel()

		middleware, logBuf := setup()

		req := httptest.NewRequest("GET", "https://example.com", nil)
		req.Header.Set("Authorization", secretToken)

		var nextMiddlewareRan bool
		middleware.Middleware()(req, func(req *http.Request) (*http.Response, error) {
			nextMiddlewareRan = true

			// The request sent down through middleware shouldn't be mutated.
			require.Equal(t, secretToken, req.Header.Get("Authorization"))

			return &http.Response{}, nil
		})

		require.True(t, nextMiddlewareRan)
		require.Contains(t, logBuf.String(), "Authorization: "+redactedPlaceholder)
	})

	t.Run("RedactsOnlySecretInAuthorizationHeader", func(t *testing.T) {
		t.Parallel()

		middleware, logBuf := setup()

		req := httptest.NewRequest("GET", "https://example.com", nil)
		req.Header.Set("Authorization", "Bearer "+secretToken)

		var nextMiddlewareRan bool
		middleware.Middleware()(req, func(req *http.Request) (*http.Response, error) {
			nextMiddlewareRan = true

			return &http.Response{}, nil
		})

		require.True(t, nextMiddlewareRan)
		require.Contains(t, logBuf.String(), "Authorization: Bearer "+redactedPlaceholder)
	})

	t.Run("RedactsMultipleAuthorizationHeaders", func(t *testing.T) {
		t.Parallel()

		middleware, logBuf := setup()

		req := httptest.NewRequest("GET", "https://example.com", nil)
		req.Header.Add("Authorization", secretToken+"1")
		req.Header.Add("Authorization", secretToken+"2")

		var nextMiddlewareRan bool
		middleware.Middleware()(req, func(req *http.Request) (*http.Response, error) {
			nextMiddlewareRan = true

			// The request sent down through middleware shouldn't be mutated.
			require.Equal(t, []string{secretToken + "1", secretToken + "2"}, req.Header.Values("Authorization"))

			return &http.Response{}, nil
		})

		require.True(t, nextMiddlewareRan)

		if strings.Count(logBuf.String(), "Authorization: "+redactedPlaceholder) != 2 {
			t.Error("expected exactly two redacted placeholders in authorization headers")
		}
	})

	const customAPIKeyHeader = "X-My-Api-Key"

	t.Run("RedactsSensitiveHeaders", func(t *testing.T) {
		t.Parallel()

		middleware, logBuf := setup()

		middleware.sensitiveHeaders = []string{customAPIKeyHeader}

		req := httptest.NewRequest("GET", "https://example.com", nil)
		req.Header.Set(customAPIKeyHeader, secretToken)

		var nextMiddlewareRan bool
		middleware.Middleware()(req, func(req *http.Request) (*http.Response, error) {
			nextMiddlewareRan = true

			// The request sent down through middleware shouldn't be mutated.
			require.Equal(t, secretToken, req.Header.Get(customAPIKeyHeader))

			return &http.Response{}, nil
		})

		require.True(t, nextMiddlewareRan)
		require.Contains(t, logBuf.String(), customAPIKeyHeader+": "+redactedPlaceholder)
	})

	t.Run("RedactsMultipleSensitiveHeaders", func(t *testing.T) {
		t.Parallel()

		middleware, logBuf := setup()

		middleware.sensitiveHeaders = []string{customAPIKeyHeader}

		req := httptest.NewRequest("GET", "https://example.com", nil)
		req.Header.Add(customAPIKeyHeader, secretToken+"1")
		req.Header.Add(customAPIKeyHeader, secretToken+"2")

		var nextMiddlewareRan bool
		middleware.Middleware()(req, func(req *http.Request) (*http.Response, error) {
			nextMiddlewareRan = true

			// The request sent down through middleware shouldn't be mutated.
			require.Equal(t, []string{secretToken + "1", secretToken + "2"}, req.Header.Values(customAPIKeyHeader))

			return &http.Response{}, nil
		})

		require.True(t, nextMiddlewareRan)
		require.Equal(t, 2, strings.Count(logBuf.String(), customAPIKeyHeader+": "+redactedPlaceholder))
	})

	t.Run("DoesNotConsumeRequestBodyWhenIoReader", func(t *testing.T) {
		t.Parallel()

		middleware, logBuf := setup()
		middleware.sensitiveHeaders = []string{customAPIKeyHeader}

		const bodyContent = "test request body content"
		bodyReader := strings.NewReader(bodyContent)

		req := httptest.NewRequest("POST", "https://example.com", bodyReader)
		req.Header.Set("Authorization", secretToken)

		var nextMiddlewareRan bool
		middleware.Middleware()(req, func(req *http.Request) (*http.Response, error) {
			nextMiddlewareRan = true

			// The request body should still be fully readable after the middleware runs
			body, err := io.ReadAll(req.Body)
			require.NoError(t, err)
			require.Equal(t, bodyContent, string(body))

			// The request sent down through middleware shouldn't be mutated.
			require.Equal(t, secretToken, req.Header.Get("Authorization"))

			return &http.Response{}, nil
		})

		require.True(t, nextMiddlewareRan)
		require.Contains(t, logBuf.String(), "Authorization: "+redactedPlaceholder)
	})

	// Response headers carry the same classes of secrets that the request-side
	// list redacts. `set-cookie` is on `sensitiveHeaders` because Mercury's
	// API session cookie (HttpOnly+Secure+SameSite=Strict) appears on every
	// authenticated response. Without redaction the cookie surfaces verbatim
	// in `--debug` output, which is commonly archived in CI logs.

	t.Run("RedactsSensitiveResponseHeaders", func(t *testing.T) {
		t.Parallel()

		middleware, logBuf := setup()

		const sessionCookie = "_SESSION=4hl5Nz9ATGm8CVMkQ7c; Path=/; HttpOnly; Secure; SameSite=Strict"
		const apiKeyResp = "key-leaked-in-response"

		req := httptest.NewRequest("GET", "https://example.com", nil)
		middleware.Middleware()(req, func(req *http.Request) (*http.Response, error) {
			resp := &http.Response{
				StatusCode: http.StatusOK,
				ProtoMajor: 1, ProtoMinor: 1,
				Header: http.Header{},
				Body:   io.NopCloser(strings.NewReader("")),
			}
			resp.Header.Add("Set-Cookie", sessionCookie)
			resp.Header.Add("api-key", apiKeyResp)
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		})

		out := logBuf.String()
		// Sensitive response headers must be redacted in the dump.
		require.NotContains(t, out, sessionCookie)
		require.NotContains(t, out, apiKeyResp)
		require.Contains(t, out, "Set-Cookie: "+redactedPlaceholder)
		require.Contains(t, out, "Api-Key: "+redactedPlaceholder)
		// Non-sensitive headers must remain intact.
		require.Contains(t, out, "Content-Type: application/json")
	})

	t.Run("DoesNotMutateOriginalResponseHeaders", func(t *testing.T) {
		t.Parallel()

		middleware, _ := setup()

		const sessionCookie = "_SESSION=raw-value"
		req := httptest.NewRequest("GET", "https://example.com", nil)

		var capturedResp *http.Response
		middleware.Middleware()(req, func(req *http.Request) (*http.Response, error) {
			resp := &http.Response{
				StatusCode: http.StatusOK,
				ProtoMajor: 1, ProtoMinor: 1,
				Header: http.Header{},
				Body:   io.NopCloser(strings.NewReader("")),
			}
			resp.Header.Set("Set-Cookie", sessionCookie)
			capturedResp = resp
			return resp, nil
		})

		// The downstream consumer must still see the unredacted Set-Cookie value
		// on the actual response — redaction only applies to the dumped log output.
		require.NotNil(t, capturedResp)
		require.Equal(t, sessionCookie, capturedResp.Header.Get("Set-Cookie"))
	})

	t.Run("DoesNotConsumeResponseBody", func(t *testing.T) {
		t.Parallel()

		middleware, _ := setup()

		const bodyContent = "downstream-must-still-read-this"

		req := httptest.NewRequest("GET", "https://example.com", nil)
		var capturedBody string
		middleware.Middleware()(req, func(req *http.Request) (*http.Response, error) {
			resp := &http.Response{
				StatusCode: http.StatusOK,
				ProtoMajor: 1, ProtoMinor: 1,
				Header: http.Header{},
				Body:   io.NopCloser(strings.NewReader(bodyContent)),
			}
			// Force a redaction path by including a sensitive header.
			resp.Header.Set("Set-Cookie", "x=y")

			result := resp
			middleware.Middleware()(req, func(_ *http.Request) (*http.Response, error) {
				return resp, nil
			})

			body, err := io.ReadAll(result.Body)
			require.NoError(t, err)
			capturedBody = string(body)
			return resp, nil
		})

		require.Equal(t, bodyContent, capturedBody)
	})

	t.Run("DoesNotRedactNonSensitiveResponseHeaders", func(t *testing.T) {
		t.Parallel()

		middleware, logBuf := setup()

		req := httptest.NewRequest("GET", "https://example.com", nil)
		middleware.Middleware()(req, func(req *http.Request) (*http.Response, error) {
			resp := &http.Response{
				StatusCode: http.StatusOK,
				ProtoMajor: 1, ProtoMinor: 1,
				Header: http.Header{},
				Body:   io.NopCloser(strings.NewReader("")),
			}
			resp.Header.Set("X-Mercury-Request-Id", "trace-abc-123")
			resp.Header.Set("Content-Type", "application/json; charset=utf-8")
			return resp, nil
		})

		out := logBuf.String()
		require.Contains(t, out, "X-Mercury-Request-Id: trace-abc-123")
		require.Contains(t, out, "Content-Type: application/json; charset=utf-8")
	})
}
