package gen

import (
	"log"
	"path/filepath"

	"github.com/dabao-zhao/ddltoall/model"
	bizGen "github.com/dabao-zhao/ddltoall/module/biz/gen"
	dataGen "github.com/dabao-zhao/ddltoall/module/data/gen"
	serviceGen "github.com/dabao-zhao/ddltoall/module/service/gen"
	typeGen "github.com/dabao-zhao/ddltoall/module/type/gen"
	"github.com/dabao-zhao/ddltoall/parser"
	"github.com/dabao-zhao/ddltoall/util/filex"
	"github.com/dabao-zhao/ddltoall/util/stringx"
)

const pwd = "."

type (
	DefaultGenerator struct {
		originDir string
		dir       string
		pkg       string
	}
)

// NewDefaultGenerator creates an instance for defaultGenerator
func NewDefaultGenerator(dir string) (*DefaultGenerator, error) {
	if dir == "" {
		dir = pwd
	}
	originDir := dir
	dirAbs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	dir = dirAbs
	pkg := stringx.SafeString(filepath.Base(dirAbs))
	err = filex.MkdirIfNotExist(dir)
	if err != nil {
		return nil, err
	}

	generator := &DefaultGenerator{
		originDir: originDir,
		dir:       dir,
		pkg:       pkg,
	}

	return generator, nil
}

func (g *DefaultGenerator) StartFromDDL(filename string, strict bool, database string) error {
	tables, err := parser.Parse(filename, database, strict)
	if err != nil {
		return err
	}
	return g.CreateFromTables(tables)
}

func (g *DefaultGenerator) StartFromInformationSchema(tables map[string]*model.Table, strict bool) error {
	var m []*parser.Table
	for _, each := range tables {
		table, err := parser.ConvertDataType(each, strict)
		if err != nil {
			return err
		}
		m = append(m, table)
	}
	return g.CreateFromTables(m)
}

func (g *DefaultGenerator) CreateFromTables(tables []*parser.Table) error {
	biz, err := bizGen.NewDefaultGenerator(g.dir + "/biz")
	if err != nil {
		return err
	}
	err = biz.CreateFromTables(tables)
	if err != nil {
		return err
	}

	typ, err := typeGen.NewDefaultGenerator(g.dir + "/types")
	if err != nil {
		return err
	}
	err = typ.CreateFromTables(tables)
	if err != nil {
		return err
	}

	service, err := serviceGen.NewDefaultGenerator(g.dir+"/service", g.originDir)
	if err != nil {
		return err
	}
	err = service.CreateFromTables(tables)
	if err != nil {
		return err
	}

	data, err := dataGen.NewDefaultGenerator(g.dir+"/data", g.originDir)
	if err != nil {
		return err
	}
	err = data.CreateFromTables(tables)
	if err != nil {
		return err
	}

	log.Println("done.")
	return nil
}
