package template

// TypeTpl defines a template for types in biz.
const TypeTpl = `
type (
	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}

	{{.lowerStartCamelObject}}Repo interface{
		{{.method}}
	}

	{{.upperStartCamelObject}}UseCase struct {
		repo   {{.lowerStartCamelObject}}Repo
		logger *logging.Logger
	}
)
`
