package template

const ImportTpl = `import (
	"context"
	{{if .time}}"time"{{end}}

	"github.com/op/go-logging"
	"github.com/dabao-zhao/where-builder"
)
`
