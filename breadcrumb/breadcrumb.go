package breadcrumb

import (
	"bytes"
	"html/template"
)

type Breadcrumb struct {
	Items []*breadcrumbItem
}

type breadcrumbItem struct {
	Name, Url string
}

func New() *Breadcrumb {
	return new(Breadcrumb)
}

func (self *Breadcrumb) Add(name, url string) *Breadcrumb {
	self.Items = append(self.Items, &breadcrumbItem{name, url})
	return self
}

func (self *Breadcrumb) Render() template.HTML {
	// Create a new template and parse the letter into it.
	var out bytes.Buffer
	tBreadcrumb := template.Must(template.New("breadcrumb").Parse(tmplBreadcrumb))
	tMap := map[string]interface{}{
		"breadcrumb": self,
	}
	tBreadcrumb.Execute(&out, tMap)
	return template.HTML(out.String())
}

const tmplBreadcrumb = `
<ol class="breadcrumb">
  {{ range .breadcrumb.Items }}
    {{ if len .Url }}
      <li><a href="{{.Url}}">{{.Name}}</a></li>
    {{ else }}
      <li class="active">{{.Name}}</li>
    {{ end }}
  {{ end }}
</ol>
`
