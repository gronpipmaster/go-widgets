package pagination

import (
	"bytes"
	"fmt"
	"html/template"
	"math"
)

type Pagination struct {
	total       int
	currentPage int
	limit       int
	urlPattern  string
	numLinks    int
	start       int
	end         int
	links       []*page
}

type page struct {
	Active bool
	Number int
	Link   string
}

// Example:
//         pagination := pagination.New(total, currentPage, pageSize, routes.Users.Index()+"page/")
// Example Print:
//				 pagination.Render() #or template usage {{.pagination.Render}}
//				 pagination.Summary() #or template usage {{.pagination.Summary}}
func New(total, currentPage, limit int, urlPattern string) *Pagination {
	p := new(Pagination)
	p.total = total
	p.currentPage = currentPage
	p.limit = limit
	p.urlPattern = urlPattern + "%d"
	p.numLinks = 20
	return p
}

func (self *Pagination) SetNumLinks(numLinks int) *Pagination {
	self.numLinks = numLinks
	return self
}

func (self *Pagination) Render() template.HTML {
	// Create a new template and parse the letter into it.
	var out bytes.Buffer
	tPagination := template.Must(template.New("pagination").Parse(tmplPagination))
	tMap := map[string]interface{}{
		"links": self.links,
	}
	tPagination.Execute(&out, tMap)
	return template.HTML(out.String())
}

func (self *Pagination) Summary() template.HTML {
	// Create a new template and parse the letter into it.
	var out bytes.Buffer
	tSummary := template.Must(template.New("summary").Parse(tmplSummary))
	tMap := map[string]interface{}{
		"start": self.start,
		"end":   self.end,
		"total": self.total,
	}
	tSummary.Execute(&out, tMap)
	return template.HTML(out.String())
}

func (self *Pagination) Init() {
	if self.currentPage < 1 {
		self.currentPage = 1
	}

	if self.limit == 0 {
		self.limit = 10
	}
	numPages := int(math.Ceil(float64(self.total) / float64(self.limit)))
	if numPages > 1 {
		if numPages < self.numLinks {
			self.start = 1
			self.end = numPages
		} else {
			self.start = self.currentPage - int(math.Floor(float64(self.numLinks)/float64(2)))
			self.end = self.currentPage + int(math.Floor(float64(self.numLinks)/float64(2)))

			if self.start < 1 {
				self.end += int(math.Abs(float64(self.start))) + 1
				self.start = 1
			}

			if self.end > numPages {
				self.start -= (self.end - numPages)
				self.end = numPages
			}
		}

		for i := self.start; i <= self.end; i++ {
			page := new(page)
			page.Number = i
			page.Link = fmt.Sprintf(self.urlPattern, page.Number)
			if page.Number == self.currentPage {
				page.Active = true
			} else {
				page.Active = false
			}
			self.links = append(self.links, page)
		}
	}

	if self.total > 0 {
		self.start = ((self.currentPage - 1) * self.limit) + 1
	}

	if ((self.currentPage - 1) * self.limit) > (self.total - self.limit) {
		self.end = self.total
	} else {
		self.end = ((self.currentPage - 1) * self.limit) + self.limit
	}
}

const (
	tmplPagination = `
  {{if .links}}
  <ul class="pagination">
    {{range .links}}
      {{ if .Active }}
        <li class="active"><a href="#">{{.Number}}</a></li>
      {{ else }}
        <li><a href="{{.Link}}">{{.Number}}</a></li>
      {{ end }}
    {{end}}
  </ul>
  {{end}}
  `
	tmplSummary = `{{if .total}}<div class="summary text-right">Displaying {{.start}}-{{.end}} of {{.total}} results.</div>{{end}}`
)
