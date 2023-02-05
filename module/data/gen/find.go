package gen

import (
	"github.com/dabao-zhao/ddltoall/module/data/template"
	"github.com/dabao-zhao/ddltoall/output"
	"github.com/dabao-zhao/ddltoall/parser"
	"github.com/dabao-zhao/ddltoall/util/filex"
	"github.com/dabao-zhao/ddltoall/util/stringx"
)

func genFindOne(table *parser.Table) (string, error) {
	camel := table.Name.ToCamel()
	text, err := filex.LoadTemplate(findOneMethodTemplateFile, template.FindOneMethod)
	if err != nil {
		return "", err
	}

	methodBuffer, err := output.With("findOneMethod").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"upperStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Title()),
			"lowerStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", err
	}

	return methodBuffer.String(), nil
}
