package template

const ImportTpl = `import (
	"context"
	"errors"
	"math"

	"{{.bizPackage}}"
	"{{.typesPackage}}"

	"github.com/dabao-zhao/where-builder"
	"github.com/jinzhu/copier"
	"github.com/op/go-logging"
)
`
