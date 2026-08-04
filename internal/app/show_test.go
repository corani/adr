package app

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/corani/adr/internal/adr"
)

func TestShowJSON(t *testing.T) {
	t.Parallel()

	entry := &adr.Adr{
		Filename: testFilename,
		Type:     "",
		Number:   1,
		Title:    testTitle,
		Status:   adr.StatusAccepted,
		Date:     testDate,
		Link:     0,
		Body:     []byte("  We chose PostgreSQL.\n"),
	}

	var buf bytes.Buffer

	if err := showJSON(&buf, entry); err != nil {
		t.Fatalf("showJSON: %v", err)
	}

	var got adrShowEntry

	if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal: %v\noutput: %s", err, buf.String())
	}

	if got.Number != 1 {
		t.Errorf("Number = %d, want 1", got.Number)
	}

	if got.Title != testTitle {
		t.Errorf("Title = %q", got.Title)
	}

	if got.Status != testStatus {
		t.Errorf("Status = %q", got.Status)
	}

	if got.Body != "We chose PostgreSQL." {
		t.Errorf("Body = %q", got.Body)
	}
}

func TestShowMarkdownRaw(t *testing.T) {
	t.Parallel()

	entry := &adr.Adr{
		Filename: testFilename,
		Type:     "",
		Number:   1,
		Title:    testTitle,
		Status:   adr.StatusAccepted,
		Date:     testDate,
		Link:     0,
		Body:     []byte("We chose PostgreSQL.\n"),
	}

	var buf bytes.Buffer

	if err := showMarkdown(&buf, entry, FormatRaw); err != nil {
		t.Fatalf("showMarkdown: %v", err)
	}

	out := buf.String()

	if !strings.Contains(out, "| Status | "+testStatus+" |") {
		t.Errorf("output missing status row\ngot: %s", out)
	}

	if !strings.Contains(out, "We chose PostgreSQL.") {
		t.Errorf("output missing body\ngot: %s", out)
	}
}
