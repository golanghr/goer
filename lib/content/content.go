package content

import (
	"io"
	"io/ioutil"

	"github.com/russross/blackfriday"
)

// HTML converst content of reader containg markdown content to HTML
func HTML(r io.Reader) (html string, err error) {
	mk, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	htmlbytes := blackfriday.MarkdownCommon(mk)

	return string(htmlbytes), nil
}
