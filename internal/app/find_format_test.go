package app

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/corani/adr/internal/adr"
)

func TestFindJSON(t *testing.T) {
	t.Parallel()

	entries := []adrListEntry{
		{Number: 3, Title: "Use Kafka", Status: "deprecated", Date: "2024-03-01", Filepath: "0003-use-kafka.md"},
	}

	var buf bytes.Buffer

	if err := findJSON(&buf, entries); err != nil {
		t.Fatalf("findJSON: %v", err)
	}

	var got []adrListEntry

	if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal: %v\noutput: %s", err, buf.String())
	}

	if len(got) != 1 || got[0] != entries[0] {
		t.Errorf("got %+v, want %+v", got, entries)
	}
}

func TestFindMarkdownRaw(t *testing.T) {
	t.Parallel()

	entries := []adrListEntry{
		{Number: 3, Title: "Use Kafka", Status: "deprecated", Date: "2024-03-01", Filepath: "0003-use-kafka.md"},
	}

	var buf bytes.Buffer

	if err := renderMarkdownTable(&buf, entries, FormatRaw, "find"); err != nil {
		t.Fatalf("renderMarkdownTable: %v", err)
	}

	out := buf.String()

	const wantRow = "| 0003 | 2024-03-01 | deprecated | Use Kafka |"

	if !strings.Contains(out, wantRow) {
		t.Errorf("output missing %q\ngot: %s", wantRow, out)
	}
}

func TestMatchesFormat(t *testing.T) {
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

	re := buildQuery("postgres")

	if !matches(re, entry, false) {
		t.Error("expected title match")
	}
}
