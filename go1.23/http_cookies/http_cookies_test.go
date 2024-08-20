package http_cookies

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ParseCookie parses a Cookie header value and returns all the cookies
// which were set in it. Since the same cookie name can appear multiple times
// the returned Values can contain more than one value for a given key.
func TestParseCookie(t *testing.T) {
	line := "session_id=abc123; dnt=1; lang=en; lang=de"
	cookies, err := http.ParseCookie(line)
	if err != nil {
		t.Fatalf("Failed to parse cookies: %v", err)
	}

	expected := []http.Cookie{
		{Name: "session_id", Value: "abc123"},
		{Name: "dnt", Value: "1"},
		{Name: "lang", Value: "en"},
		{Name: "lang", Value: "de"},
	}

	for i, cookie := range cookies {
		t.Logf("Parsed cookie: %+v", cookie)
		if expected[i].Value != cookie.Value {
			t.Errorf("Expected %s: %s, but got %s", cookie.Name, expected[i].Value, cookie.Value)
		}
	}
}

// ParseSetCookie parses a Set-Cookie header value and returns a cookie.
// It returns an error on syntax error.
//
// The "Partitioned" field identifies cookies with the Partitioned attribute
// (restricts the cookie's scope to a specific partition of the browsing context, such as a top-level site or a specific subdomain).
//
// The "Quoted" field indicates whether the value was originally quoted
func TestParseSetCookie(t *testing.T) {
	line := "session_id=abc123; SameSite=None; Secure; Partitioned; Path=/; Domain=.old.com; Domain=.new.com"
	cookie, err := http.ParseSetCookie(line)
	if err != nil {
		t.Fatalf("Failed to parse Set-Cookie: %v", err)
	}

	t.Logf("Parsed cookie: %+v", cookie)

	if cookie.Name != "session_id" || cookie.Value != "abc123" || !cookie.Secure || !cookie.Partitioned || cookie.Path != "/" || cookie.Domain != ".new.com" {
		t.Errorf("Failed to parse Set-Cookie correctly: %+v", cookie)
	}
}

// Test retrieving cookies by name using CookiesNamed
func TestCookiesNamed(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: "abc123"})
	req.AddCookie(&http.Cookie{Name: "session", Value: "123abc"})

	w := httptest.NewRecorder()
	handler := func(w http.ResponseWriter, r *http.Request) {
		cookies := r.CookiesNamed("session")
		if len(cookies) > 0 {
			fmt.Fprintf(w, "session cookie = %s", cookies[0].Value)
		} else {
			fmt.Fprint(w, "session cookie not found")
		}
	}
	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	if string(body) != "session cookie = abc123" {
		t.Errorf("Expected to find session cookie, but got: %s", string(body))
	}
}
