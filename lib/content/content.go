package content

import (
	"io"
	"io/ioutil"

	"github.com/russross/blackfriday"
)

// HTML converts markdown.
func HTML(r io.Reader) (string, error) {
	md, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	htmlBytes := blackfriday.MarkdownCommon(md)

	return string(htmlBytes), nil
}
