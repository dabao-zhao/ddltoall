package gen

import (
	"github.com/dabao-zhao/ddltoall/module/biz/template"
	"github.com/dabao-zhao/ddltoall/output"
	"github.com/dabao-zhao/ddltoall/parser"
	"github.com/dabao-zhao/ddltoall/util/filex"
	"github.com/dabao-zhao/ddltoall/util/stringx"
)

func genUpdate(table *parser.Table) (string, string, error) {
	for _, field := range table.Fields {
		camel := stringx.SafeString(field.Name.ToCamel())
		if camel == "CreateTime" || camel == "UpdateTime" || camel == "CreateAt" || camel == "UpdateAt" {
			continue
		}

		if field.Name.Source() == table.PrimaryKey.Name.Source() {
			continue
		}

	}

	camelTableName := table.Name.ToCamel()
	text, err := filex.LoadTemplate(updateMethodTemplateFile, template.UpdateMethod)
	if err != nil {
		return "", "", err
	}

	methodBuffer, err := output.With("updateMethod").Parse(text).Execute(
		map[string]interface{}{
			"upperStartCamelObject": camelTableName,
			"lowerStartCamelObject": stringx.From(camelTableName).Untitle(),
			"data":                  table,
		},
	)
	if err != nil {
		return "", "", nil
	}

	// update interface method
	text, err = filex.LoadTemplate(updateInterfaceTemplateFile, template.UpdateMethodInterface)
	if err != nil {
		return "", "", err
	}

	interfaceBuffer, err := output.With("updateInterface").Parse(text).Execute(
		map[string]interface{}{
			"upperStartCamelObject": camelTableName,
			"data":                  table,
		},
	)
	if err != nil {
		return "", "", nil
	}

	return methodBuffer.String(), interfaceBuffer.String(), nil
}
