// Code generated by hero.
// source: /home/sl/gocode/src/github.com/SlinSo/goTemplateBenchmark/hero/simple.html
// DO NOT EDIT!
package template

import (
	"bytes"
	"html"

	"github.com/shiyanhui/hero"
)

import "github.com/SlinSo/goTemplateBenchmark/model"

func SimpleQtc(u *model.User) *bytes.Buffer {
	_buffer := hero.GetBuffer()
	_buffer.WriteString(`

<html>
    <body>
        <h1>`)
	_buffer.WriteString(html.EscapeString(u.FirstName))
	_buffer.WriteString(`</h1>

        <p>Here's a list of your favorite colors:</p>
        <ul>
            `)
	for _, colorName := range u.FavoriteColors {
		_buffer.WriteString(`
                <li>`)
		_buffer.WriteString(html.EscapeString(colorName))
		_buffer.WriteString(`</li>
            `)
	}
	_buffer.WriteString(`
        </ul>
    </body>
</html>
`)

	return _buffer
}
