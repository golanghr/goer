package content

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseMkToHTML(t *testing.T) {
	r := strings.NewReader(`
# Main title

## Subtitle

This is some text with a [link](https://www.google.com).
`)

	html, err := HTML(r)

	if err != nil {
		t.Fatal(fmt.Sprintf("Content parsing failed with error %s", err))
	}

	if !strings.Contains(html, "<h1>Main title</h1>") {
		t.Fatal(fmt.Sprintf("Expected \"%s\" not recived in \"%s\"", "<h1>Main title</h1>", html))
	}
	if !strings.Contains(html, "<h2>Subtitle</h2>") {
		t.Fatal(fmt.Sprintf("Expected \"%s\" not recived in \"%s\"", "<h2>Subtitle</h2>", html))
	}
	if !strings.Contains(html, "<a href=\"https://www.google.com\">link</a>") {
		t.Fatal(fmt.Sprintf("Expected \"%s\" not recived in \"%s\"", "<a href=\"https://www.google.com\">link</a>", html))
	}
}
