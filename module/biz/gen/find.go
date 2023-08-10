package gen

import (
	"github.com/dabao-zhao/ddltoall/module/biz/template"
	"github.com/dabao-zhao/ddltoall/output"
	"github.com/dabao-zhao/ddltoall/parser"
	"github.com/dabao-zhao/ddltoall/util/filex"
	"github.com/dabao-zhao/ddltoall/util/stringx"
)

func genFindOne(table *parser.Table) (string, string, error) {
	camel := table.Name.ToCamel()
	text, err := filex.LoadTemplate(findOneMethodTemplateFile, template.FindOneMethod)
	if err != nil {
		return "", "", err
	}

	methodBuffer, err := output.With("findOneMethod").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"lowerStartCamelObject":     stringx.From(camel).Untitle(),
			"lowerStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	text, err = filex.LoadTemplate(findOneInterfaceTemplateFile, template.FindOneInterface)
	if err != nil {
		return "", "", err
	}

	interfaceBuffer, err := output.With("findOneInterface").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"lowerStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	return methodBuffer.String(), interfaceBuffer.String(), nil
}

func genFindAll(table *parser.Table) (string, string, error) {
	camel := table.Name.ToCamel()
	text, err := filex.LoadTemplate(findAllMethodTemplateFile, template.FindAllMethod)
	if err != nil {
		return "", "", err
	}

	methodBuffer, err := output.With("findAllMethod").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"lowerStartCamelObject":     stringx.From(camel).Untitle(),
			"lowerStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	text, err = filex.LoadTemplate(findAllInterfaceTemplateFile, template.FindAllInterface)
	if err != nil {
		return "", "", err
	}

	interfaceBuffer, err := output.With("findAllInterface").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"lowerStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	return methodBuffer.String(), interfaceBuffer.String(), nil
}

func genPaginate(table *parser.Table) (string, string, error) {
	camel := table.Name.ToCamel()
	text, err := filex.LoadTemplate(paginateMethodTemplateFile, template.PaginateMethod)
	if err != nil {
		return "", "", err
	}

	methodBuffer, err := output.With("paginateMethod").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"lowerStartCamelObject":     stringx.From(camel).Untitle(),
			"lowerStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	text, err = filex.LoadTemplate(paginateInterfaceTemplateFile, template.PaginateInterface)
	if err != nil {
		return "", "", err
	}

	interfaceBuffer, err := output.With("paginateInterface").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"lowerStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	return methodBuffer.String(), interfaceBuffer.String(), nil
}

func genCount(table *parser.Table) (string, string, error) {
	camel := table.Name.ToCamel()
	text, err := filex.LoadTemplate(countMethodTemplateFile, template.CountMethod)
	if err != nil {
		return "", "", err
	}

	methodBuffer, err := output.With("countMethod").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"lowerStartCamelObject":     stringx.From(camel).Untitle(),
			"lowerStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	text, err = filex.LoadTemplate(countInterfaceTemplateFile, template.CountInterface)
	if err != nil {
		return "", "", err
	}

	interfaceBuffer, err := output.With("countInterface").
		Parse(text).
		Execute(map[string]interface{}{
			"upperStartCamelObject":     camel,
			"lowerStartCamelPrimaryKey": stringx.EscapeGolangKeyword(stringx.From(table.PrimaryKey.Name.ToCamel()).Untitle()),
			"dataType":                  table.PrimaryKey.DataType,
			"data":                      table,
		})
	if err != nil {
		return "", "", err
	}

	return methodBuffer.String(), interfaceBuffer.String(), nil
}
