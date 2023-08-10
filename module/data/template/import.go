package template

const ImportTpl = `import (
	"context"

	"{{.bizPackage}}"

	"github.com/dabao-zhao/where-builder"
	"github.com/op/go-logging"
	"gorm.io/gorm"
)
`
