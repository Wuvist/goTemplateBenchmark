// Generated by ego.
// DO NOT EDIT

package ego

import (
	"fmt"
	"github.com/SlinSo/goTemplateBenchmark/model"
	"html"
	"io"
)

var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
//line footer.ego:1
func EgoFooter(w io.Writer) error {
//line footer.ego:2
	_, _ = io.WriteString(w, "\n<div class=\"footer\">copyright 2016</div>")
	return nil
}

//line header.ego:1
func EgoHeader(w io.Writer, title string) error {
//line header.ego:2
	_, _ = io.WriteString(w, "\n<title>")
//line header.ego:2
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", title)))
//line header.ego:2
	_, _ = io.WriteString(w, "'s Home Page</title>\n<div class=\"header\">Page Header</div>")
	return nil
}

//line index.ego:2
func EgoComplex(w io.Writer, u *model.User, nav []*model.Navigation, title string) error {
//line index.ego:2
	_, _ = io.WriteString(w, "\n")
//line index.ego:3
	_, _ = io.WriteString(w, "\n<!DOCTYPE html>\n<html>\n<body>\n\n<header>\n")
//line index.ego:8
	EgoHeader(w, title)
//line index.ego:9
	_, _ = io.WriteString(w, "\n</header>\n\n<nav>\n")
//line index.ego:12
	EgoNavigation(w, nav)
//line index.ego:13
	_, _ = io.WriteString(w, "\n</nav>\n\n<section>\n<div class=\"content\">\n\t<div class=\"welcome\">\n\t\t<h4>Hello ")
//line index.ego:18
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", u.FirstName)))
//line index.ego:18
	_, _ = io.WriteString(w, "</h4>\n\t\t\n\t\t<div class=\"raw\">")
//line index.ego:20
	_, _ = fmt.Fprintf(w, "%v", u.RawContent)
//line index.ego:20
	_, _ = io.WriteString(w, "</div>\n\t\t<div class=\"enc\">")
//line index.ego:21
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", u.EscapedContent)))
//line index.ego:21
	_, _ = io.WriteString(w, "</div>\n\t</div>\n\n")
//line index.ego:25
	for i := 1; i <= 5; i++ {
		if i == 1 {
//line index.ego:28
			_, _ = io.WriteString(w, "\n\t\t\t<p>")
//line index.ego:28
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", u.FirstName)))
//line index.ego:28
			_, _ = io.WriteString(w, " has ")
//line index.ego:28
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", i)))
//line index.ego:28
			_, _ = io.WriteString(w, " message</p>\n\t")
//line index.ego:29
		} else {
//line index.ego:30
			_, _ = io.WriteString(w, "\n\t\t\t<p>")
//line index.ego:30
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", u.FirstName)))
//line index.ego:30
			_, _ = io.WriteString(w, " has ")
//line index.ego:30
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", i)))
//line index.ego:30
			_, _ = io.WriteString(w, " messages</p>\n\t\t")
//line index.ego:31
		}
	}

//line index.ego:34
	_, _ = io.WriteString(w, "\n</div>\n</section>\n\n<footer>\n")
//line index.ego:38
	EgoFooter(w)
//line index.ego:39
	_, _ = io.WriteString(w, "\n</footer>\n\n</body>\n</html>")
	return nil
}

//line navigation.ego:2
func EgoNavigation(w io.Writer, nav []*model.Navigation) error {
//line navigation.ego:2
	_, _ = io.WriteString(w, "\n")
//line navigation.ego:3
	_, _ = io.WriteString(w, "\n<ul class=\"navigation\">\n")
//line navigation.ego:5
	for _, item := range nav {
//line navigation.ego:7
		_, _ = io.WriteString(w, "\n\t\t<li><a href=\"")
//line navigation.ego:7
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", item.Link)))
//line navigation.ego:7
		_, _ = io.WriteString(w, "\">")
//line navigation.ego:7
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", item.Item)))
//line navigation.ego:7
		_, _ = io.WriteString(w, "</a></li>\n")
//line navigation.ego:8
	}
//line navigation.ego:9
	_, _ = io.WriteString(w, "\n</ul>")
	return nil
}

//line simple.ego:2
func EgoSimple(w io.Writer, u *model.User) error {
//line simple.ego:2
	_, _ = io.WriteString(w, "\n")
//line simple.ego:3
	_, _ = io.WriteString(w, "\n<html>\n    <body>\n        <h1>")
//line simple.ego:5
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", u.FirstName)))
//line simple.ego:5
	_, _ = io.WriteString(w, "</h1>\n\n        <p>Here's a list of your favorite colors:</p>\n        <ul>\n        ")
//line simple.ego:9
	for _, colorName := range u.FavoriteColors {
//line simple.ego:10
		_, _ = io.WriteString(w, "\n            <li>")
//line simple.ego:10
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", colorName)))
//line simple.ego:10
		_, _ = io.WriteString(w, "</li>")
//line simple.ego:10
	}
//line simple.ego:11
	_, _ = io.WriteString(w, "\n        </ul>\n    </body>\n</html>")
	return nil
}
