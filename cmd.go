package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/dabao-zhao/ddltoall/command"
	"github.com/dabao-zhao/ddltoall/version"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ddltoall",
		Short: "A cli tool to generate model code",
		Long: "A cli tool to generate model code\n\n" +
			"GitHub: https://github.com/dabao-zhao/ddltoall\n" +
			"Copy From: https://github.com/zeromicro/go-zero/tools/goctl",
	}

	ddlCmd = &cobra.Command{
		Use:   "ddl",
		Short: "Generate all from ddl",
		RunE:  command.MysqlDDL,
	}

	datasourceCmd = &cobra.Command{
		Use:   "datasource",
		Short: "Generate all from datasource",
		RunE:  command.MySqlDataSource,
	}
)

func init() {
	ddlCmd.Flags().StringVarP(&command.VarStringSrc, "src", "s", "", "The path or path globbing patterns of the ddl")
	ddlCmd.Flags().StringVarP(&command.VarStringDir, "dir", "d", "", "The target dir")
	ddlCmd.Flags().StringVar(&command.VarStringDatabase, "database", "", "The name of database [optional]")

	datasourceCmd.Flags().StringVarP(&command.VarStringURL, "url", "u", "", `The data source of database,like "root:password@tcp(127.0.0.1:3306)/database"`)
	datasourceCmd.Flags().StringSliceVarP(&command.VarStringSliceTable, "table", "t", nil, "The table or table globbing patterns in the database")
	datasourceCmd.Flags().StringVarP(&command.VarStringDir, "dir", "d", "", "The target dir")
}

func main() {
	rootCmd.Version = fmt.Sprintf(
		"%s %s/%s", version.BuildVersion,
		runtime.GOOS, runtime.GOARCH)

	rootCmd.AddCommand(ddlCmd)
	rootCmd.AddCommand(datasourceCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Println(aurora.Red(err.Error()))
		os.Exit(1)
	}
}
