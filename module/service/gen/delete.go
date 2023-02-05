package gen

import (
	"github.com/dabao-zhao/ddltoall/module/service/template"
	"github.com/dabao-zhao/ddltoall/output"
	"github.com/dabao-zhao/ddltoall/parser"
	"github.com/dabao-zhao/ddltoall/util/filex"
	"github.com/dabao-zhao/ddltoall/util/stringx"
)

func genDelete(table *parser.Table) (string, error) {
	camel := table.Name.ToCamel()
	text, err := filex.LoadTemplate(deleteMethodTemplateFile, template.DeleteMethod)
	if err != nil {
		return "", err
	}

	methodBuffer, err := output.With("delete").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"upperStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Title()),
			"data":                      table,
		})
	if err != nil {
		return "", err
	}

	return methodBuffer.String(), nil
}
