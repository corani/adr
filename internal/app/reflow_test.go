package app

import "testing"

func TestReflowBody(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "joins indented continuation lines",
			input: "- item one with a long description that\n  wraps onto the next line\n- item two\n",
			want:  "- item one with a long description that wraps onto the next line\n- item two\n",
		},
		{
			name:  "does not join across blank line",
			input: "- item one\n  \n  indented after blank\n",
			want:  "- item one\n  \n  indented after blank\n",
		},
		{
			name:  "does not join across double newline",
			input: "- item one\n\n  indented after blank\n",
			want:  "- item one\n\n  indented after blank\n",
		},
		{
			name:  "does not join paragraph soft breaks",
			input: "A paragraph with\na manual break.\n",
			want:  "A paragraph with\na manual break.\n",
		},
		{
			name:  "multiple continuation lines",
			input: "- The service's logic is mature\n  and well-tested.\n  Minimising changes reduces risk.\n- item two\n",
			want:  "- The service's logic is mature and well-tested. Minimising changes reduces risk.\n- item two\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := reflowBody(tt.input); got != tt.want {
				t.Errorf("got:\n%q\nwant:\n%q", got, tt.want)
			}
		})
	}
}
