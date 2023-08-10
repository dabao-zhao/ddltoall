package template

// TypeTpl defines a template for type in type.
const TypeTpl = `
type (
	Create{{.upperStartCamelObject}}Request struct {
		{{.fields}}
	}
	Create{{.upperStartCamelObject}}Response struct {
		{{.fieldsOnlyJsonTag}}
	}

	Update{{.upperStartCamelObject}}Request struct {
		{{.fields}}
	}
	Update{{.upperStartCamelObject}}Response struct {
		{{.fieldsOnlyJsonTag}}
	}

	Delete{{.upperStartCamelObject}}Request struct {
		{{.fields}}
	}
	Delete{{.upperStartCamelObject}}Response struct {
	}

	FindOne{{.upperStartCamelObject}}Request struct {
		{{.fields}}
	}
	FindOne{{.upperStartCamelObject}}Response struct {
		{{.fieldsOnlyJsonTag}}
	}

	FindAll{{.upperStartCamelObject}}Request struct {
		{{.fields}}
	}
	FindAll{{.upperStartCamelObject}}Response struct {
		Items []*FindAll{{.upperStartCamelObject}}ResponseItem {{.itemTag}}
	}
	FindAll{{.upperStartCamelObject}}ResponseItem struct {
		{{.fieldsOnlyJsonTag}}
	}

	Paginate{{.upperStartCamelObject}}Request struct {
		{{.fields}}

		{{.paginateRequestTag}}
	}
	Paginate{{.upperStartCamelObject}}Response struct {
		Items []*Paginate{{.upperStartCamelObject}}ResponseItem {{.itemTag}}

		{{.paginateResponseTag}}
	}
	Paginate{{.upperStartCamelObject}}ResponseItem struct {
		{{.fieldsOnlyJsonTag}}
	}
)
`

// TypeGen defines a template for type
var TypeGen = `
package {{.pkg}}
{{.imports}}
{{.types}}
`
