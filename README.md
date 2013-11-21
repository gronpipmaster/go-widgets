Go widgets
==========

Go widgets, pagination, breadcrumb, menu.

### Example ###

Pagination docs http://godoc.org/github.com/gronpipmaster/go-widgets/pagination
```go
pagination := pagination.New(total, currentPage, pageSize, "/me/page/")
pagination.SetNumLinks(10).Init()
//print in your template
// {{.pagination.Summary}}
// {{.pagination.Render}}

```

Breadcrumb docs http://godoc.org/github.com/gronpipmaster/go-widgets/breadcrumb
```go
breadcrumb := breadcrumb.New().Add("Home", "/").Add("Foo", "")
//print in your template
// {{.breadcrumb.Render}}
```

Menu docs http://godoc.org/github.com/gronpipmaster/go-widgets/menu
```go
sidebar := menu.New().AddItem("Bar", "/bar")
//print in your template
// {{.sidebar.Render}}
```


