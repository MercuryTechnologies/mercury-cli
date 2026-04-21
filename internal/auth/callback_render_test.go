package auth

import (
	"html"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderSuccess(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	renderSuccess(w)

	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, strings.HasPrefix(resp.Header.Get("Content-Type"), "text/html"))

	body := w.Body.String()
	assert.Contains(t, body, "<!DOCTYPE html>")
	assert.Contains(t, body, "You're signed in")
	assert.Contains(t, body, "You can close this window and return to your terminal.")
	assert.Contains(t, body, `aria-label="Mercury"`)
}

func TestRenderError(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		title  string
		detail string
	}{
		{
			name:   "state mismatch",
			title:  "That link didn't look right",
			detail: "For security, the sign-in link couldn't be verified. Please try again from your terminal.",
		},
		{
			name:   "missing code",
			title:  "Something went wrong",
			detail: "We didn't receive a sign-in code from Mercury. Please try again from your terminal.",
		},
		{
			name:   "upstream error",
			title:  "Sign-in didn't complete",
			detail: "User denied consent.",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			renderError(w, tc.title, tc.detail, nil)

			resp := w.Result()
			assert.Equal(t, 200, resp.StatusCode)
			assert.True(t, strings.HasPrefix(resp.Header.Get("Content-Type"), "text/html"))

			body := html.UnescapeString(w.Body.String())
			assert.Contains(t, body, tc.title)
			assert.Contains(t, body, tc.detail)
			assert.Contains(t, body, `aria-label="Mercury"`)
			assert.NotContains(t, body, `class="cta"`, "no link should be rendered when link is nil")
		})
	}
}

func TestRenderError_WithLink(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	renderError(w, "Title", "Detail.", &errorLink{
		URL:  "https://app.mercury.com/settings/users",
		Text: "Open Team Settings",
	})

	body := w.Body.String()
	assert.Contains(t, body, `href="https://app.mercury.com/settings/users"`)
	assert.Contains(t, body, "Open Team Settings")
}

func TestRenderError_EscapesDetail(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	renderError(w, "Bad", `<script>alert("xss")</script>`, nil)

	body := w.Body.String()
	assert.NotContains(t, body, `<script>alert("xss")</script>`)
	assert.Contains(t, body, "&lt;script&gt;")
}
