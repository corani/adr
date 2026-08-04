package app

import (
	"testing"
)

func TestParseFormat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input   string
		want    Format
		wantErr bool
	}{
		{"md", FormatMd, false},
		{"raw", FormatRaw, false},
		{"json", FormatJSON, false},
		{"", FormatMd, true},
		{"html", FormatMd, true},
		{"MD", FormatMd, true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			got, err := ParseFormat(test.input)

			if (err != nil) != test.wantErr {
				t.Errorf("ParseFormat(%q) error = %v, wantErr %v", test.input, err, test.wantErr)
			}

			if !test.wantErr && got != test.want {
				t.Errorf("ParseFormat(%q) = %v, want %v", test.input, got, test.want)
			}
		})
	}
}
