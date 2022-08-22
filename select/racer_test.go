package racer

import "testing"

func TestRecer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.co.uk"

	want := fastURL
	got := Recer(slowURL, fastURL)

	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
}
