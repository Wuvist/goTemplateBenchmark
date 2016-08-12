// This file is automatically generated by qtc from "navigation.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line navigation.qtpl:1
package quicktemplate

//line navigation.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line navigation.qtpl:1
import "github.com/SlinSo/goTemplateBenchmark/model"

//line navigation.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line navigation.qtpl:2
func StreamNavigation(qw422016 *qt422016.Writer, nav []*model.Navigation) {
	//line navigation.qtpl:2
	qw422016.N().S(`
<ul class="navigation">
`)
	//line navigation.qtpl:4
	for _, item := range nav {
		//line navigation.qtpl:4
		qw422016.N().S(`
	<li><a href="`)
		//line navigation.qtpl:5
		qw422016.E().S(item.Link)
		//line navigation.qtpl:5
		qw422016.N().S(`">`)
		//line navigation.qtpl:5
		qw422016.E().S(item.Item)
		//line navigation.qtpl:5
		qw422016.N().S(`</a></li>
`)
		//line navigation.qtpl:6
	}
	//line navigation.qtpl:6
	qw422016.N().S(`
</ul>
`)
//line navigation.qtpl:8
}

//line navigation.qtpl:8
func WriteNavigation(qq422016 qtio422016.Writer, nav []*model.Navigation) {
	//line navigation.qtpl:8
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line navigation.qtpl:8
	StreamNavigation(qw422016, nav)
	//line navigation.qtpl:8
	qt422016.ReleaseWriter(qw422016)
//line navigation.qtpl:8
}

//line navigation.qtpl:8
func Navigation(nav []*model.Navigation) string {
	//line navigation.qtpl:8
	qb422016 := qt422016.AcquireByteBuffer()
	//line navigation.qtpl:8
	WriteNavigation(qb422016, nav)
	//line navigation.qtpl:8
	qs422016 := string(qb422016.B)
	//line navigation.qtpl:8
	qt422016.ReleaseByteBuffer(qb422016)
	//line navigation.qtpl:8
	return qs422016
//line navigation.qtpl:8
}
