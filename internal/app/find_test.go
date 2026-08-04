package app

import (
	"testing"

	"github.com/corani/adr/internal/adr"
)

func TestBuildQuery(t *testing.T) {
	t.Parallel()

	const fooBar = "foo bar"

	const fooDotBar = "foo.bar"

	tests := []struct {
		query string
		input string
		want  bool
	}{
		{"foo", fooBar, true},
		{"foo", "bar baz", false},
		{fooBar, "foo baz bar", true},
		{fooBar, "bar foo", false},
		{fooBar, "foobar", true},
		{fooDotBar, fooDotBar, true},
		{fooDotBar, "fooXbar", false},
	}

	for _, test := range tests {
		t.Run(test.query+"/"+test.input, func(t *testing.T) {
			t.Parallel()

			re := buildQuery(test.query)

			if got := re.MatchString(test.input); got != test.want {
				t.Errorf("buildQuery(%q).MatchString(%q) = %v, want %v", test.query, test.input, got, test.want)
			}
		})
	}
}

func TestMatches(t *testing.T) {
	t.Parallel()

	entry := &adr.Adr{
		Filename: testFilename,
		Type:     "",
		Number:   1,
		Title:    testTitle,
		Status:   adr.StatusAccepted,
		Date:     testDate,
		Link:     0,
		Body:     []byte("We chose PostgreSQL because it supports JSONB."),
	}

	tests := []struct {
		name     string
		query    string
		fullText bool
		want     bool
	}{
		{"title match", "postgres storage", false, true},
		{"title case-insensitive", "POSTGRES", false, true},
		{"title no match", "mysql", false, false},
		{"body not searched without flag", "jsonb", false, false},
		{"body searched with flag", "jsonb", true, true},
		{"status searched with flag", testStatus, true, true},
		{"date searched with flag", "2024-01", true, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			re := buildQuery(test.query)

			if got := matches(re, entry, test.fullText); got != test.want {
				t.Errorf("matches(%q, fullText=%v) = %v, want %v", test.query, test.fullText, got, test.want)
			}
		})
	}
}
