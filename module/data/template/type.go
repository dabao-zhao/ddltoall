package template

// TypeTpl defines a template for types in service.
const TypeTpl = `
type (
	{{.upperStartCamelObject}}Repo struct {
		db     *gorm.DB
		logger *logging.Logger
	}
)
`
