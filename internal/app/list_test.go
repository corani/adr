package app

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestListJSON(t *testing.T) {
	t.Parallel()

	entries := []adrListEntry{
		{Number: 1, Title: "Use PostgreSQL", Status: testStatus, Date: testDate, Filepath: testFilename},
		{Number: 2, Title: "Use Redis", Status: "proposed", Date: "2024-02-01", Filepath: "0002-use-redis.md"},
	}

	var buf bytes.Buffer

	if err := listJSON(&buf, entries); err != nil {
		t.Fatalf("listJSON: %v", err)
	}

	var got []adrListEntry

	if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal: %v\noutput: %s", err, buf.String())
	}

	if len(got) != len(entries) {
		t.Fatalf("len = %d, want %d", len(got), len(entries))
	}

	for i, want := range entries {
		if got[i] != want {
			t.Errorf("[%d] = %+v, want %+v", i, got[i], want)
		}
	}
}

func TestListMarkdownRaw(t *testing.T) {
	t.Parallel()

	entries := []adrListEntry{
		{Number: 1, Title: "Use PostgreSQL", Status: testStatus, Date: testDate, Filepath: testFilename},
	}

	var buf bytes.Buffer

	if err := renderMarkdownTable(&buf, entries, FormatRaw, "list"); err != nil {
		t.Fatalf("renderMarkdownTable: %v", err)
	}

	out := buf.String()

	const wantHeader = "| # | date | status | title |"

	if !strings.Contains(out, wantHeader) {
		t.Errorf("output missing header %q\ngot: %s", wantHeader, out)
	}

	const wantRow = "| 0001 | 2024-01-15 | accepted | Use PostgreSQL |"

	if !strings.Contains(out, wantRow) {
		t.Errorf("output missing row %q\ngot: %s", wantRow, out)
	}

	const wantHint = "adr show <id>"

	if !strings.Contains(out, wantHint) {
		t.Errorf("output missing hint %q\ngot: %s", wantHint, out)
	}
}
