Go widgets
==========

Go widgets, pagination, breadcrumb, menu.

### Example ###

Pagination:
```go
pagination := pagination.New(total, currentPage, pageSize, "/me/page/")
pagination.SetNumLinks(10).Init()
//print in your template
// {{.pagination.Summary}}
// {{.pagination.Render}}

```

Breadcrumb:
```go
breadcrumb := breadcrumb.New().Add("Home", "/").Add("Foo", "")
//print in your template
// {{.breadcrumb.Render}}
```

Menu:
```go
sidebar := menu.New().AddItem("Bar", "/bar")
//print in your template
// {{.sidebar.Render}}
```


