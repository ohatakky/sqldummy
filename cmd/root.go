package cmd

import (
	"log"

	"github.com/k0kubun/pp"
	"github.com/ohatakky/sqldummy/adapter/mysql"
	"github.com/ohatakky/sqldummy/internal/parser"
	"github.com/spf13/cobra"
)

var (
	user     string
	password string
	host     string
	port     int
	database string
	table    string

	rootCmd = &cobra.Command{
		Use:   "sqldummy",
		Short: "sqldummy is dummy data generator for mysql.",
		Long:  "sqldummy is dummy data generator for mysql.",
		Run: func(cmd *cobra.Command, args []string) {
			db, err := mysql.New(user, password, host, database, port)
			if err != nil {
				log.Fatal(err)
			}
			err = Run(db, table)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolP("help", "", false, "helping myself with the command "+rootCmd.Name())
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "mysql user")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "mysql password")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "P", 0, "mysql port")
	rootCmd.PersistentFlags().StringVarP(&host, "host", "h", "", "mysql host")
	rootCmd.PersistentFlags().StringVarP(&database, "database", "d", "", "mysql database")
	rootCmd.PersistentFlags().StringVarP(&table, "table", "t", "", "mysql table")
}

func Run(db *mysql.Mysql, table string) error {
	// todo: parse
	p := parser.New(db)
	ddl, err := p.DumpTableDDL(table)
	if err != nil {
		return err
	}
	tm, err := p.ParseDDL(ddl)
	if err != nil {
		return err
	}
	pp.Println(tm)

	// todo: generate

	// todo: insert

	return nil
}
