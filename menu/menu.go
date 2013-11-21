package menu

import (
	"bytes"
	"html/template"
)

type Menu struct {
	Items []*menuItem
}

type menuItem struct {
	Name, Url string
	Items     []*menuItem
}

func New() *Menu {
	return new(Menu)
}

func (self *Menu) AddItem(name, url string) *Menu {
	self.Items = append(self.Items, &menuItem{name, url, nil})
	return self
}

func (self *Menu) AddSubItem(parent, name, url string) *Menu {
	for key, item := range self.Items {
		if item.Name != parent {
			continue
		}
		self.Items[key].Items = append(self.Items[key].Items, &menuItem{name, url, nil})
	}
	return self
}

func (self *Menu) Sort() {

}

func (self *Menu) Render() template.HTML {
	// Create a new template and parse the letter into it.
	var out bytes.Buffer
	tMenu := template.Must(template.New("menu").Parse(tmplMenu))
	tMap := map[string]interface{}{
		"menu": self,
	}
	tMenu.Execute(&out, tMap)
	return template.HTML(out.String())
}

const tmplMenu = `
<ul class="nav bs-sidenav">
  {{ range $index, $element := .menu.Items }}
    {{ if len $element.Items }}
      <li>
          <a href="#">{{$element.Name}}</a>
          <ul class="nav">
            {{ range $element.Items }}
              <li><a href="{{.Url}}">{{.Name}}</a></li>
            {{ end }}
          </ul>
      </li>
    {{ else }}
      <li><a href="{{$element.Url}}">{{$element.Name}}</a></li>
    {{ end }}
  {{ end }}
<ul>
`
