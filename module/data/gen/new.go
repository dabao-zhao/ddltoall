package gen

import (
	"fmt"
	"github.com/dabao-zhao/ddltoall/module/data/template"
	"github.com/dabao-zhao/ddltoall/output"
	"github.com/dabao-zhao/ddltoall/parser"
	"github.com/dabao-zhao/ddltoall/util/filex"
	"github.com/dabao-zhao/ddltoall/util/stringx"
)

func genNew(table *parser.Table) (string, error) {
	text, err := filex.LoadTemplate(dataNewTemplateFile, template.NewTpl)
	if err != nil {
		return "", err
	}

	t := fmt.Sprintf(`"%s"`, table.Name.Source())
	camel := table.Name.ToCamel()
	buffer, err := output.With("new").
		Parse(text).
		Execute(map[string]interface{}{
			"table":                 t,
			"upperStartCamelObject": camel,
			"lowerStartCamelObject": stringx.From(table.Name.ToCamel()).Untitle(),
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
