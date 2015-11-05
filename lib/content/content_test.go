package content

import (
	"strings"
	"testing"
)

func TestHTML(t *testing.T) {
	r := strings.NewReader(`
# Main title

## Subtitle

This is some text with a [link](https://www.golang.hr).
`)

	html, err := HTML(r)

	if err != nil {
		t.Fatalf("Content parsing failed with error %s", err)
	}

	tests := []struct {
		want string
	}{
		{"<h1>Main title</h1>"},
		{"<h2>Subtitle</h2>"},
		{"<a href=\"https://www.golang.hr\">link</a>"},
	}

	for _, tt := range tests {
		if !strings.Contains(html, tt.want) {
			t.Errorf("Expected %q not found", tt.want)
		}
	}
}
